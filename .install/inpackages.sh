function install::inpackages() {
    command mapfile -t packages < <(
        command cat "${root}/.install/packages.txt"
    )

    for line in "${packages[@]}"; do
        [[ -z "${line}" ]] && continue
        [[ "${line}" =~ ^# ]] && continue
        echo -e "${color_B}[*] ${color_N}Installing: ${color_GG}${line}${color_N}"
        install::zparser "${line}"
    done
}; readonly -f install::inpackages