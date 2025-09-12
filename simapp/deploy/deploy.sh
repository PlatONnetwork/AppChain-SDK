SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_BASE="$(dirname "$SCRIPT_DIR")"
PROJECT_ROOT="$(dirname "$PROJECT_BASE")"
source "${PROJECT_ROOT}/lib/core.sh"
loadLibraries
debug "SCRIPT_DIR:$SCRIPT_DIR"
debug "PROJECT_ROOT:$PROJECT_ROOT"

output=

nodes=
source decodenodes.sh
source decodechildcontract.sh
source findnodedir.sh
source envtmpl.sh
declare -A nodesMap
declare -A childchainContracts
declare -a nodeDirs
declare -A envMap
createToolConfig(){
    rootDir=$1
    owner=$2
    stakeAmount=$3
    key=$4
    user=$5
    rootchainurl=$6
    childchainurl=$7
    toolsPath=$8
    targetAddr=$9
    envMap["userkey"]=$key
    envMap["user"]=$user
    envMap["rootchainurl"]=$rootchainurl
    envMap["childchainurl"]=$childchainurl
    envMap["token"]=${childchainContracts["token"]}
    envMap["l1staking"]=${childchainContracts["stake_manager"]}
    envMap["l1depositmanager"]=${childchainContracts["deposit_manager"]}
    envMap["l1exithelper"]=${childchainContracts["exit_helper"]}
    for node in "${nodeDirs[@]}"; do
        nodePubKey=$(cat "$rootDir/$node/nodepub")
        blsPubKey=$(cat "$rootDir/$node/blspub")
        validatorAddress=${nodesMap[$nodePubKey]}
        
        
        if [[ ${#targetAddr} -ne 0 && "$targetAddr" -ne "$validatorAddress" ]]; then
            debug "unmatch target address, stop create config, $targetAddr"
            continue
        fi
        validator="{\"owner\":\"$owner\",\"stakeAmount\":\"$stakeAmount\",\"commissionRate\":2,\"pubKey\":\"0x$nodePubKey\",\"blsKey\":\"0x$blsPubKey\"}"
        envMap["validator"]=$validator
        envMap["validatoraddr"]=$validatorAddress
        $(mkdir -p $rootDir/$node/scripts/config)
        $(cp -rf $PROJECT_BASE/tools/* $rootDir/$node/scripts)
        $(cp -f $toolsPath $rootDir/$node/scripts/bin/tools)
        $(cp -rf $PROJECT_ROOT/lib $rootDir/$node/scripts/)
        genEnv "$rootDir/$node/scripts/config/env" envMap
        debug "create config success, validator:$validatorAddress, node:$node"

    done
}

deploy(){
    output=$1
    owner=$2
    stakeAmount=$3
    key=$4
    user=$5
    rootchainurl=$6
    childchainurl=$7
    toolsPath=$8
    targetAddr=$9
    nodeTomlPath="$output/nodes.toml"
    childChainContractsTomlPath="$output/childchaincontract.toml"
    nodeDir="$output/ansible/playbooks/files"

    debug "nodeTomlPath:$nodeTomlPath"
    debug "childChainContractsTomlPath:$childChainContractsTomlPath"
    debug "nodedir:$nodeDir"

    debug "parse nodes.toml"
    decodeNodesToml $nodeTomlPath nodesMap
    debug "parse childchaincontracts.toml"
    decodeChildcontractToml $childChainContractsTomlPath childchainContracts
    debug "parse nodes dirs"
    getNodeFolders $nodeDir nodeDirs
    if [[ ${#nodesMap[@]} -eq 0 ||  ${#childchainContracts[@]} -eq 0 || ${#nodeDirs[@]} -eq 0 ]]; then  
        echo "element is empty, nodesMap:${#nodesMap[@]}, childchainContracts:${#childchainContracts[@]}, nodeDirs:${#nodeDirs[@]}"  
    fi
    createToolConfig $nodeDir $owner $stakeAmount $key $user $rootchainurl $childchainurl $toolsPath $targetAddr
    
}


printConfig() {
    for key in "${!childchainContracts[@]}"; do
        echo "$key: ${childchainContracts[$key]}"
    done
    for key in "${!nodesMap[@]}"; do
        echo "$key: ${nodesMap[$key]}"
    done
    for node in "${nodeDirs[@]}"; do
        echo "  $node"
    done
}

showHelp(){
    cat << EOF
Usage: $0 COMMAND [ARGS]

A simple staking operation script.
    
Commands:
  all      All scripts generated
             args: [output] [owner] [stakeAmount] [key] [user] [rootchainurl] [childchainurl] [toolsPath]
  single       Script to generate the folder of validator addresses,
                args: [output] [owner] [stakeAmount] [key] [user] [rootchainurl] [childchainurl] [toolsPath] [targetAddr]
Options:
    output  simapp output directory
    owner   validator's owner
    stakeAmount  validator stake amount
    key     transactions sender key
    user    transactions sender address
    rootchainurl layer1 rpc url
    childchainurl layer2 rpc url
    toolsPath    tools binary path
    targetAddr   generate validator config
EOF
}
case $1 in
    "all")
        deploy $2 $3 $4 $5 $6 $7 $8 $9
        ;;
    "single")
        deploy $2 $3 $4 $5 $6 $7 $8 $9 ${10}
    ;;
    *)
        showHelp
    ;;
esac
