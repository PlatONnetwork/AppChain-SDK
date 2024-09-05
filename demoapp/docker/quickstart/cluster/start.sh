#!/bin/bash

N=$1
DATADIR=node$N
P2P_PORT=1800$N
WS_PORT=660$N
HTTP_PORT=880$N
PPROF_PORT=700$N

./demoapp --datadir $DATADIR/data --nodekey $DATADIR/data/nodekey --cbft.blskey $DATADIR/data/blskey --port $P2P_PORT --verbosity 4 --ws --ws.addr 0.0.0.0 --ws.port $WS_PORT --ws.api platon,debug,personal,admin,net,web3,txpool --http --http.vhosts "*" --http.corsdomain "*" --http.addr 0.0.0.0 --http.port $HTTP_PORT --http.api platon,debug,personal,admin,net,web3,txpool,checkpoint --pprof --pprof.addr 0.0.0.0 --pprof.port $PPROF_PORT --cache 256 --metrics --ipcdisable --txpool.locals lat1wrfq0sfj9n9eq6wn0yxkw6yxdk4l7yp4hpmt8p,lat1xsuh90mr6yrzwcd242y36f6s7q7tfvhh5844xy --maxpeers 100 --maxconsensuspeers 75 --txpool.globaltxcount 1000 --nodiscover  --networkid 102 --allow-insecure-unlock --statesync.startblock 1 --l2.keystore ./75a8f9b1-73cc-4d57-952d-673efe1efee2 --l2.password ./password
