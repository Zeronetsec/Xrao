function install::installer() {
    if [[ ! -d "${HOME}/.${targetins}" ]]; then
        install::getinstall \
            "command mkdir -p ${HOME}/.${targetins}" \
            "Create directory: ${color_GG}${HOME}/.${targetins}${color_N}"
    fi

    if [[ ! -f "${HOME}/.${targetins}/config.xr" ]]; then
        install::getinstall \
            "
                command cp \
                    ${opt}/${targetins}/config/config.xr \
                    ${HOME}/.${targetins}/
            " \
            "Copying: ${color_GG}${opt}/${targetins}/config/config.xr ${color_DG}-> ${color_GG}${HOME}/.${targetins}/${color_N}"
    fi

    (
        cd "${opt}/${targetins}"
        install::getinstall \
            "
                command go mod tidy
                command go build -o ${targetins}
            " \
            "Compiling: ${color_GG}${targetins}${color_N}"
    )
}; readonly -f install::installer