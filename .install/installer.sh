function install::installer() {
    if [[ "${__BACKUP__}" == true && -d "${opt}/xrao" ]]; then
        (
            cd "${opt}"
            install::getinstall \
                "
                    command zip -r \
                        xrao_${bkdate}.bak.zip \
                        xrao
                " \
                "Backup: ${GG}${opt}/xrao ${DG}-> ${GG}${opt}/xrao_${bkdate}.bak.zip${N}"
            cd
        )
    fi

    if [[ -d "${opt}/xrao" ]]; then
        install::getinstall \
            "command rm -rf ${opt}/xrao" \
            "Removing old source..."
    fi

    install::getinstall \
        "command mv ${root} ${opt}/xrao" \
        "Moving: ${GG}${root} ${DG}-> ${GG}${opt}/xrao${N}"

    if [[ ! -d "${HOME}/.xrao" ]]; then
        install::getinstall \
            "command mkdir -p ${HOME}/.xrao" \
            "Create directory: ${GG}${HOME}/.xrao${N}"
    fi

    if [[ ! -f "${HOME}/.xrao/config.xr" ]]; then
        install::getinstall \
            "
                command cp \
                    ${opt}/xrao/config/config.xr \
                    ${HOME}/.xrao/
            " \
            "Copying: ${GG}${opt}/xrao/config/config.xr ${DG}-> ${GG}${HOME}/.xrao/${N}"
    fi

    (
        cd "${opt}/xrao"
        install::getinstall \
            "command go mod tidy" \
            "Retidy: ${GG}xrao${N}"

        install::getinstall \
            "command go build -o xrao" \
            "Compiling: ${GG}xrao${N}"
        cd
    )

    install::getinstall \
        "
            command ln -sf \
                ${opt}/xrao/xrao \
                ${bin}/xrao
        " \
        "Symlink: ${GG}${opt}/xrao/xrao ${DG}-> ${GG}${bin}/xrao${N}"
}; readonly -f install::installer