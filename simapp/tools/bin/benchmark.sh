#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
source "${PROJECT_ROOT}/lib/core.sh"
loadLibraries
source "${PROJECT_ROOT}/config/benchmarkenv"


benchmarkclient=./benchmarkclient

showHelp(){

    cat << EOF
Usage: $0 COMMAND [ARGS]

A simple benchmark operation script.

Commands:
  genrawtx          The node generates stress test transactions
  gencontracttx     The node generates stress test transactions
  startrawtx        Start transaction send
  startcontracttx   Start transaction send
  stop              Stop transaction send
  status            Test status
EOF

}

gencontracttx(){
    count=$1
    if [[ -z $count ]];then
        count=$BENCHMARK_COUNT
    fi
    $benchmarkclient genservertx --url $BENCHMARK_URL --startaddr $BENCHMARK_START_ADDR --endaddr $BENCHMARK_END_ADDR --rawtx $BENCHMARK_CONTRACT_RAWTX --count $count
}

genrawtx(){
    count=$1
    if [[ -z $count ]];then
        count=$BENCHMARK_COUNT
    fi
    $benchmarkclient genservertx --url $BENCHMARK_URL --startaddr $BENCHMARK_START_ADDR --endaddr $BENCHMARK_END_ADDR --rawtx $BENCHMARK_RAWTX --count $count

}

startcontracttx(){
    $benchmarkclient start --urls $BENCHMARK_URL --tps $BENCHMARK_CONTRACT_TPS --txsperaccount $BENCHMARK_CONTRACT_TXSPERACCOUNT
}
startrawtx(){
    $benchmarkclient start --urls $BENCHMARK_URL --tps $BENCHMARK_RAW_TPS --txsperaccount $BENCHMARK_RAW_TXSPERACCOUNT
}

stop(){
    $benchmarkclient stop --urls $BENCHMARK_URL
}
status(){
    $benchmarkclient status --urls $BENCHMARK_URL
}
case $1 in
    "genrawtx")
        genrawtx $2
    ;;
    "gencontracttx")
        gencontracttx $2
    ;;
    "startrawtx")
        startrawtx
    ;;
    "startcontracttx")
        startcontracttx
    ;;
    "stop")
        stop
    ;;
    "status")
        status
    ;;
    "help"|"--help")
        showHelp
    ;;
esac