# Generate Configuration

*   `cd scripts/bin` to enter the script directory.
*   `./generate.sh --stakeamount 1000000000`

Generates the private keys required for the node. `1000000000` indicates the number of tokens the node will stake.

# Start Node

*   Switch to the project root directory.
*   `./init.sh`
*   `./start.sh`

# Staking Script
Simple staking and unstaking script.

!!!~~~~ Before using this script, please ensure your node is fully synchronized to the latest block. Otherwise, penalties may be incurred when executing staking operations.

## Execution Process

*   Check if the parameters in `config/env` are correct. The default issued configuration is already consistent with the node.
*   Enter the `bin` directory.
*   Execute `./staking.sh stake` to perform staking. If the balance is insufficient, please contact the official team.
*   Execute `./staking.sh validators` to check if you have become a validator node. This usually takes some time.
*   Execute `./staking.sh unstake` to perform unstaking. The node will exit the validator list.
*   Execute `./staking.sh validators` to check if the node has exited the validator list. This usually takes some time.

# Stress Testing

*   `cd scripts/bin`
*   `./benchmarkcluster.sh gencontracttx 200000` to generate 200,000 ERC20 contract transfer transactions.
*   `./benchmarkcluster.sh status` to check if generation is complete (e.g., cache:200000).
*   `./benchmarkcluster.sh startcontracttx` to start the stress test.
*   `./benchmarkcluster.sh status` to check if all transactions have been sent (e.g., sent:200000).

Block transaction details can be queried via the block explorer.