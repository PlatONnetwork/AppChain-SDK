#!/bin/bash

declare -A SIMAPP_ENV_TEMPLATE=(
    ["OUTPUT"]=""
    ["BIN"]=""
    ["DEPLOY_KEY"]=""
    ["INITIALIZE_KEY"]=""
    ["ROOTCHAIN_URL"]=""
    ["GENESIS_NODE"]=""
    ["NORMAL_NODE"]=""
    ["REGISTRY_MANAGER_OWNER"]=""
    ["CHILDCHAIN_OWNER"]=""
    ["STAKE_OWNER"]=""
    ["STAKE_AMOUNT"]="1000000000"
    ["USERNAME"]=""
    ["PASSWORD"]=""
    ["CHECKPOINT_KEYSTORE"]="l1checkpointsender.json"
    ["CHECKPOINT_KEYSTORE_PASSWORD"]="l1checkpointsender_password"
    ["L2TX_KEYSTORE"]="l2txsender.json"
    ["L2TX_KEYSTORE_PASSWORD"]="l2txsender_password"
    ["START_ARGS"]="--identity node --verbosity 4 --http.api platon,debug,personal,admin,net,web3,txpool,benchmark,checkpoint --http.vhosts \"*\" --cache 256 --metrics --ipcdisable  --maxpeers 100 --maxconsensuspeers 75 --txpool.globalslots 1000000 --txpool.accountslots 1000000 --txpool.globalqueue 1000000 --txpool.accountqueue 1000000 --txpool.cacheSize 1000000 --txpool.globaltxcount 1000000 --networkid 102 --allow-insecure-unlock"
    ["EXTRA_ARGS"]="--log.debug --asyncblock.computersenderthread 4 --asyncblock.entrysize 2000 --asyncblock.splitthreshold 3000 --asyncblock.concurrency_level 7 --asyncblock.txs_batch 512 --cbft.wal.disabled --nontxpool.broadcast"
    ["STAKE_WITHDRAWAL_WAIT_PERIOD"]="2"
    ["STAKE_DELEGATE_WITHDRAWAL_WAIT_PERIOD"]="2"
)
SIMAPP_ENV_TEMPLATE_KEY=(
    OUTPUT
    BIN
    DEPLOY_KEY
    INITIALIZE_KEY
    ROOTCHAIN_URL
    GENESIS_NODE
    NORMAL_NODE
    REGISTRY_MANAGER_OWNER
    CHILDCHAIN_OWNER
    STAKE_OWNER
    STAKE_AMOUNT
    USERNAME
    PASSWORD
    CHECKPOINT_KEYSTORE
    CHECKPOINT_KEYSTORE_PASSWORD
    L2TX_KEYSTORE
    L2TX_KEYSTORE_PASSWORD
    START_ARGS
    EXTRA_ARGS
    STAKE_WITHDRAWAL_WAIT_PERIOD
    STAKE_DELEGATE_WITHDRAWAL_WAIT_PERIOD
)
genSimappEnv(){
    file=$1
    local  overrides="$2" 
    {
        echo "#!/bin/bash"
        echo "#"
        echo "# $comment"
        echo "# Generated on: $(date)"
        echo "#"
        echo ""
        
        for key in "${SIMAPP_ENV_TEMPLATE_KEY[@]}"; do
            local value="${overrides[$key]:-${SIMAPP_ENV_TEMPLATE[$key]}}"
            echo "export $key='$value'"
        done
    } > $file
}