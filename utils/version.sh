# Xrao Project

function __version__() {
    project="Xrao"
    ver="v1.0"
    creator="Zeronetsec"
    homepage="https://github.com/Zeronetsec/Xrao"

    echo
    echo -e "Project: $project"
    echo -e "Version: $ver"
    echo -e "Creator: $creator"
    echo -e "Homepage: $homepage"

    if [[ "$newline" == true ]]; then
        echo
    fi
}

# Copyright (c) 2026 Zeronetsec