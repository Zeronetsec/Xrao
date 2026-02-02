# Xrao Project

function __package_status__() {
    echo -e "\n[*] Package Status"
    for package in $affect; do
        echo -e "Lookup: $package"
        dumpsys game | command grep "$package"
    done
}

function __list_mode__() {
    echo -e "\n[*] List Mode"
    for package in $affect; do
        echo -e "Lookup: $package"
        cmd game list-modes "$package"
    done
}

function __list_config__() {
    echo -e "\n[*] List Config"
    for package in $affect; do
        echo -e "Lookup: $package"
        cmd game list-configs "$package"
    done
}

function __dss__() {
    echo -e "\n[*] Driver Selection Status"
    settings get global angle_gl_driver_selection_pkgs
    settings get global angle_gl_driver_selection_values
}

function __duds__() {
    echo -e "\n[*] Dynamic Updatable Driver Status"
    settings get global updatable_driver_production_opt_in_apps
}

function __touch_booster_status__() {
    echo -e "\n[*] Touch Booster Status"
    settings get system input_boost_game_pkg
}

function __gmode_status__() {
    echo -e "\n[*] Game Mode Status"
    settings get secure game_mode_list
}

function __llms__() {
    echo -e "\n[*] Low Latency Mode Status"
    settings get global game_low_latency_mode
}

function __whitelist_apps__() {
    echo -e "\n[*] Whitelist Apps"
    dumpsys deviceidle whitelist
}

function __anim_status__() {
    echo -e "\n[*] Animation Status"
    echo -e "Animation: $(settings get global window_animation_scale)"
    echo -e "Transition: $(settings get global transition_animation_scale)"
    echo -e "Animator: $(settings get global animator_duration_scale)"
}

# Copyright (c) 2026 Zeronetsec