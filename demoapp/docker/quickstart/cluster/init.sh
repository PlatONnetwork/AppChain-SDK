#!/bin/bash

N=$1
node=node$N
if [ ! -d $node/data/sdk ];
then
  ./demoapp init --datadir $node/data ./genesis.json
fi
