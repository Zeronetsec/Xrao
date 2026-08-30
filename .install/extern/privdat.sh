function install::extern::privdat() {
    if [[ -d "${root}/.privdat" ]]; then
        echo -e "${color_B}[*] ${color_N}Setting up private data..."
        install::getinstall \
            "command rm -rf ${root}/config" \
            "Removing: ${color_GG}${root}/config${color_N}"

        install::getinstall \
            "
                command cp -r \
                    ${root}/.privdat \
                    ${root}/config
            " \
            "Copying: ${color_GG}${root}/.privdat ${color_DG}-> ${color_GG}${root}/config${color_N}"
    fi
}; readonly -f install::extern::privdat