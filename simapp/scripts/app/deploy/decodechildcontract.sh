#!/bin/bash

decodeChildcontractToml() {
    local config_file="$1"
    local -n config_map="$2" 
    
    if [[ ! -f "$config_file" ]]; then
        debug "error: Config file not found: $config_file" >&2
        return 1
    fi
    
    local target_fields=("token" "deposit_manager" "exit_helper" "stake_manager")
    
    while IFS= read -r line; do
        line=$(echo "$line" | sed 's/^[[:space:]]*//;s/[[:space:]]*$//' | sed 's/#.*$//')
        if [[ -z "$line" ]]; then
            continue
        fi

        if [[ "$line" =~ ^([a-zA-Z_][a-zA-Z0-9_]*)[[:space:]]*=[[:space:]]*\'([^\']+)\' ]]; then
            local key="${BASH_REMATCH[1]}"
            local value="${BASH_REMATCH[2]}"
            
            if [[ " ${target_fields[@]} " =~ " $key " ]]; then
                config_map["$key"]="$value"
                debug "found: $key = $value"
            fi
        fi
    done < "$config_file"
    
    debug "successfully parsed ${#config_map[@]} configuration items"
}

