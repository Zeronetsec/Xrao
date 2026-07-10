// https://github.com/Zeronetsec/Xrao

package banner

import (
    "embed"
    "fmt"
    "github.com/Zeronetsec/Xrao/utils/color"
)

//go:embed ascii/*.txt
var asciiFS embed.FS

func Show() {
    data, err := asciiFS.ReadFile("ascii/banner.txt")
    if err != nil {
        fmt.Printf(
            "%s[!] %sError loading banner!\n",
            color.R, color.N,
        )
        return
    }

    fmt.Printf(
        "%s%s%s\n",
        color.B, string(data), color.N,
    )
}

// Copyright (c) 2026 Zeronetsec