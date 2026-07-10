// https://github.com/Zeronetsec/Xrao

package variable

import (
    "os"
    "path/filepath"
)

const (
    Dlt = "/data/local/tmp"
    Sdx = "/sdcard/Download/xrao"
)

var (
    Prefix string
    Tmp string
    Config string
    Home string
)

func init() {
    Prefix = os.Getenv("PREFIX")
    Home = os.Getenv("HOME")

    if Prefix == "" {
        Prefix = "/usr"
    }

    Tmp = filepath.Join(Prefix, "tmp")
    Config = filepath.Join(Home, ".xrao", "config.xr")
}

// Copyright (c) 2026 Zeronetsec