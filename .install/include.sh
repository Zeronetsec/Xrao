function include() {
    if [[ "${1}" != ":" ]]; then
        return 1
    fi

    shift
    local input_data="${1}"
    if [[ -z "${input_data}" ]]; then
        return 1
    fi

    local line
    local inside_bracket=0

    while read -r line; do
        line="${line#"${line%%[![:space:]]*}"}"
        line="${line%"${line##*[![:space:]]}"}"
        if [[ "${line}" == *"("* ]]; then
            inside_bracket=1
            continue
        fi

        if [[ "${line}" == *")"* ]]; then
            inside_bracket=0
            continue
        fi

        if (( inside_bracket )); then
            [[ -z "${line}" ]] && continue
            [[ "${line}" =~ ^# ]] && continue
            if [[ ! -f "${root}/${line}.sh" ]]; then
                echo -e "\x1b[1;31m[!] \x1b[0mInclude: \x1b[0;32m${line} \x1b[0mnot found!"
                return 1
            fi
            source "${root}/${line}.sh"
        fi
    done <<< "${input_data}"
    return 0
}; readonly -f include