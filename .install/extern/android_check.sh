function install::extern::androidCheck() {
    function __android_check__() {
        if [[ -f /etc/os-release ]]; then
            local os_id="$(
                command grep -E '^ID=' /etc/os-release | \
                command cut -d= -f2 | \
                command tr -d '"'
            )"

            if [[ "${os_id}" =~ ^(ubuntu|debian|kali|parrot|fedora|arch|alpine|linux)$ ]]; then
                return 1
            fi
        fi

        if [[ ! "$(command whoami)" =~ ^u0_a[0-9]+ ]]; then
            if [[ "$(command id -u)" != "0" ]]; then
                return 1
            fi
        fi

        if [[ ! -f "/system/bin/linker" && ! -f "/system/bin/linker64" ]]; then
            return 1
        fi

        if [[ ! -d "/data/data/com.termux/files/home" ]]; then
            return 1
        fi

        if [[ -z "${PREFIX}" || "${PREFIX}" != *"/com.termux/"* || ! -d "${PREFIX}" ]]; then
            return 1
        fi

        if [[ -z "${TERMUX_VERSION}" ]]; then
            return 1
        fi

        return 0
    }

    echo -e "${B}[*] ${N}Checking android environment..."
    __android_check__ || {
        echo -e "${R}[!] ${N}Termux environment not detected."
        echo -e "${R}[!] ${N}This tool is designed exclusively for the Termux Android app."
        return 1
    }
}; readonly -f install::extern::androidCheck