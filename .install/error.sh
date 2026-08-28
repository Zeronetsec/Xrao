function install::error() {
    local cmd="${BASH_COMMAND}"

    echo -e " ${DG}- ${N}error: ${R}${?}${N}"

    if [[ "${__ERRSHOW__}" == true ]]; then
        echo -e "${N}${cmd}"
    fi
}; readonly -f install::error

trap install::error ERR EXIT