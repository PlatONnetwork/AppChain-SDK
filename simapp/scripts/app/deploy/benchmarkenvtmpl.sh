#!/bin/bash
declare -A BENCHMARK_ENV_TEMPLATE=(
    ["BENCHMARK_URLS"]=""
    ["BENCHMARK_URL"]="http://127.0.0.1:8801"
    ["BENCHMARK_ADDRS"]="280"
    ["BENCHMARK_START_ADDR"]="1719"
    ["BENCHMARK_END_ADDR"]="1999"
    ["BENCHMARK_RAWTX"]="100"
    ["BENCHMARK_CONTRACT_RAWTX"]="0"
    ["BENCHMARK_COUNT"]="5000000"
    ["BENCHMARK_RAW_TPS"]="25000"
    ["BENCHMARK_CONTRACT_TPS"]="12000"
    ["BENCHMARK_RAW_TXSPERACCOUNT"]="100"
    ["BENCHMARK_CONTRACT_TXSPERACCOUNT"]="50"

)

BENCHMARK_ENV_TEMPLATE_KEY=(
    BENCHMARK_URLS
    BENCHMARK_URL
    BENCHMARK_ADDRS
    BENCHMARK_START_ADDR
    BENCHMARK_END_ADDR
    BENCHMARK_RAWTX
    BENCHMARK_CONTRACT_RAWTX
    BENCHMARK_COUNT
    BENCHMARK_RAW_TPS
    BENCHMARK_CONTRACT_TPS
    BENCHMARK_RAW_TXSPERACCOUNT
    BENCHMARK_CONTRACT_TXSPERACCOUNT
)
genBenchmarkEnv(){
    file=$1
    local overrides="$2" 
    {
        echo "#!/bin/bash"
        echo "#"
        echo "# $comment"
        echo "# Generated on: $(date)"
        echo "#"
        echo ""
        
        for key in "${BENCHMARK_ENV_TEMPLATE_KEY[@]}"; do
            local value="${overrides[$key]:-${BENCHMARK_ENV_TEMPLATE[$key]}}"
            echo "export $key='$value'"
        done
    } > $file
}