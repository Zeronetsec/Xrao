function install::error() {
    local err="${?}"
    local cmd="${BASH_COMMAND}"
    local cmd2="${__cmd__}"

    echo -e " ${DG}- ${N}error: ${R}${err}${N}"
    echo -e "${R}1${DG}[${N}${cmd}${DG}]${N}"
    echo -e "${R}2${DG}[${N}${cmd2}${DG}]${N}"

}; readonly -f install::error
trap install::error ERR EXIT