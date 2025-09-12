#!/bin/bash
envfile=
writeLog(){
	echo "$1=$2" >> $envfile
}
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color
LOG_LEVEL="DEBUG"
declare -A LOG_LEVELS=(
    ["DEBUG"]=0
    ["INFO"]=1
    ["WARN"]=2
    ["ERROR"]=3
    ["FATAL"]=4
)

setLogLevel() {
    local level="$1"
    if [[ -n "${LOG_LEVELS[$level]}" ]]; then
        LOG_LEVEL="$level"
        log_debug "set log level: $level"
    else
        log_error "invalid log level: $level"
    fi
}

shouldLog() {
    local level="$1"
    [[ "${LOG_LEVELS[$level]}" -ge "${LOG_LEVELS[$LOG_LEVEL]}" ]]
}
log() {
    local level="${1:-INFO}"
    local message="$2"
    local timestamp=$(date '+%Y-%m-%d %H:%M:%S')
    
    local line_number="${BASH_LINENO[1]}"
    local script_name=$(basename "${BASH_SOURCE[2]}")
    
    local color
    case "$level" in
        "DEBUG") color="$PURPLE" ;;
        "INFO") color="$GREEN" ;;
        "WARN") color="$YELLOW" ;;
        "ERROR") color="$RED" ;;
        "FATAL") color="$RED" ;;
        *) color="$NC" ;;
    esac
    
    if [[ "$level" == "ERROR" || "$level" == "FATAL" ]]; then
        echo -e "${color}[$timestamp] [$level] [$script_name:$line_number]${NC} $message" >&2
    else
        echo -e "${color}[$timestamp] [$level] [$script_name:$line_number]${NC} $message" >&2
    fi
}
debug() {
    if shouldLog "DEBUG"; then
        log "DEBUG" "$1"
    fi
}

info() {
    if shouldLog "INFO"; then
        log "INFO" "$1" 
    fi
}

warn() {
    if shouldLog "WARN"; then
        log "WARN" "$1"
    fi
}

error() {
    if shouldLog "ERROR"; then
        log "ERROR" "$1" 
    fi
}

fatal() {
    if shouldLog "FATAL"; then
        log "FATAL" "$1" 
    fi
}


