function install::symlink() {
    install::getinstall \
        "
            command ln -sf \
                ${opt}/${targetins}/${targetsyml} \
                ${bin}/${targetins}
        " \
        "Symlink: ${GG}${opt}/${targetins}/${targetsyml} ${DG}-> ${GG}${bin}/${targetins}${N}"
}; readonly -f install::symlink