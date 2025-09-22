#!/bin/bash


SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_BASE="$(dirname "$SCRIPT_DIR")"
PROJECT_ROOT="$(dirname "$PROJECT_BASE")"
source "${PROJECT_ROOT}/lib/core.sh"
loadLibraries
debug "SCRIPT_DIR:$SCRIPT_DIR"
debug "PROJECT_ROOT:$PROJECT_ROOT"

output=

nodes=
source decodenodes.sh
source decodechildcontract.sh
source findnodedir.sh
source envtmpl.sh
source benchmarkenvtmpl.sh
source simappenvtmpl.sh
declare -A simappConfigMap
deployPrivateNetwork(){
    rootDir=$1
    simappDir=$2
    simappClient=$3
    tools=$4
    benchmarkclient=$5
    debug "$rootDir"

    if [[ -d $rootDir ]]; then
        rm -rf $rootDir
    fi
    debug "$rootDir"
    mkdir -p $rootDir
    mkdir -p $rootDir/bin
    mkdir -p $rootDir/scripts/bin
    mkdir -p $rootDir/scripts/config
    $(cp -f $simappDir $rootDir/bin/simapp)
    $(cp -f $simappClient $rootDir/bin/simappclient)
    $(cp -f $tools $rootDir/bin/tools)
    $(cp -f $benchmarkclient $rootDir/bin/benchmarkclient)
    $(cp -rf $PROJECT_ROOT/lib $rootDir/scripts/)
    $(cp -rf $PROJECT_BASE/tools/privatenetwork_en.md $rootDir/README.md)
    $(cp -rf $PROJECT_BASE/tools/bin/gensimapp.sh $rootDir/scripts/bin)
    $(cp -rf $PROJECT_BASE/tools/bin/benchmarkcluster.sh $rootDir/scripts/bin)
    genSimappEnv $rootDir/scripts/config/simappenv $simappConfigMap
    genBenchmarkEnv $rootDir/scripts/config/benchmarkenv $simappConfigMap
}

showHelp() {
    cat << EOF
Usage: $0 [OUTPUTDIR] [SIMAPPDIR] [SIMAPPCLIENT] [TOOLS] [BENCHMARKCLIENT]

A simple gen simapp script.
OUTPUTDIR       output dir
SIMAPPDIR       simapp binary path
SIMAPPCLIENT    simapp client binary path
TOOLS           tools binary path
BENCHMARKCLIENT benchmarkclient binary path
EOF
}

case "$1" in
    "all")
    deployPrivateNetwork $2 $3 $4 $5 $6
    ;;
esac