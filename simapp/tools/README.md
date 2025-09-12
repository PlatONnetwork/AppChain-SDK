# Staking Script
Simple staking and unstaking script

!!!~~~~ Prior to using this script, verify that your node is fully synchronized with the latest block. Failure to do so may lead to penalties when executing stake operations.
# Execution Process

* Check the correctness of `config/env` parameters. The default issued configuration is already consistent with the node.
* Navigate to the `bin` directory.
* Execute `./staking.sh stake` to perform staking. If the balance is insufficient, please contact the official team.
* Execute `./staking.sh validators` to check if you have become a validator node. This usually takes some time.
* Execute `./staking.sh unstake` to perform unstaking. The node will exit the validator list.
* Execute `./staking.sh validators` to check if the node has exited the validator list. This usually takes some time.
