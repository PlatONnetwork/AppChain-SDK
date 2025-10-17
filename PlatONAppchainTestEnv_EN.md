## PlatON Application Chain Test Environment

PlatON has implemented an L2 application chain based on PlatON, adopting the same consensus algorithm, virtual machine, and economic model as PlatON. The source code for this project is located at https://github.com/PlatONnetwork/AppChain-SDK.

### Testing the Application Chain

PlatON has deployed a set of PlatON L2 application chains for testing. The test environment consists of 4 nodes with IP addresses 8.219.210.248, 8.222.253.224, 8.222.207.89, and 8.219.242.242. The browser address is [http://8.219.177.143:8000/](http://8.219.177.143:8000/ "http://8.219.177.143:8000/").

Users can develop and test contracts and dAPPs in the test environment, or they can set up their own nodes to join the test chain and perform performance verification. The method is as follows:

#### Download Test Environment Configuration
The operating system is Ubuntu 24.04 LTS. The download link is:
https://app-chain.oss-ap-southeast-1.aliyuncs.com/testnetnode.tar.gz

#### Generate Configuration

* `cd scripts/bin` Navigate to the script directory.
* `./generate.sh  --stakeamount 1000000000`

Generate the private key required for the node. 1000000000 represents the number of tokens the node will stake.

#### Test Token

After executing the `generate.sh` script, the following information will be displayed. Please provide the addresses corresponding to `user address` and `checkpoint address` to us to obtain test tokens.
```
generate keys success please contact the official team to ensure sufficient balance for both the user address:0x6f9f59826097D25f0D6F4d719805D1cF8ba7a07d and checkpoint address:0x1Df85FA030d2E00bA28EfA7115eF5dB88B14CcDD
```

* **user address:** Validator account address, requires obtaining PlatON testnet LAT and application chain test tokens (SIMAPP).
* **checkpoint address:** Account address for sending checkpoint transactions, requires obtaining PlatON testnet LAT.

#### Start Node

* Switch to the project root directory.
* `./init.sh`
* `./start.sh`

#### Node Staking

Before performing the staking operation, please ensure your node is fully synchronized to the latest block. Otherwise, penalties may be incurred during the staking operation.

##### Execution Process

* Check if the `config/env` parameters are correct. The default configuration issued is consistent with the node.
* Enter the `bin` directory.
* Execute `./staking.sh stake` to perform staking. If the balance is insufficient, please contact the official team.
* Execute `./staking.sh validators` to check if it has become a validator node. This usually takes some time.
* Execute `./staking.sh unstake` to unstake. The node will exit the validator list.
* Execute `./staking.sh validators` to check if the node has exited the validator list. This usually takes some time.

#### Performance Testing

* `cd scripts/bin`
* `./benchmark.sh gencontracttx 200000` Generate 200000 ERC20 contract transfer transactions.
* `./benchmark.sh status` Check if cache:200000 has been generated.
* `./benchmark.sh startcontracttx` Start the stress test.
* `./benchmark.sh status` Check if sent:200000 has been completed.

The block transaction status can be queried through the browser.

### Private Application Chain Deployment

Users can build their own PlatON L2 application chain based on this source code. The following provides a setup using a compiled private chain configuration environment:

#### Download Version
The operating system is Ubuntu 24.04 LTS (install ansible). The download link is:
https://app-chain.oss-ap-southeast-1.aliyuncs.com/privatenetwork.tar.gz

#### Generate Configuration Parameters

- `cd scripts/bin` Navigate to the script directory.
- `./gensimapp.sh [IPS] [USERNAME] [PASSWORD]`. For example, if the machines are 10.1.1.33, 10.1.1.34, 10.1.1.35, and the server account password is simapp, 123456, then the command is `./gensimapp.sh gen 10.1.1.33,10.1.1.34,10.1.1.35 simapp 123456`. The following files will be generated in `scripts/config`:
  - **benchmarkenv:** Environment variable configuration for stress testing, mainly RPC addresses and performance parameter configuration.
  - **l1checkpointsender.json:** Keystore account file for sending checkpoint transactions to Layer1. Ensure it has sufficient balance.
  - **l1checkpointsender_password:** Keystore password for `l1checkpointsender.json`.
  - **l2txsender.json:** Account for sending system transactions on Layer2.
  - **l2txsender_password:** Keystore password for `l2txsender.json`.
  - **simappenv:** Environment variable configuration for application chain deployment.

#### Test Tokens

After executing the `gensimapp.sh` script, the following information will be displayed. Please ensure the addresses corresponding to `deploy address`, `initialize address`, and `checkpoint address` have sufficient tokens.
```
generate keys success please contact the official team to ensure sufficient balance for both the deploy address:0x311F5F59B1170e095709eE8da4f6880d912946F3, initialize address:0x8CFF458f3d7cFF30943701979b336adFf42D2776 and checkpoint address:0xabEe1cCfb84853d954f8a005a1Becc18d37A86ec
```
* **deploy address:** Contract deployment account address, requires PlatON testnet LAT.
* **initialize address:** Contract initialization account address, requires PlatON testnet LAT.
* **checkpoint address:** Account address for sending checkpoint transactions, requires PlatON testnet LAT.

#### Create Chain

- `source scripts/config/simappenv` Apply environment variables.
- `cd bin` Switch to the application directory.
- `./simappclient deploytemplate` Deploy contract template on Layer1.
- `./simappclient createnode` Generate node configuration.
- `./simappclient deploychildchain` Deploy Layer2 contract on Layer1.
- `./simappclient creategenesis` Create genesis file.
- `./simappclient createansible` Create ansible directory.
- `./simappclient createansiblenode` Create ansible node.
- `cd scripts/output/ansible` Switch to the ansible directory.
- `ansible-playbook -i inventories/hosts.yml playbooks/deploy.yml` Deploy to target machines.
- `ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=init"` Initialize nodes.
- `ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=start"` Start nodes.

#### Performance Testing

- `cd scripts/bin`
- `./benchmarkcluster.sh gencontracttx 200000` Generate 200000 ERC20 contract transfer transactions.
- `./benchmarkcluster.sh status` Check if cache:200000 has been generated.
- `./benchmarkcluster.sh startcontracttx` Start the stress test.
- `./benchmarkcluster.sh report [startblock] [endblock] report.csv` After the stress test ends, use the report command to generate stress test block data.