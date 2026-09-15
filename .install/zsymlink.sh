function install::zsymlink() {
    local installed_name="${1}"
    local target_alias="${2}"

    if [[ "${installed_name}" == "${target_alias}" ]]; then
        return 0
    fi

    local src_path
    src_path="$(
        command -v "${installed_name}" 2>/dev/null
    )"

    if [[ -z "${src_path}" ]]; then
        echo -e "${color_R}[!] ${color_N}Binary: ${color_GG}${installed_name} ${color_N}not found to symlink!"
        return 1
    fi

    local target_path="${bin}/${target_alias}"

    if [[ "${src_path}" == "${target_path}" ]]; then
        return 0
    fi

    echo -e "${color_DG}-> ${color_N}Symlink: ${color_GG}${installed_name} ${color_DG}-> ${color_GG}${target_path}${color_N}"
    command ln -sf \
        "${src_path}" \
        "${target_path}"

    return ${?}
}; readonly -f install::zsymlink