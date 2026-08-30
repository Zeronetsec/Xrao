function install::postins() {
    if [[ -d "${opt}/${targetins}" ]]; then
        install::getinstall \
            "command rm -rf ${opt}/${targetins}" \
            "Removing old source..."
    fi

    install::getinstall \
        "command mv ${root} ${opt}/${targetins}" \
        "Moving: ${color_GG}${root} ${color_DG}-> ${color_GG}${opt}/${targetins}${color_N}"
}; readonly -f install::postins