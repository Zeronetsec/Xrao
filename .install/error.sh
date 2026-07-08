function install::error() {
    echo -e " ${DG}- ${N}error: ${R}${?}${N}"
}; readonly -f install::error

trap install::error EXIT