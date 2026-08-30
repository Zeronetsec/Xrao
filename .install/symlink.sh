function install::symlink() {
    install::getinstall \
        "
            command ln -sf \
                ${opt}/${targetins}/${targetsyml} \
                ${bin}/${targetins}
        " \
        "Symlink: ${color_GG}${opt}/${targetins}/${targetsyml} ${color_DG}-> ${color_GG}${bin}/${targetins}${color_N}"
}; readonly -f install::symlink