function install::getinstall() {
    local cmd="${1}"
    local desc="${2}"

    echo -ne "${color_B}[*] ${color_N}${desc}"
    set +o errexit

    local error_msg
    error_msg="$(
        eval "${cmd}" 2>&1 >/dev/null
    )"
    local status=${?}

    set -o errexit

    if [[ ${status} -eq 0 ]]; then
        echo -e " ${color_DG}- ${color_N}exit: ${color_GG}${status}${color_N}"
    else
        echo -e " ${color_DG}- ${color_N}error: ${color_R}${status}${color_N}"
        echo -e "${color_N}${error_msg}${color_N}" | \
            command sed 's/^/    /'
    fi

    return ${status}
}; readonly -f install::getinstall