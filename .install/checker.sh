function install::checker() {
    if command -v ${targetins} &>/dev/null; then
        echo -e "${color_GG}[+] ${color_N}${targetins^} installed!"
        echo -e "${color_GG}[+] ${color_N}Usage: ${color_GG}${targetins} --help ${color_N}to show helper"
        return 0
    else
        echo -e "${color_R}[!] ${color_N}Failed installing: ${color_GG}${targetins}${color_N}"
        return 1
    fi
}; readonly -f install::checker