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
source benchmarkenvtmpl.sh
declare -A nodesMap
declare -A childchainContracts
declare -a nodeDirs
declare -A envMap
declare -A benchmarkEnvMap
createToolConfig(){
    rootDir=$1
    owner=$2
    stakeAmount=$3
    key=$4
    user=$5
    rootchainurl=$6
    childchainurl=$7
    toolsPath=$8
    benchmarkclient=$9
    targetAddr=${10}
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
        $(cp -f $benchmarkclient $rootDir/$node/scripts/bin/benchmarkclient)
        $(cp -rf $PROJECT_ROOT/lib $rootDir/$node/scripts/)
        genEnv "$rootDir/$node/scripts/config/env" envMap
        genBenchmarkEnv "$rootDir/$node/scripts/config/benchmarkenv" $benchmarkEnvMap
        debug "create config success, validator:$validatorAddress, node:$node"

    done
}
createEmptyToolConfig(){
    sampleNodeDir=$1
    newNodeDir=$2
    rootchainurl=$3
    childchainurl=$4
    toolsPath=$5
    benchmarkclient=$6
    $(cp -f $sampleNodeDir/*.sh $newNodeDir/)
    $(cp -f $sampleNodeDir/genesis.json $newNodeDir/)
    $(cp -f $sampleNodeDir/node $newNodeDir/)

    envMap["rootchainurl"]=$rootchainurl
    envMap["childchainurl"]=$childchainurl
    envMap["token"]=${childchainContracts["token"]}
    envMap["l1staking"]=${childchainContracts["stake_manager"]}
    envMap["l1depositmanager"]=${childchainContracts["deposit_manager"]}
    envMap["l1exithelper"]=${childchainContracts["exit_helper"]}
    
    $(mkdir -p $newNodeDir/scripts/config)
    $(cp -rf $PROJECT_BASE/tools/* $newNodeDir/scripts)
    $(cp -f $toolsPath $newNodeDir/scripts/bin/tools)
    $(cp -f $benchmarkclient $newNodeDir/scripts/bin/benchmarkclient)
    $(cp -rf $PROJECT_ROOT/lib $newNodeDir/scripts/)
    genEnv "$newNodeDir/scripts/config/env" envMap
    genBenchmarkEnv "$newNodeDir/scripts/config/benchmarkenv" $benchmarkEnvMap
    debug "create config success, $newNodeDir"

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
    benchmarkclient=$9
    targetAddr=${10}
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
    createToolConfig $nodeDir $owner $stakeAmount $key $user $rootchainurl $childchainurl $toolsPath $benchmarkclient $targetAddr
    
}

createEmpty(){
    output=$1
    rootchainurl=$2
    childchainurl=$3
    toolsPath=$4
    bechmarkclient=$5
    newNodeDir=$6
    mkdir -p $newNodeDir

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
    createEmptyToolConfig "$nodeDir/${nodeDirs[0]}" $newNodeDir $rootchainurl $childchainurl $toolsPath $bechmarkclient
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
  all       All scripts generated
             args: [output] [owner] [stakeAmount] [key] [user] [rootchainurl] [childchainurl] [toolsPath] [benchmarkclient]
  single    Script to generate the folder of validator addresses,
                args: [output] [owner] [stakeAmount] [key] [user] [rootchainurl] [childchainurl] [toolsPath] [benchmarkclient] [targetAddr]
  empty     Script to generate the empty folder,
                args: [output] [rootchainurl] [childchainurl] [toolsPath] [benchmarkclient] [newNodeDir]
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
    newNodeDir  generate node directory
EOF
}
case $1 in
    "all")
        deploy $2 $3 $4 $5 $6 $7 $8 $9 ${10}
        ;;
    "single") 
        deploy $2 $3 $4 $5 $6 $7 $8 $9 ${10} ${11}
    ;;
    "empty")
        createEmpty $2 $3 $4 $5 $6 $7
        ;;
    *)
        showHelp
    ;;
esac
