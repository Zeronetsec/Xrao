function install::extern::privdat() {
    if [[ -d "${root}/.privdat" ]]; then
        echo -e "${B}[*] ${N}Setting up private data..."
        command rm -rf "${root}/config"
        command mv "${root}/.privdat" "${root}/config"
    fi
}