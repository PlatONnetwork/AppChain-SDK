#!/bin/bash

N=$1
node=node$N
if [ ! -d $node/data/sdk ];
then
  simapp init --datadir $node/data ./genesis.json
fi