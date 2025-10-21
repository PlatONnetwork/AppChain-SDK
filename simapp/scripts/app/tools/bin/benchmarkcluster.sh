#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
source "${PROJECT_ROOT}/lib/core.sh"
loadLibraries
source "${PROJECT_ROOT}/config/benchmarkenv"


benchmarkclient=../../bin/benchmarkclient

showHelp(){
    cat << EOF
Usage: $0 COMMAND [ARGS]

A simple benchmark operation script.

Commands:
  genrawtx [COUNT]          Generate stress test raw transactions
                            COUNT: Number of transactions (default: value from config file BENCHMARK_COUNT)
  
  gencontracttx [COUNT]     Generate stress test contract transactions
                            COUNT: Number of transactions (default: value from config file BENCHMARK_COUNT)
  
  startrawtx                Start raw transaction sending
  
  startcontracttx           Start contract transaction sending
  
  stop                      Stop transaction sending
  
  report START_BLOCK END_BLOCK OUTPUT_FILE
                            Generate TPS report
                            START_BLOCK: Starting block height
                            END_BLOCK: Ending block height
                            OUTPUT_FILE: Output file path
  
  status                    Show test status

Examples:
  $0 genrawtx 1000          Generate 1000 raw transactions
  $0 gencontracttx          Generate contract transactions (using default count)
  $0 report 100 200 report.txt Generate report from block 100 to 200
EOF
}

gencontracttx(){
    count=$1
    if [[ -z $count ]];then
        count=$BENCHMARK_COUNT
    fi
    $benchmarkclient gentx --urls $BENCHMARK_URLS --addrs $BENCHMARK_ADDRS --rawtx $BENCHMARK_CONTRACT_RAWTX --count $count
}

genrawtx(){
    count=$1
    if [[ -z $count ]];then
        count=$BENCHMARK_COUNT
    fi
    $benchmarkclient gentx --urls $BENCHMARK_URLS --addrs $BENCHMARK_ADDRS --rawtx $BENCHMARK_RAWTX --count $count

}

startcontracttx(){
    $benchmarkclient start --urls $BENCHMARK_URLS --tps $BENCHMARK_CONTRACT_TPS --txsperaccount $BENCHMARK_CONTRACT_TXSPERACCOUNT
}
startrawtx(){
    $benchmarkclient start --urls $BENCHMARK_URLS --tps $BENCHMARK_RAW_TPS --txsperaccount $BENCHMARK_RAW_TXSPERACCOUNT
}

stop(){
    $benchmarkclient stop --urls $BENCHMARK_URLS
}
status(){
    $benchmarkclient status --urls $BENCHMARK_URLS
}
report(){
    local start=$1
    local end=$2
    local output=$3
    $benchmarkclient report --urls $BENCHMARK_URLS --start $start --end $end --output $output
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
    "report")
        report $2 $3 $4
    ;;
    "help"|"--help")
        showHelp
    ;;
esac