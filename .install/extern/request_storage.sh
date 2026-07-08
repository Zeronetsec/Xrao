function install::extern::requestStorage() {
    set +o errexit

    function __test__() {
        command touch "/sdcard/__xrao_test_storage_access__" > /dev/null 2>&1 || return 1
        command rm "/sdcard/__xrao_test_storage_access__" > /dev/null 2>&1 || return 1
        return 0
    }

    if ! __test__; then
        set -o errexit
        echo -e "${R}[!] ${N}Storage access denied!"
        echo -e "${R}[!] ${N}Please change access permissions using: ${GG}termux-setup-storage${N}"
        exit 1
    fi

    set -o errexit
}; readonly -f install::extern::requestStorage