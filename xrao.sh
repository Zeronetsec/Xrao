#!/system/bin/sh
# Xrao Project

name=$(basename "$0")

if [[ -z "$1" && -z "$2" ]]; then
    echo -e "\n[!] Invalid input!"
    echo -e "[!] Try: $name --path=\"/path/to/Xrao\" --help"
    echo -e "[!] Example: $name --path=\"/sdcard/Download/Xrao\" --status\n"
    exit 1
fi

if [[ ! "$1" == "--path="* ]]; then
    echo -e "\n[!] Invalid Options: $1"
    echo -e "[!] Try: $name --path=\"/path/to/Xrao\" --help"
    echo -e "[!] Example: $name --path=\"/sdcard/Download/Xrao\" --status\n"
    exit 1
fi

path=${1#*=}
. $path/init/init.sh

case "$2" in
    "--inject")
        __inject__
        exit 0
        ;;
    "--eject")
        __eject__
        exit 0
        ;;
    "--status")
        __status__
        exit 0
        ;;
    "--show-config")
        newline=true
        __show_config__
        newline=false
        exit 0
        ;;
    "--help")
        __helper__
        exit 0
        ;;
    "--version")
        newline=true
        __version__
        newline=false
        exit 0
        ;;
    *)
        if [[ -n "$path" && -z "$2" ]]; then
            echo -e "\n[!] Invalid input!"
            echo -e "[!] Try: $name --path=\"$path\" --help"
            echo -e "[!] Example: $name --path=\"$path\" --status\n"
            exit 1
        fi
        echo -e "\n[!] Invalid options: $2"
        echo -e "[!] Try: $name --path=\"$path\" --help"
        echo -e "[!] Example: $name --path=\"$path\" --status\n"
        exit 1
        ;;
esac

# Copyright (c) 2026 Zeronetsec