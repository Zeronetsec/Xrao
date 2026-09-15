function install::zparser() {
    local raw_input="${1}"
    if [[ -z "${raw_input}" ]]; then
        return 1
    fi

    local target_alias=""
    local input="${raw_input}"

    if [[ "${raw_input}" =~ "->" ]]; then
        target_alias="${raw_input#*->}"
        target_alias="$(
            echo -e "${target_alias}" | \
                command tr -d '[:space:]'
        )"

        input="${raw_input%%->*}"
        input="$(
            echo -e "${input}" | \
                command tr -s '[:space:]' ' ' | \
                command sed 's/ $//'
        )"
    fi

    local installed_pkg=""

    if [[ "${input}" == *"::"* ]]; then
        local p1="${input%%::*}"
        local fallback_raw="${input#*::}"

        set +o errexit
        install::zinstall "${p1}"
        local status=${?}
        set -o errexit

        if [[ ${status} -eq 0 ]]; then
            installed_pkg="${p1}"
        else
            local clean="${fallback_raw//[\{\}]/}"
            local IFS=','
            read -ra package_fallback <<< "${clean}"

            local allow_skip=false

            for pkg in "${package_fallback[@]}"; do
                pkg="$(
                    echo -e "${pkg}" | \
                        command tr -d '[:space:]'
                )"

                [[ -z "${pkg}" ]] && continue

                if [[ "${pkg}" == "-" ]]; then
                    allow_skip=true
                    continue
                fi

                set +o errexit
                install::zinstall "${pkg}"
                local status_fallback=${?}
                set -o errexit

                if [[ ${status_fallback} -eq 0 ]]; then
                    installed_pkg="${pkg}"
                    break
                fi
            done

            if [[ -z "${installed_pkg}" && "${allow_skip}" == true ]]; then
                echo -e "${color_DG}-> ${color_N}Skipped..."
                return 0
            fi
        fi
    else
        set +o errexit
        install::zinstall "${input}"
        local status_single=${?}
        set -o errexit

        if [[ ${status_single} -eq 0 ]]; then
            installed_pkg="${input}"
        fi
    fi

    if [[ -n "${installed_pkg}" && -n "${target_alias}" ]]; then
        install::zsymlink "${installed_pkg}" "${target_alias}"
    elif [[ -z "${installed_pkg}" ]]; then
        return 1
    fi

    return 0
}; readonly -f install::zparser