# Xrao Project

config=$path/config
utils=$path/utils
init=$path/init

. $config/affected_apps.sh
. $config/xrao_settings.sh
. $config/kill_apps.sh

. $utils/inject.sh
. $utils/eject.sh
. $utils/status.sh
. $utils/version.sh
. $utils/helper.sh
. $utils/show_config.sh

# Copyright (c) 2026 Zeronetsec