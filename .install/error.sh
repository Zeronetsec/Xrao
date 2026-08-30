function install::error() {
    local err="${?}"
    local cmd="${BASH_COMMAND}"

    echo -e " ${color_DG}- ${color_N}error: ${color_R}${err}${color_N}"
    echo -e "${color_N}${cmd}${color_N}"

}; readonly -f install::error
trap install::error ERR EXIT