function install::getinstall() {
    local cmd="${1}"
    local desc="${2}"

    echo -ne "${color_B}[*] ${color_N}${desc}"
    eval "${cmd}" > /dev/null 2>&1
    local status=${?}

    echo -e " ${color_DG}- ${color_N}exit: ${color_GG}${status}${color_N}"
}; readonly -f install::getinstall