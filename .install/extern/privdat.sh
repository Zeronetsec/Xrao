function install::extern::privdat() {
    if [[ -d "${root}/.privdat" ]]; then
        echo -e "${B}[*] ${N}Setting up private data..."
        install::getinstall \
            "command rm -rf ${root}/config" \
            "Removing: ${GG}${root}/config${N}"

        install::getinstall \
            "
                command cp -r \
                    ${root}/.privdat \
                    ${root}/config
            " \
            "Copying: ${GG}${root}/.privdat ${DG}-> ${GG}${root}/config${N}"
    fi
}; readonly -f install::extern::privdat