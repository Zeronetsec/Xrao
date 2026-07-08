function install::inpackages() {
    command mapfile -t packages < <(
        command cat "${root}/.install/packages.txt"
    )

    for line in "${packages[@]}"; do
        [[ -z "${line}" ]] && continue
        [[ "${line}" =~ ^# ]] && continue
        echo -e "${B}[*] ${N}Installing: ${GG}${line}${N}"
        install::zparser "${line}"
    done
}; readonly -f install::inpackages