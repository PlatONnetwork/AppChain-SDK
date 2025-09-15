#!/bin/bash
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
source "${PROJECT_ROOT}/lib/core.sh"
loadLibraries
envfile="${PROJECT_ROOT}/config/env"

node=../../node
defaultNodeDir=../..
defaultPassword="123456"

owner=
stakeAmount=
l1checkpointAddr=
userAddr=

genNodeKey(){
    local nodeDir=$1
    if [[ "$1" == "" ]];then
        nodeDir=$defaultNodeDir
    fi
    result=$($node keytool genkeypair)
    address=$(echo "$result" | grep -oP 'Address:\s*\K[0-9a-f]+')
    privateKey=$(echo "$result" | grep -oP 'PrivateKey:\s*\K[0-9a-f]+')  
    publicKey=$(echo "$result" | grep -oP 'PublicKey\s*:\s*\K[0-9a-f]+')
    echo $address > "$nodeDir/nodeaddr"
    echo $privateKey > "$nodeDir/nodekey"
    echo $publicKey > "$nodeDir/nodepub"
    debug "generate node key success"

}

genBlsKey(){
    local nodeDir=$1
    if [[ "$1" == "" ]];then
        nodeDir=$defaultNodeDir
    fi
    result=$($node keytool genblskeypair)
    privateKey=$(echo "$result" | grep -oP 'PrivateKey:\s*\K[0-9a-f]+')  
    publicKey=$(echo "$result" | grep -oP 'PublicKey\s*:\s*\K[0-9a-f]+')
    echo $privateKey > "$nodeDir/blskey"
    echo $publicKey > "$nodeDir/blspub"
    debug "generate bls key success"
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
genKeyfile(){
    local password=$1
    local keyfile=$2
    local passwordFile=$3
    rm -f $keyfile
    result=$(echo -e "$password\n$password" | $node keytool generate $keyfile)
    echo $password > "$passwordFile"
    debug "generate $keyfile, $passwordFile success"
}
genL1CheckpointKey(){
    local nodeDir=$1
    local password=$2
    local keyfile=$3
    local passwordFile=$4
    if [[ "$1" == "" ]];then
        nodeDir=$defaultNodeDir
    fi
    if [[ "$2" == "" ]];then
        password=$defaultPassword
    fi
    if [[ "$3" == "" ]];then
        keyfile="$nodeDir/l1checkpointsender.json"
    fi
    if [[ "$4" == "" ]];then
        passwordFile="$nodeDir/l1checkpointsender_password"
    fi
    genKeyfile $password $keyfile $passwordFile
    latAddr=$(parseAddressWithSed $keyfile)
    l1checkpointAddr=$(bech32Decode $latAddr)
    debug "l1checkpointaddr:$l1checkpointAddr"
}

genL2TxSenderKey(){
    local nodeDir=$1
    local password=$2
    local keyfile=$3
    local passwordFile=$4
    if [[ "$1" == "" ]];then
        nodeDir=$defaultNodeDir
    fi
    if [[ "$2" == "" ]];then
        password=$defaultPassword
    fi
    if [[ "$3" == "" ]];then
        keyfile="$nodeDir/l2txsender.json"
    fi
    if [[ "$4" == "" ]];then
        passwordFile="$nodeDir/l2txsender_password"
    fi
    genKeyfile $password $keyfile $passwordFile
    addr=$(parseAddressWithSed $keyfile)
    addr=$(bech32Decode $addr)
    debug "l2txsenderaddr:$addr"
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
genUserKey(){
    result=$($node keytool genkeypair)
    userAddr=$(echo "$result" | grep -i "Address:" | awk '{print $2}')
    privateKey=$(echo "$result" | grep -oP 'PrivateKey:\s*\K[0-9a-f]+')  
    publicKey=$(echo "$result" | grep -oP 'PublicKey\s*:\s*\K[0-9a-f]+')
    replaceEnvContent $envfile "userkey" $privateKey
    replaceEnvContent $envfile "user" $userAddr
    if [ -z "$owner" ]; then
        owner=$userAddr
    fi
    debug "generate userkey success, addr:$userAddr"
}

 
genValidatorStakeFor(){
    local nodeDir=$1
    if [[ "$1" == "" ]];then
        nodeDir=$defaultNodeDir
    fi
    validatorAddress=$(cat "$nodeDir/nodeaddr")
    nodePubKey=$(cat "$nodeDir/nodepub")
    blsPubKey=$(cat "$nodeDir/blspub")
    validator="{\"owner\":\"$owner\",\"stakeAmount\":\"$stakeAmount\",\"commissionRate\":2,\"pubKey\":\"0x$nodePubKey\",\"blsKey\":\"0x$blsPubKey\"}"
    replaceEnvContent $envfile "validatoraddr" $validatorAddress
    replaceEnvContent $envfile "validator" $validator
}


#!/bin/bash

# Enhanced parsing with getopt (GNU version)
parseArgsWithGetopt() {
    # Define options
    local options="o:s:n:h"
    local long_options="owner:,stakeamount:,nodedir:,help"
    
    # Parse arguments
    local parsed_args
    if ! parsed_args=$(getopt -o "$options" -l "$long_options" -n "$0" -- "$@"); then
        exit 1
    fi
    
    eval set -- "$parsed_args"
    
    owner=""
    local help=false
    local remaining_args=()
    
    while true; do
        case "$1" in
            -o|--owner)
                owner="$2"
                shift 2
                ;;
            -s|--stakeamount)
                stakeAmount="$2"
                shift 2
                ;;
            -n|--nodedir)
                defaultNodeDir="$2"
                shift 2
                ;;
            -h|--help)
                help=true
                shift
                ;;
            --)
                shift
                remaining_args=("$@")
                break
                ;;
            *)
                echo "Error: Invalid option" >&2
                exit 1
                ;;
        esac
    done
    
    if [ "$help" = true ]; then
        showHelp
        exit 0
    fi

    
    if [ -z "$stakeAmount" ]; then
        echo "Error: --stakeamount is required" >&2
        exit 1
    fi    
}

showHelp(){
    cat << EOF
Usage: $0 [OPTIONS]

Options:
  -o, --owner ADDRESS   Owner address (default:\$user)
  -s, --stakeamount AMOUNT         validator stake amount (required)
  -n, --nodedir DIRECTORY   node directory(default ../../)
  -h, --help            Show this help message

Examples:
  $0 --stakeamount 1000000000 
EOF
}
gen(){
    genNodeKey
    genBlsKey
    genL1CheckpointKey
    genL2TxSenderKey
    genUserKey
    genValidatorStakeFor

    echo "generate keys success please contact the official team to ensure sufficient balance for both the user address:$userAddr and checkpoint address:$l1checkpointAddr"
}

parseArgsWithGetopt $@
gen