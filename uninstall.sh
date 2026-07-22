#!/usr/bin/env bash

set -o errexit

src="${BASH_SOURCE[0]}"
while [[ -h "${src}" ]]; do
    dir="$(
        cd -P "$(
            command dirname "${src}"
        )" > /dev/null 2>&1 && pwd
    )"
    src="$(command readlink "${src}")"
    [[ "${src}" != /* ]] && src="${dir}/${src}"
done

dir="$(
    cd -P "$(
        command dirname "${src}"
    )" > /dev/null 2>&1 && pwd
)"

export root="${dir}"; readonly root
source "${root}/.install/include.sh"

include : '(
    .install/color
    .install/variable
    .install/error
    .install/getinstall
)'

__RMBK__=false

while [[ ${#} -gt 0 ]]; do
    case "${1}" in
        "--remove-backup") export __RMBK__=true ;;
    esac
    shift
done

if [[ "${__RMBK__}" == true ]]; then
    install::getinstall \
        "command rm -f ${opt}/xrao_*.zip.bak" \
        "Removing all backup..."
fi

install::getinstall \
    "command rm -rf ${opt}/xrao" \
    "Removing: ${GG}${opt}/xrao${N}"

install::getinstall \
    "command rm -f ${bin}/xrao" \
    "Removing: ${GG}${bin}/xrao${N}"

install::getinstall \
    "command rm -rf ${HOME}/.xrao"
    "Removing: ${GG}${HOME}/.xrao${N}"

echo -e "${GG}[+] ${N}Xrao removed"

trap - EXIT
exit ${?}