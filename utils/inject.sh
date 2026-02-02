# Xrao Project

function __inject__() {
    echo -e '
Spawned Xrao Ducky...
   _
__(•)<  ‹—(Honk
\___)

[+] Applying system optimizations'

    . $init/inject_functions.sh
    . $config/inject_settings.sh

    __disable_thermal__
    __downscale__
    __angle__
    __edud__
    __touch_booster__
    __gmode__
    __low_latency__
    __whitelist__
    __system_anim__
    __kill__

    echo
}

# Copyright (c) 2026 Zeronetsec