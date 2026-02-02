# Xrao Project

function __eject__() {
    echo -e '
Xrao Ducky Back To Sleep
   _
__(-)<  ‹—(zZZ
\___)

[-] Ejecting Xrao'

    . $init/eject_functions.sh
    . $config/eject_settings.sh

    __reset_thermal__
    __reset_perfm__
    __reset_driver__
    __rdud__
    __reset_touch__
    __rgame_mode__
    __rllm__
    __reset_whitelist__
    __reset_anim__

    echo
}

# Copyright (c) 2026 Zeronetsec