# Xrao Project

function __show_config__() {
    echo -e "\n[*] Xrao settings"
    echo -e "Basefolder: $path"
    echo -e "Downscale: $downscale"
    echo -e "FPS: $fps"
    echo -e "Animation: $animation"
    echo -e "Driver: $driver"

    echo -e "\n[*] Affected Apps Settings"
    for package in $affect; do
        echo -e "Apps: $package"
    done

    echo -e "\n[*] Kill Apps Settings"
    for package in $kill; do
        echo -e "Apps: $package"
    done

    echo -ne "\n[*] Inject Settings"
    var=$(cat $config/inject_settings.sh | grep -v '#')
    echo -ne "$var\n"

    echo -ne "\n[*] Eject Settings"
    var=$(cat $config/eject_settings.sh | grep -v '#')
    echo -ne "$var\n"

    if [[ "$newline" == true ]]; then
        echo
    fi
}

# Copyright (c) 2026 Zeronetsec