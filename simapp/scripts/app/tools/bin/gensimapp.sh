#!/bin/bash
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
source "${PROJECT_ROOT}/lib/core.sh"
loadLibraries
configDir="${PROJECT_ROOT}/config"
simappEnvfile="${PROJECT_ROOT}/config/simappenv"
benchmarkEnvfile="${PROJECT_ROOT}/config/benchmarkenv"

node=$(realpath "../../bin/simapp")
defaultoutputDir=$(realpath "../output")
defaultNodeDir=../..
defaultPassword="123456"
defaultRootchainUrl="https://devnet2openapi.platon.network/rpc"
deployAddr=
initializeAddr=

genNodeKey(){
    local nodeDir=$1
    local addrfile=$2
    local keyfile=$3
    local pubfile=$4
    if [[ "$1" == "" ]];then
        nodeDir=$defaultNodeDir
    fi
    result=$($node keytool genkeypair)
    address=$(echo "$result" | grep -oP 'Address:\s*\K[0-9a-zA-Z]+')
    privateKey=$(echo "$result" | grep -oP 'PrivateKey:\s*\K[0-9a-f]+')  
    publicKey=$(echo "$result" | grep -oP 'PublicKey\s*:\s*\K[0-9a-f]+')
    echo -n $address > "$addrfile"
    echo -n $privateKey > "$keyfile"
    echo -n $publicKey > "$pubfile"
    debug "generate key success"

}
genKeyfile(){
    local password=$1
    local keyfile=$2
    local passwordFile=$3
    rm -f $keyfile
    result=$(echo -e "$password\n$password" | $node keytool generate $keyfile)
    echo -n $password > "$passwordFile"
    debug "generate $keyfile, $passwordFile success"
}
parseAddressWithSed() {
    local keyfile="$1"
    
    local address=$(cat $keyfile | sed -n 's/.*"address":"\([^"]*\)".*/\1/p')
    
    if [ -n "$address" ]; then
        echo "$address"
        return 0
    else
        error "error: Address not found in JSON"
        return 1
    fi
}
replaceEnvContent() {  
    local file_path="$1"
    local var="$2"  
    local new_env_content="$3"  
    
    # Check if file exists  
    if [ ! -f "$file_path" ]; then  
        echo "Error: File '$file_path' does not exist" >&2  
        return 1  
    fi  
    
    # Escape special characters for sed  
    local escaped_content=$(echo "$new_env_content" | sed 's/[\/&]/\\&/g')  
    
    # Use sed to replace validator= content  
    sed -i "s/^export $var=.*$/export $var='$escaped_content'/" "$file_path"      
} 

replaceSimappEnvContent(){
    nodeIps=$1
    user=$2
    password=$3
    rootchainurl=$4
    if [[ -z $rootchainurl ]];then
        rootchainurl=$defaultRootchainUrl
    fi
    result=$($node keytool genkeypair)
    addr=$(echo "$result" | grep -i "Address:" | awk '{print $2}')
    privateKey=$(echo "$result" | grep -oP 'PrivateKey:\s*\K[0-9a-f]+')  
    publicKey=$(echo "$result" | grep -oP 'PublicKey\s*:\s*\K[0-9a-f]+')
    replaceEnvContent $simappEnvfile "DEPLOY_KEY" $privateKey
    deployAddr=$addr
    result=$($node keytool genkeypair)
    addr=$(echo "$result" | grep -i "Address:" | awk '{print $2}')
    privateKey=$(echo "$result" | grep -oP 'PrivateKey:\s*\K[0-9a-f]+')  
    publicKey=$(echo "$result" | grep -oP 'PublicKey\s*:\s*\K[0-9a-f]+')
    replaceEnvContent $simappEnvfile "INITIALIZE_KEY" $privateKey
    initializeAddr=$addr
    replaceEnvContent $simappEnvfile "REGISTRY_MANAGER_OWNER" $initializeAddr
    replaceEnvContent $simappEnvfile "CHILDCHAIN_OWNER" $privateKey
    replaceEnvContent $simappEnvfile "STAKE_OWNER" $initializeAddr
    replaceEnvContent $simappEnvfile "GENESIS_NODE" $nodeIps
    replaceEnvContent $simappEnvfile "ROOTCHAIN_URL" $rootchainurl
    replaceEnvContent $simappEnvfile "USERNAME" $user
    replaceEnvContent $simappEnvfile "PASSWORD" $password
    replaceEnvContent $simappEnvfile "OUTPUT" $defaultoutputDir
    replaceEnvContent $simappEnvfile "BIN" $node
    replaceEnvContent $simappEnvfile "CHECKPOINT_KEYSTORE" $configDir/l1checkpointsender.json
    replaceEnvContent $simappEnvfile "CHECKPOINT_KEYSTORE_PASSWORD" $configDir/l1checkpointsender_password
    replaceEnvContent $simappEnvfile "L2TX_KEYSTORE" $configDir/l2txsender.json 
    replaceEnvContent $simappEnvfile "L2TX_KEYSTORE_PASSWORD" $configDir/l2txsender_password
}

replaceBenchmarkEnvContent(){
    nodeIps=$1
    IFS=',' read -ra nodeArray <<< "$nodeIps"
    urlArray=()
    for ip in "${nodeArray[@]}"; do
        urlArray+=("http://${ip}:8801")
    done
    result=$(IFS=,; echo "${urlArray[*]}")
    replaceEnvContent $benchmarkEnvfile "BENCHMARK_URLS" $result
}

generate(){
    ips=$1
    username=$2
    password=$3
    rootchainurl=$4
    mkdir -p $defaultoutputDir
    replaceSimappEnvContent $ips $username $password $rootchainurl
    replaceBenchmarkEnvContent $ips
    genKeyfile $defaultPassword $configDir/l1checkpointsender.json $configDir/l1checkpointsender_password
    latAddr=$(parseAddressWithSed $configDir/l1checkpointsender.json)
    l1checkpointAddr=$(bech32Decode $latAddr)
    debug "l1checkpointaddr:$l1checkpointAddr"
    genKeyfile $defaultPassword $configDir/l2txsender.json $configDir/l2txsender_password
    echo "generate keys success please contact the official team to ensure sufficient balance for both the deploy address:$deployAddr, initialize address:$initializeAddr and checkpoint address:$l1checkpointAddr"

}
showHelp() {
    cat << EOF
Usage: $0 [IPS] [USERNAME] [PASSWORD] [ROOTCHAINURL]

A simple gen simapp script.
IPS         The IP addresses of the servers to be deployed, connected by commas, e.g. 10.1.1.33,10.1.1.34,10.1.1.35,10.1.1.36.
USERNAME    Server username
PASSWORD    Server password
ROOTCHAINURL Rootchain url, default:https://devnet2openapi.platon.network/rpc
EOF
}
case $1 in
    "gen")
        generate $2 $3 $4 $5
        ;;
    "help"|"--help")
        showHelp
        ;;
esac