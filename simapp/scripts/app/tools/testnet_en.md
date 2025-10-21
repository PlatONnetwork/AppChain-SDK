# Generate Configuration

* `cd scripts/bin` Navigate to the script directory
* `./generate.sh --stakeamount 1000000000`

Generate the private keys required for the node. 1000000000 represents the amount of tokens to be staked by this node.

# Start Node

* Switch to the project root directory
* `./init.sh`
* `./start.sh`

# Node Staking
> Before executing the following scripts, please ensure your node has fully synchronized to the latest block. Otherwise, penalties may be incurred during staking operations.
> Check if the parameters in `config/env` are correct. The default delivered configuration is already consistent with the node.

Node staking operation is as follows:

```
cd scripts/bin
./staking.sh stake # Perform staking. If the balance is insufficient, please contact the official team.
./staking.sh validators # Check if the node has become a validator. This usually takes some time.
```

Node unstaking operation is as follows:

```
cd scripts/bin
./staking.sh unstake # Perform unstaking. The node will exit the validator list.
./staking.sh validators # Check if the node has exited the validator list. This usually takes some time.
```

# Stress Testing

* `cd scripts/bin`
* `./benchmark.sh gencontracttx 200000` Generate 200,000 ERC20 contract transfer transactions
* `./benchmark.sh status` Check if cache:200000 has been fully generated
* `./benchmark.sh startcontracttx` Start stress testing
* `./benchmark.sh status` Check if sent:200000 has been fully transmitted

Block transaction status can be queried through the browser.