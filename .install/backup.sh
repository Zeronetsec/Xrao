function install::backup() {
    if [[ "${__BACKUP__}" == true && -d "${opt}/${targetins}" ]]; then
        (
            cd "${opt}"
            install::getinstall \
                "
                    command zip -r \
                        ${targetins}_${bkdate}.bak.zip \
                        ${targetins}
                " \
                "Backup: ${color_GG}${opt}/${targetins} ${color_DG}-> ${color_GG}${opt}/${targetins}_${bkdate}.bak.zip${color_N}"
        )
    fi
}; readonly -f install::backup