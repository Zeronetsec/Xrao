function install::getinstall() {
    local cmd="${1}"
    local desc="${2}"

    echo -ne "${B}[*] ${N}${desc}"
    eval "${cmd}" > /dev/null 2>&1

    local status=${?}
    echo -e " ${DG}- ${N}exit: ${GG}${status}${N}"
}; readonly -f install::getinstall