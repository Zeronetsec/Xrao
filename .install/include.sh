function include() {
    local input_data
    local line
    local ecode

    if [[ "${1}" != ":" ]]; then
        return 1
    fi

    shift
    input_data="${1}"

    if [[ -z "${input_data}" ]]; then
        return 1
    fi

    input_data="$(
        echo -e "${input_data}" | \
        command sed -n '/(/,/)/ {
            s/[()]//g;
            s/^[[:space:]]*//;
            s/[[:space:]]*$//;
            /^$/d;
            p;
        }'
    )"

    while read -r line; do
        [[ -z "${line}" ]] && continue
        [[ "${line}" =~ ^# ]] && continue
        if [[ ! -f "${root}/${line}.sh" ]]; then
            echo -e "\x1b[1;31m[!] \x1b[0mInclude: \x1b[0;32m${line} \x1b[0mnot found!"
            return 1
        fi
        source "${root}/${line}.sh"
    done <<< "${input_data}"

    return 0
}; readonly -f include