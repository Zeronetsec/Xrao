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
                "Backup: ${GG}${opt}/${targetins} ${DG}-> ${GG}${opt}/${targetins}_${bkdate}.bak.zip${N}"
            cd
        )
    fi
}; readonly -f install::backup