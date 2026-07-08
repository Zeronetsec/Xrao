function install::zinstall() {
    function __exit__() {
        local code="${1}"
        if [[ "${code}" -gt 0 ]]; then
            echo -e " ${DG}- ${N}error: ${R}${code}${N}"
            return ${code}
        fi
        echo -e " ${DG}- ${N}exit: ${GG}${code}${N}"
        return ${code}
    }

    local pkg="${1}"
    if [[ -z "${pkg}" ]]; then
        return 1
    fi

    echo -ne "${DG}-> ${N}Try: ${GG}${pkg}${N}"

    if command -v pacman &>/dev/null; then
        if command -v paru &>/dev/null; then
            command paru -S --noconfirm "${pkg}" > /dev/null 2>&1
            __exit__ ${?}
        elif command -v yay &>/dev/null; then
            command yay -S --noconfirm "${pkg}" > /dev/null 2>&1
            __exit__ ${?}
        else
            command pacman -S --noconfirm "${pkg}" > /dev/null 2>&1
            __exit__ ${?}
        fi
    elif command -v apt &>/dev/null; then
        command apt install -y "${pkg}" > /dev/null 2>&1
        __exit__ ${?}
    elif command -v dnf &>/dev/null; then
        command dnf install -y "${pkg}" > /dev/null 2>&1
        __exit__ ${?}
    elif command -v zypper &>/dev/null; then
        command zypper in -y "${pkg}" > /dev/null 2>&1
        __exit__ ${?}
    elif command -v apk &>/dev/null; then
        command apk add "${pkg}" > /dev/null 2>&1
        __exit__ ${?}
    elif command -v flatpak &>/dev/null; then
        command flatpak install -y "${pkg}" > /dev/null 2>&1
        __exit__ ${?}
    elif command -v snap &>/dev/null; then
        command snap install "${pkg}" > /dev/null 2>&1
        __exit__ ${?}
    else
        echo -e "${R}[!] ${N}Unknown package manager!"
        return 1
    fi
}; readonly -f install::zinstall