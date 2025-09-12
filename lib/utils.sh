#!/bin/bash

getReceiptHash(){
    log=$1
    local hash=$(echo "$log" | grep -o '"hash": "[^"]*"' | awk -F'"' '{print $4}')
    echo $hash
}

hexToInt(){
    idhex=$1
    local id=$(printf "%d" $(echo "$idhex" | tr -d '"'))
    echo $id
}

waitCondition(){
    local cmd=$1
    local condition=$2
    local msg=$3
    local time=0
    while true; do
		result=$(eval "$cmd")
        flag=$("$condition" "$result")
		if [ "$flag" -eq "1" ]; then
			echo "$msg wait success"
			break
		else
			echo -ne "result:[$result] had elapse ${time}s retry wait 5s...\r"
			sleep 5
            time=$((time+5))
		fi
	done
}

checkTx(){
	local node=$1
	local url=$2
	local hash=$3
	local status=$($node attach $url --exec "platon.getTransactionReceipt(\"$hash\").status")
	echo "status:$status"
	if [ $status == "\"0x0\"" ];then
		debug "transaction receipt is failed:$hash"
		exit
	else
		debug "transaction receipt is success:$hash"
	fi
}
checkTxHash() {  
    local tx_hash="$1"  
    
    if [[ -z "$tx_hash" ]]; then  
        debug "Error: Transaction hash cannot be empty"  
        exit  
    fi  
    
    if [[ ${#tx_hash} -ne 66 && ${#tx_hash} -ne 64 ]]; then  
        debug "Error: Transaction hash length should be 64 or 66 characters (current: ${#tx_hash})"  
        exit 
    fi  
    
    if [[ "$tx_hash" == 0x* ]]; then  
        local without_prefix="${tx_hash#0x}"  
        if [[ ${#without_prefix} -ne 64 ]]; then  
            debug "Error: Hash length after 0x prefix should be 64 characters"  
            exit
        fi  
        tx_hash="$without_prefix"  
    fi  
    
    if [[ ! "$tx_hash" =~ ^[0-9a-fA-F]{64}$ ]]; then  
        debug "Error: Not a valid hexadecimal hash"  
        exit
    fi      
}
waitcheckpoint(){
    id=$1
	debug "prepare wait checkpoint"
	equalStr(){
		result=$1
		if [[ "$result" == *"ExitEvent"* ]]; then
   			echo 1
		else
			echo 0
		fi
	}
	cmd="curl -s --noproxy '*' '$childchainurl' -X POST -H 'Content-Type: application/json' --data '{\"jsonrpc\":\"2.0\",\"method\":\"checkpoint_generateExitProof\",\"params\":['$id'],\"id\":0}'"
	waitCondition "$cmd" equalStr "waitcheckpoint success"
}

checkBalance(){
    local node=$1
    local url=$2
    local user=$3
    local amount=$4
    debug "prepare getBalance"
    balance=$($node attach $url --exec "platon.getBalance(\"$user\")")
    balance=$(printf "%.0f" "$balance")
    result=$(echo "$balance > $amount" | bc -l)
    if [ $result -eq 0 ]; then
        warn "insufficient user balance:$balance, require:$amount"
        exit 
    else
        debug "user has enough coin:$user balance:$balance"
    fi
}

checkTokenBalance(){
    local tools=$1
    local url=$2
    local user=$3
    local amount=$4
    balance=$($tools cast --module erc20 --address $token --rpc $url  --type call --method balanceOf $user)
    balance=$(printf "%.0f" "$balance")
    result=$(echo "$balance > $amount" | bc -l)
    if [ $result -eq 0 ]; then
        warn "insufficient user token balance:$balance, require:$amount"
        exit 
    else
        debug "user has enough token coin:$user balance:$balance"
    fi
}

checkNumber() {
    local str="$1"
    if [[ $str =~ ^-?[0-9]+$ ]]; then
        return 0
    else
        debug "check number failed, str:$str"
        exit
    fi
}