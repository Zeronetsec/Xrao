function install::postins() {
    if [[ -d "${opt}/${targetins}" ]]; then
        install::getinstall \
            "command rm -rf ${opt}/${targetins}" \
            "Removing old source..."
    fi

    install::getinstall \
        "command mv ${root} ${opt}/${targetins}" \
        "Moving: ${GG}${root} ${DG}-> ${GG}${opt}/${targetins}${N}"
}; readonly -f install::postins