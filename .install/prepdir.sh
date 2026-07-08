function install::prepdir() {
    if [[ ! -d "${tmp}" ]]; then
        command mkdir -p "${tmp}"
    fi

    if [[ ! -d "${opt}" ]]; then
        command mkdir -p "${opt}"
    fi

    if [[ ! -d "${bin}" ]]; then
        command mkdir -p "${bin}"
    fi
}; readonly -f install::prepdir