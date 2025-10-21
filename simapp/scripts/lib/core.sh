#!/bin/bash

loadLibraries() {
    local lib_dir="$(dirname "${BASH_SOURCE[0]}")"
    
    for lib in utils log; do
        if [[ -f "${lib_dir}/${lib}.sh" ]]; then
            source "${lib_dir}/${lib}.sh"
        else
            echo "WARN: library ${lib}.sh non-exists" >&2
        fi
    done
}