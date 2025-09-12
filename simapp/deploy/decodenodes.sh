#!/bin/bash



decodeNodesToml(){
    
CONFIG_FILE=$1
if [[ ! -f "$CONFIG_FILE" ]]; then
    echo "错误: 文件 $CONFIG_FILE 不存在"
    exit 1
fi
local -n node_map="$2"

local current_section=""
local node_count=0
local node_public_key=""
# node_address=""
local node_address=""
while IFS= read -r line; do
    line_clean=$(echo "$line" | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//')
    
    if [[ "$line_clean" == "[[nodes]]" ]]; then
        ((node_count++))
        current_section="nodes"
        continue
    fi
    
    if [[ "$current_section" == "nodes" ]]; then

        if [[ "$line_clean" =~ ^node_public_key[[:space:]]*=[[:space:]]*\'([^\']+)\' ]]; then
            node_public_key="${BASH_REMATCH[1]}"
        fi
        
        if [[ "$line_clean" =~ ^node_address[[:space:]]*=[[:space:]]*\'([^\']+)\' ]]; then
            node_address="${BASH_REMATCH[1]}"

        fi
        if [[ ${#node_public_key} -gt 0 && ${#node_address} -gt 0 ]]; then
                echo "node_public_key:$node_public_key,node_address:$node_address"

            node_map["$node_public_key"]="$node_address"
            node_public_key=""
            node_address=""
        fi

        if [[ "$line_clean" =~ ^\[[[:alnum:]_]+\] ]] || [[ -z "$line_clean" ]]; then
            current_section=""
        fi
    fi
    
done < "$CONFIG_FILE"

}

