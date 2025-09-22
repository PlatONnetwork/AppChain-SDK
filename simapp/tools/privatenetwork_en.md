# Private Network Deployment

## Generate Configuration Parameters

*   `cd scripts/bin` Navigate to the script directory.
*   `./gensimapp.sh [IPS] [USERNAME] [PASSWORD]` For example, if the machine IPs are 10.1.1.33, 10.1.1.34, 10.1.1.35, and the server account password is `simapp,123456`, the command would be: `./gensimapp.sh 10.1.1.33,10.1.1.34,10.1.1.35 simapp 123456`. This will generate the files `benchmarkenv`, `l1checkpointsender.json`, `l1checkpointsender_password`, `l2txsender.json`, `l2txsender_password`, and `simappenv` in the `scripts/config/` directory.
    *   `benchmarkenv`: Environment variable configuration for benchmarking, mainly RPC addresses and performance parameter settings.
    *   `l1checkpointsender.json`, `l1checkpointsender_password`: Account for sending checkpoint transactions to Layer1. Ensure it has sufficient balance.
    *   `l2txsender.json`, `l2txsender_password`: Account for sending system transactions on Layer2.
    *   `simappenv`: Environment variable configuration for application deployment.

## Create Chain

*   `cd bin` Switch to the application directory.
*   `source ../scripts/config/simappenv` Source the application environment variables.
*   `./simappclient deploytemplate` Deploy contract templates on Layer1.
*   `./simappclient createnode` Generate node configuration.
*   `./simappclient deploychildchain` Deploy the Layer2 contract on Layer1.
*   `./simappclient creategenesis` Create the genesis file.
*   `./simappclient createansible` Create the Ansible directory structure.
*   `./simappclient createansiblenode` Create Ansible node configurations.
*   `cd ../scripts/output/ansible` Switch to the Ansible directory.
*   `ansible-playbook -i inventories/hosts.yml playbooks/deploy.yml` Deploy to target machines.
*   `ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=init"` Initialize nodes.
*   `ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=start"` Start nodes.

## Benchmarking

*   `cd scripts/bin`
*   `./benchmarkcluster.sh gencontracttx 200000` Generate 200,000 ERC20 contract transfer transactions.
*   `./benchmarkcluster.sh status` Check if generation is complete (e.g., `cache:200000`).
*   `./benchmarkcluster.sh startcontracttx` Start the benchmark.
*   `./benchmarkcluster.sh report [startblock] [endblock] report.csv` After the benchmark ends, use the `report` command to generate block data statistics into `report.csv`.