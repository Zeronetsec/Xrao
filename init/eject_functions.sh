# Xrao Project

function __reset_thermal__() {
    echo -e "\n[-] Reset Thermal Throttle"
    if [ $reset_thermal_throttle == true ]; then
        for package in $affect; do
            echo -e "Reset: $package"
            cmd game set thermal-throttle "$package" true
            cmd game reset thermal-throttle "$package"
        done
    else
        echo -e "false"
    fi
}

function __reset_perfm__() {
    echo -e "\n[-] Reset Performance Mode"
    if [ $reset_performance_mode == true ]; then
        for package in $affect; do
            echo -e "Reset: $package"
            cmd game reset "$package"
        done
    else
        echo -e "false"
    fi
}

function __reset_driver__() {
    echo -e "\n[-] Reset Driver Selection"
    if [ $reset_driver_selection == true ]; then
        settings delete global angle_gl_driver_selection_pkgs
        settings delete global angle_gl_driver_selection_values
    else
        echo -e "false"
    fi
}

function __rdud__() {
    echo -e "\n[-] Reset Dynamic Updatable Driver"
    if [ $reset_dynamic_updatable_driver == true ]; then
        settings delete global updatable_driver_production_opt_in_apps
    else
        echo -e "false"
    fi
}

function __reset_touch__() {
    echo -e "\n[-] Reset Touch Booster"
    if [ $reset_touch_booster == true ]; then
        settings delete system input_boost_game_pkg
    else
        echo -e "false"
    fi
}

function __rgame_mode__() {
    echo -e "\n[-] Reset Game Mode"
    if [ $reset_game_mode == true ]; then
        settings delete secure game_mode_list
    else
        echo -e "false"
    fi
}

function __rllm__() {
    echo -e "\n[-] Reset low Latency Mode"
    if [ $reset_low_latency_mode == true ]; then
        settings delete global game_low_latency_mode
    else
        echo -e "false"
    fi
}

function __reset_whitelist__() {
    echo -e "\n[-] Reset Whitelist Apps"
    if [ $reset_whitelist_apps == true ]; then
        for package in $affect; do
            echo -e "Reset: $package"
            dumpsys deviceidle whitelist -"$package"
        done
    else
        echo -e "false"
    fi
}

function __reset_anim__() {
    echo -e "\n[-] Reset System Animation"
    if [ $reset_system_animation == true ]; then
        echo -e "Window: 1.0"
        settings put global window_animation_scale 1.0
        echo -e "Transition: 1.0"
        settings put global transition_animation_scale 1.0
        echo -e "Animator: 1.0"
        settings put global animator_duration_scale 1.0
    else
        echo -e "false"
    fi
}

# Copyright (c) 2026 Zeronetsec