function install::error() {
    local cmd="${BASH_COMMAND}"
    echo -e " ${DG}- ${N}error: ${R}${?}${N}"
    echo -e "${DG}[${GG}${cmd}${DG}]${N}"
}; readonly -f install::error

trap install::error ERR EXIT