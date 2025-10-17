## PlatON Application Chain Test Environment

PlatON has implemented an L2 application chain based on PlatON, utilizing the same consensus algorithm, virtual machine, and economic model as PlatON. The source code for this project is located at xxxxx.

### Test Application Chain

PlatON has deployed a set of PlatON L2 application chains for testing purposes. The test environment consists of 4 nodes with IP addresses: 8.219.210.248, 8.222.253.224, 8.222.207.89, and 8.219.242.242. The block explorer is available at [http://8.219.177.143:8000/](http://8.219.177.143:8000/ "http://8.219.177.143:8000/").

Users can conduct contract and dAPP development testing in this test environment, or set up their own nodes to join the test chain and perform performance verification. The methods are as follows:

#### Download Test Environment Configuration
The test environment configuration system is Ubuntu 24.04 LTS. Please ensure consistency. The download link is:
https://app-chain.oss-ap-southeast-1.aliyuncs.com/testnetnode.tar.gz
#### Generate Configuration

* `cd scripts/bin` to enter the script directory
* `./generate.sh --stakeamount 1000000000`

Generates the private keys required for the node. 1000000000 indicates the number of tokens the node will stake.

#### Test Tokens

Please contact us to obtain test tokens.

#### Start Node

* Switch to the project root directory
* `./init.sh`
* `./start.sh`

#### Staking

#### Node Staking

Before performing the staking operation, please ensure your node is fully synchronized to the latest block. Otherwise, penalties may be imposed when executing the staking operation.

##### Execution Process

* Check if the parameters in `config/env` are correct. The default issued configuration is already consistent with the node.
* Enter the `bin` directory.
* Execute `./staking.sh stake` to perform staking. If the balance is insufficient, please contact the official team.
* Execute `./staking.sh validators` to check if you have become a validator node. This usually takes some time.
* Execute `./staking.sh unstake` to perform unstaking. The node will exit the validator list.
* Execute `./staking.sh validators` to check if the node has exited the validator list. This usually takes some time.

#### Performance Testing

* `cd scripts/bin`
* `./benchmark.sh gencontracttx 200000` to generate 200,000 ERC20 contract transfer transactions
* `./benchmark.sh status` to check if generation is complete (cache:200000)
* `./benchmark.sh startcontracttx` to start the stress test
* `./benchmark.sh status` to check if all transactions have been sent (sent:200000)

Block transaction details can be queried via the block explorer.

### Private Application Chain Deployment

Users can build their own PlatON L2 application chain based on this source code. Below provides the setup using a compiled private chain configuration environment:

#### Download Version
The test environment configuration system is Ubuntu 24.04 LTS. Please ensure consistency. The download link is:
https://app-chain.oss-ap-southeast-1.aliyuncs.com/privatenetwork.tar.gz

#### Generate Configuration Parameters

- `cd scripts/bin` to enter the script directory
- `./gensimapp.sh [IPS] [USERNAME] [PASSWORD]`, for example, if the machines are 10.1.1.33, 10.1.1.34, 10.1.1.35, and the server account password is simapp, 123456, then the command is `./gensimapp 10.1.1.33,10.1.1.34,10.1.1.35 simapp 123456`. The `scripts/config` directory will generate the following files: `benchmarkenv`, `l1checkpointsender.json`, `l1checkpointsender_password`, `l2txsender.json`, `l2txsender_password`, `simappenv`
    - `benchmarkenv`: Environment variable configuration for stress testing, mainly RPC addresses and performance parameter configurations
    - `l1checkpointsender.json`, `l1checkpointsender_password`: Account for sending checkpoint transactions to Layer1, must have sufficient balance
    - `l2txsender.json`, `l2txsender_password`: Account for sending system transactions on Layer2
    - `simappenv`: Environment variable configuration for app deployment

#### Create Chain

- `cd bin` to switch to the application directory
- `source scripts/config/simappenv` to apply environment variables
- `./simappclient deploytemplate` to deploy contract templates on Layer1
- `./simappclient createnode` to generate node configuration
- `./simappclient deploychildchain` to deploy Layer2 contracts on Layer1
- `./simappclient creategenesis` to create genesis file
- `./simappclient createansible` to create ansible directory
- `./simappclient createansiblenode` to create ansible node
- `cd scripts/output/ansible` to switch to the ansible directory
- `ansible-playbook -i inventories/hosts.yml playbooks/deploy.yml` to deploy to target machines
- `ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=init"` to initialize nodes
- `ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=start"` to start nodes

#### Performance Testing

- `cd scripts/bin`
- `./benchmarkcluster.sh gencontracttx 200000` to generate 200,000 ERC20 contract transfer transactions
- `./benchmarkcluster.sh status` to check if generation is complete (cache:200000)
- `./benchmarkcluster.sh startcontracttx` to start the stress test
- `./benchmarkcluster.sh report [startblock] [endblock] report.csv` after the stress test ends, use the report command to generate stress test block data