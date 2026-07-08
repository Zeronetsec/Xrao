function install::checker() {
    local name="xrao"
    if command -v ${name} &>/dev/null; then
        echo -e "${GG}[+] ${N}${name^} installed!"
        echo -e "${GG}[+] ${N}Usage: ${GG}${name} --help ${N}to show helper"
        return 0
    else
        echo -e "${R}[!] ${N}Failed installing ${name}!"
        return 1
    fi
}; readonly -f install::checker