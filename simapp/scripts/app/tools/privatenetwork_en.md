# Private Network Deployment

## Generate Configuration Parameters

```
cd scripts/bin
./gensimapp.sh gen [IPS] [USERNAME] [PASSWORD]
```

Assuming four nodes are to be installed: 10.1.1.33, 10.1.1.34, 10.1.1.35, 10.1.1.36, with the server account credentials uniformly set as simapp/123456, the command would be `./gensimapp gen 10.1.1.33,10.1.1.34,10.1.1.35,10.1.1.36 simapp 123456`. The following files will be generated in the `scripts/config` directory:

- **benchmarkenv**: Environment variable configuration for stress testing, including RPC addresses and performance parameters.
- **l1checkpointsender.json**: Keystore of the account used to send checkpoint transactions to Layer1. Ensure it has sufficient balance.
- **l1checkpointsender_password**: Password for the keystore in `l1checkpointsender.json`.
- **l2txsender.json**: Keystore of the account used to send system transactions on Layer2.
- **l2txsender_password**: Password for the keystore in `l2txsender.json`.
- **simappenv**: Environment variable configuration for application chain deployment.

## Create the Chain

```
source scripts/config/simappenv # Load application environment variables
cd bin # Switch to the application directory
./simappclient deploytemplate # Deploy contract templates on Layer1
./simappclient createnode # Generate node configurations
./simappclient deploychildchain # Deploy Layer2 contracts on Layer1
./simappclient creategenesis # Create the genesis file
./simappclient createansible # Create the ansible directory
./simappclient createansiblenode # Create ansible nodes
cd scripts/output/ansible # Switch to the ansible directory
ansible-playbook -i inventories/hosts.yml playbooks/deploy.yml # Deploy to target machines
ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=init" # Initialize nodes
ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=start" # Start nodes
```

## Stress Testing

```
cd scripts/bin
./benchmarkcluster.sh gencontracttx 200000 # Generate 200,000 ERC20 contract transfer transactions
./benchmarkcluster.sh status # Check if cache:200000 is fully generated
./benchmarkcluster.sh startcontracttx # Start stress testing
./benchmarkcluster.sh report [startblock] [endblock] report.csv # After stress testing, use the report command to generate block data for analysis
```