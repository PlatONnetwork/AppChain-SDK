#!/bin/bash

declare -A ENV_TEMPLATE=(
    ["userkey"]=""
    ["rootchainurl"]=""
    ["childchainurl"]=""
    ["user"]=""
    ["token"]=""
    ["l1staking"]=""
    ["l1depositmanager"]=""
    ["l1exithelper"]=""
    ["l2statesender"]="0x1000000000000000000000000000000000000001"
    ["l2statesync"]="0x1000000000000000000000000000000000000002"
    ["l2staking"]="0x1000000000000000000000000000000000000005"
    ["l2deposit"]="0x1000000000000000000000000000000000000007"
    ["l2votetoken"]="0x100000000000000000000000000000000000000B"
    ["validator"]=""
)
ENV_TEMPLATE_KEY=(userkey rootchainurl childchainurl user token l1staking l1depositmanager l1exithelper l2statesender l2statesync l2staking l2deposit l2votetoken validatoraddr validator)
genEnv(){
    file=$1
    local -n overrides="$2" 
    {
        echo "#!/bin/bash"
        echo "#"
        echo "# $comment"
        echo "# Generated on: $(date)"
        echo "#"
        echo ""
        
        for key in "${ENV_TEMPLATE_KEY[@]}"; do
            local value="${overrides[$key]:-${ENV_TEMPLATE[$key]}}"
            echo "export $key='$value'"
        done
    } > $file
}