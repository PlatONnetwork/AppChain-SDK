#!/bin/bash

getNodeFolders() {
    local target_dir="${1:-.}"
    local -n result_array="$2" 
    result_array=()
    
    while IFS= read -r -d '' folder; do
        folder_name=$(basename "$folder")
        if [[ "$folder_name" =~ ^node[0-9]+$ ]]; then
            result_array+=("$folder_name")
        fi
    done < <(find "$target_dir" -maxdepth 1 -type d -name "node[0-9]*" -print0 2>/dev/null)
    
    IFS=$'\n' result_array=($(sort -V <<<"${result_array[*]}"))
    unset IFS
    
    return ${#result_array[@]} 
}