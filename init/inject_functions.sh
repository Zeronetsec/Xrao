# Xrao Project

function __disable_thermal__() {
    echo -e "\n[+] Disable Thermal Throttle"
    if [ $disable_thermal_throttle == true ]; then
        for package in $affect; do
            echo -e "Boost: $package"
            cmd game set thermal-throttle "$package" false
        done
    else
        echo -e "false"
    fi
}

function __downscale__() {
    echo -e "\n[+] Set Performance Mode"
    if [ $set_performance_mode == true ]; then
        echo -e "[+] Downscalling Ratio"
        echo -e "[+] Lock FPS"
        echo -e "Downscale: $downscale"
        echo -e "FPS: $fps"
        for package in $affect; do
            echo -e "Boost: $package"
            cmd game set --mode 2 --downscale "$downscale" --fps "$fps" "$package"
        done
    else
        echo -e "false"
    fi
}

function __angle__() {
    echo -e "\n[+] Force Enable Angle"
    if [ $force_enable_angle == true ]; then
        for package in $affect; do
            echo -e "Boost: $package"
        done
        package=""
        for b in $affect; do
            [ -z "$package" ] && package="$b" || package="$package,$b"
        done
        echo -e "Package: $package"
        settings put global angle_gl_driver_selection_pkgs "$package"
        echo -e "Driver: $driver"
        settings put global angle_gl_driver_selection_values "$driver"
    else
        echo -e "false"
    fi
}

function __edud__() {
    echo -e "\n[+] Enable Dynamic Updatable Driver"
    if [ $enable_dynamic_updatable_driver == true ]; then
        for package in $affect; do
            echo -e "Boost: $package"
        done
        package=""
        for b in $affect; do
            [ -z "$package" ] && package="$b" || package="$package,$b"
        done
        echo -e "Package: $package"
        settings put global updatable_driver_production_opt_in_apps "$package"
    else
        echo -e "false"
    fi
}

function __touch_booster__() {
    echo -e "\n[+] Enable Touch Booster"
    if [ $enable_touch_booster == true ]; then
        for package in $affect; do
            echo -e "Boost: $package"
        done
        package=""
        for b in $affect; do
            [ -z "$package" ] && package="$b" || package="$package,$b"
        done
        echo -e "Package: $package"
        settings put system input_boost_game_pkg "$package"
    else
        echo -e "false"
    fi
}

function __gmode__() {
    echo -e "\n[+] Enable Game Mode"
    if [ $enable_game_mode == true ]; then
        for package in $affect; do
            echo -e "Boost: $package"
        done
        package=""
        for b in $affect; do
            [ -z "$package" ] && package="$b" || package="$package,$b"
        done
        echo -e "Package: $package"
        settings put secure game_mode_list "$package"
    else
        echo -e "false"
    fi
}

function __low_latency__() {
    echo -e "\n[+] Enable Low Latency Mode"
    if [ $enable_low_latency_mode == true ]; then
        for package in $affect; do
            echo -e "Boost: $package"
        done
        package=""
        for b in $affect; do
            [ -z "$package" ] && package="$b" || package="$package,$b"
        done
        echo -e "Package: $package"
        settings put global game_low_latency_mode "$package"
    else
        echo -e "false"
    fi
}

function __whitelist__() {
    echo -e "\n[+] Whitelist Apps"
    if [ $whitelist_apps == true ]; then
        for package in $affect; do
            echo -e "Boost: $package"
            dumpsys deviceidle whitelist +"$package"
        done
    else
        echo -e "false"
    fi
}

function __system_anim__() {
    echo -e "\n[+] Force Down System Animation"
    if [ $force_down_system_animation == true ]; then
        echo -e "Window: $animation"
        settings put global window_animation_scale "$animation"
        echo -e "Transition: $animation"
        settings put global transition_animation_scale "$animation"
        echo -e "Animator: $animation"
        settings put global animator_duration_scale "$animation"
    else
        echo -e "false"
    fi
}

function __kill__() {
    echo -e "\n[+] Kill Apps Process"
    if [ $kill_apps_process == true ]; then
        for package in $kill; do
            echo -e "Kill: $package"
            am force-stop "$package"
        done
    else
        echo -e "false"
    fi
}

# Copyright (c) 2026 Zeronetsec