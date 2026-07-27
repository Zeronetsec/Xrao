function install::installer() {
    if [[ ! -d "${HOME}/.${targetins}" ]]; then
        install::getinstall \
            "command mkdir -p ${HOME}/.${targetins}" \
            "Create directory: ${GG}${HOME}/.${targetins}${N}"
    fi

    if [[ ! -f "${HOME}/.${targetins}/config.xr" ]]; then
        install::getinstall \
            "
                command cp \
                    ${opt}/${targetins}/config/config.xr \
                    ${HOME}/.${targetins}/
            " \
            "Copying: ${GG}${opt}/${targetins}/config/config.xr ${DG}-> ${GG}${HOME}/.${targetins}/${N}"
    fi

    (
        cd "${opt}/${targetins}"
        install::getinstall \
            "command go mod tidy" \
            "Retidy: ${GG}${targetins}${N}"

        install::getinstall \
            "command go build -o ${targetins}" \
            "Compiling: ${GG}${targetins}${N}"
        cd
    )
}; readonly -f install::installer