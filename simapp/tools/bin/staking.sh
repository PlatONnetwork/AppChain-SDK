#!/bin/bash
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
source "${PROJECT_ROOT}/lib/core.sh"
loadLibraries
source "${PROJECT_ROOT}/config/env"

node=../../node
tools=./tools
approveAmount=200000000000000000000000

# ENV_L1_STAKING_HASH=
# ENV_L1_STAKING_ID=
# ENV_L2_STAKING_WITHDRAW_HASH=
# ENV_L2_UNSTAKE_HASH=
# ENV_L2_STAKING_WITHDRAW_ID=

getStakeAmount(){
    amount=$($tools cast --module l1.staking --address $l1staking --rpc $rootchainurl --type call --method stakeOf $validatoraddr)
    checkNumber $amount
    echo $amount
}

approve(){
    debug "prepare l1 approve"
    receiptLog=$($tools cast --module erc20 --address $token --rpc $rootchainurl --key $userkey --type send --method approve $l1staking $approveAmount)
    hash=$(getReceiptHash "$receiptLog")
    checkTxHash $hash
	debug "l1 approve hash:$hash"
	checkTx $node $rootchainurl $hash
}
l1Stake(){
    debug "check $user balance"
    checkBalance $node $rootchainurl $user 0
    checkTokenBalance $tools $rootchainurl $user 0
    amount=$(getStakeAmount)
    checkNumber $amount
    if [[ $amount -gt 0 ]];then
        warn "had staked, validator:$validatoraddr, amount:$amount"
        exit
    fi
    approve
    debug "prepare l1 staking:$validatoraddr"
    local receiptLog=$($tools cast --module l1.staking --address $l1staking --rpc $rootchainurl --key $userkey --type send --method stakeFor $validator)
    debug $receiptLog
    hash=$(getReceiptHash "$receiptLog")
    checkTxHash $hash
	debug "l1 staking hash:$hash"
	checkTx $node $rootchainurl $hash

	ENV_L1_STAKING_HASH=$hash
    local idhex=$($node attach $rootchainurl --exec "platon.getTransactionReceipt(\"$hash\").logs[3].topics[1]")
	ENV_L1_STAKING_ID=$(hexToInt "$idhex")
	debug "l1 staking id:$ENV_L1_STAKING_ID"
}

waitL2Statesync(){
	debug "prepare wait l2 state sync"
	equalId(){
		result=$1
		if [ $result -eq $ENV_L1_STAKING_ID ];then
			echo 1
		else
			echo 0
		fi
	}
	waitCondition "$tools cast --module l2.statesync --address $l2statesync --rpc $childchainurl --type call --method getExecutedId"  equalId "l2 staking sync"
}

getNewestValidator(){
    round=$($node attach $childchainurl --exec "debug.consensusStatus().state.view.epoch")
    validators=$($tools cast --module l2.staking --address $l2staking --rpc $childchainurl --type call --method getValidatorAddrs 1 $round)
    echo $validators
}


waitBecomeL2Validator(){
    debug "prepare wait become l2 validator"
    equalId(){
        validators=$1
		result=$(echo "${validators[*]}" | tr '[:upper:]' '[:lower:]' | tr -d ' ')
        lowerVal=$(echo "$validatoraddr" | tr '[:upper:]' '[:lower:]')
		if [[ "$result" == *"$lowerVal"* ]];then
			echo 1
		else
			echo 0
		fi
	}
    waitCondition getNewestValidator equalId "wait become l2 validator"
}



l2Unstake(){
    checkBalance $node $childchainurl $user 0
    unamount=$1
    if [[ unamount -eq 0 ]];then
        amount=$(getStakeAmount)
        checkNumber $amount
        if [[ $amount -gt 0 ]];then
            debug "had staked, validator:$validatoraddr, amount:$amount"
            unamount=$amount
        fi
    fi
    debug "prepare unstake:$validatoraddr $unamount"
    local receiptLog=$($tools cast --module l2.staking --address $l2staking --rpc $childchainurl --key $userkey --type send --method unstake $validatoraddr $unamount)
    hash=$(getReceiptHash "$receiptLog")
    checkTxHash $hash
	debug "l2 unstake hash:$hash"
	checkTx $node $childchainurl $hash
    ENV_L2_UNSTAKE_HASH=$hash
}


checkL2Withdrable(){
    pendingWithdrawbleBalance=$($tools cast --module l2.staking --address $l2staking --rpc $childchainurl  --type call --method pendingWithdrawalsOfStake $validatoraddr)
    withdrawbleBalance=$($tools cast --module l2.staking --address $l2staking --rpc $childchainurl  --type call --method withdrawableOfStake $validatoraddr)
    if [[ $pendingWithdrawbleBalance -eq 0 && $withdrawbleBalance -eq 0 ]];then
        warn "available withdraw is zero,exit"
        exit 1
    fi
    waitL2Withdrable
}

waitL2Withdrable(){
    debug "prepare wait l2 withdraw"
	equalValue(){
		value=$1
        checkNumber $value
		if [ $value -gt 0 ]; then
			echo 1
		else
			echo 0
		fi
	}
	waitCondition "$tools cast --module l2.staking --address $l2staking --rpc $childchainurl  --type call --method withdrawableOfStake $validatoraddr" equalValue "withdrawable staking"
}

withdrawStaking(){
    local receiptLog=$($tools cast --module l2.staking --address $l2staking --rpc $childchainurl --key $userkey --type send --method withdrawUnstake $validatoraddr)
    hash=$(getReceiptHash "$receiptLog")
    checkTxHash $hash
	debug "withdraw unstake hash:$hash"
	checkTx $node $childchainurl $hash
	ENV_L2_STAKING_WITHDRAW_HASH=$hash
	local idhex=$($node attach $childchainurl --exec "platon.getTransactionReceipt(\"$hash\").logs[1].topics[1]")
	ENV_L2_STAKING_WITHDRAW_ID=$(hexToInt "$idhex")
	debug "withdraw unstaking id:$ENV_L2_STAKING_WITHDRAW_ID"
}

waitUnstakeCheckpoint(){
    debug "wait unstake checkpoint"
    waitcheckpoint $ENV_L2_STAKING_WITHDRAW_ID
}

l1exithelper(){
	debug "prepare l1 exithelper"
	receiptLog=$($tools cast l1.exithelper --address $l1exithelper --rootchainrpc $rootchainurl --rpc $childchainurl --key $userkey --exit-id $ENV_L2_STAKING_WITHDRAW_ID)
    hash=$(getReceiptHash "$receiptLog")
    checkTxHash $hash
	debug "l1 exithelper hash:$hash"
	checkTx $node $rootchainurl $hash
}

checkL1Withdraw(){
    withdrawableAmount=$($tools cast --module l1.staking --address $l1staking --rpc $rootchainurl --key $userkey --type call --method withdrawableStake $validatoraddr)
    checkNumber $withdrawableAmount
    if [[ $withdrawableAmount -gt 0 ]]; then
        unamount=$withdrawableAmount
        debug "had withdrawable balance:$unamount"
        l1withdrawunstake
    else
        warn "withdrawable balance is zero, quit"
        exit
    fi
}

l1withdrawunstake(){
	debug "prepare l1 withdraw unstake"
 	receiptLog=$($tools cast --module l1.staking --address $l1staking --rpc $rootchainurl --key $userkey --type send --method withdrawStake $validatoraddr $user  $unamount)
	hash=$(getReceiptHash "$receiptLog")
    checkTxHash $hash
	debug "withdraw unstake hash:$hash"
	checkTx $node $rootchainurl $hash
}


l1querybalance(){
	debug "prepare l1 query $user balance "
	balance=$($tools cast --module erc20 --address $token --rpc $rootchainurl --type call --method balanceOf $user)
    debug "balance:$balance"
}


stakeFlow(){
    l1Stake
    waitL2Statesync
}
unstakeFlow(){
    l2Unstake
}
withdrawUnstakeFlow(){
    checkL2Withdrable
    withdrawStaking
    waitUnstakeCheckpoint
    l1exithelper
    checkL1Withdraw
    l1querybalance
}
runFunc() {  
    if [[ $# -eq 0 ]]; then  
        warn "error: runfunc requires at least one function name parameter"  
        warn "usage: \$0 runfunc <function1> [function2] ..."  
        list_functions  
        return 1  
    fi  

    local exit_code=0  
    local functions=("$@")  
    
    debug "starting execution of ${#functions[@]} functions..."  

    for func in "${functions[@]}"; do  
        # Check if function exists  
        if [[ $(type -t "$func") != "function" ]]; then  
            warn "Error: Function '$func' is not defined"  
            list_functions  
            return 1  
        fi  

        # Execute function  
        debug "executing: $func"  
        if $func; then  
            debug "success: $func"  
        else  
            local func_exit_code=$?  
            debug "failed: $func (exit code: $func_exit_code)"  
            exit_code=$func_exit_code  
            break  # Stop execution on failure  
        fi  
        debug "----------------------------------------"  
    done  

    if [[ $exit_code -eq 0 ]]; then  
        debug "all functions executed successfully!"  
    else  
        debug "execution interrupted due to error"  
    fi  
    
    return $exit_code  
}  


showAdvancedHelp(){
    if [[ "$ADVANCED_COMMAND" == "true" ]];then
   cat << EOF
Advanced commands:
  withdrawUnstake      Withdraw unstake amount
  runfunc              Call function list
EOF
    fi
}

advancedCommand(){
    if [[ "$ADVANCED_COMMAND" == "true" ]];then
        case $1 in
            "withdrawUnstake")
                withdrawUnstakeFlow
                return
                ;;
            "runfunc")
                shift
                runFunc "$@"
                return
            ;;
        esac
    fi
    echo "unknow command, type "$0 help" to see help"
}

showHelp() {
    cat << EOF
Usage: $0 COMMAND [ARGS]

A simple staking operation script.

Commands:
  stake      Stake to become a validator
  unstake    Unstake cancel validator
  validators Query the current validators
EOF
}
case $1 in
    "stake")
        stakeFlow
        ;;
    "unstake")
        unstakeFlow
        ;;
    "validators")
        getNewestValidator
        ;;
    "--help"|"help")
        showHelp
        showAdvancedHelp
        ;;
    *)
        advancedCommand "$@"
    ;;
esac
