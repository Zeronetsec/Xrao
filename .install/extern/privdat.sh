function install::extern::privdat() {
    if [[ -d "${root}/.privdat" ]]; then
        echo -e "${B}[*] ${N}Setting up private data..."
        command rm -rf "${root}/config"
        command cp -r \
            "${root}/.privdat" \
            "${root}/config"
    fi
}; readonly -f install::extern::privdat