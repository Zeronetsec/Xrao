function install::zparser() {
    local input="${1}"
    if [[ -z "${input}" ]]; then
        return 1
    fi

    if [[ "${input}" == *"::"* ]]; then
        local p1="${input%%::*}"
        local fallback_raw="${input#*::}"

        set +o errexit
        install::zinstall "${p1}"
        local status=${?}
        set -o errexit

        if [[ ${status} -eq 0 ]]; then
            return 0
        fi

        local clean="${fallback_raw//[\{\}]/}"
        local IFS=','
        read -ra package_fallback <<< "${clean}"

        for pkg in "${package_fallback[@]}"; do
            pkg="$(
                echo -e "${pkg}" | \
                command tr -d '[:space:]'
            )"

            [[ -z "${pkg}" ]] && continue

            set +o errexit
            install::zinstall "${pkg}"
            set -o errexit
        done
        return 0
    else
        install::zinstall "${input}"
    fi
}; readonly -f install::zparser