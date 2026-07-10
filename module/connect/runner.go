// https://github.com/Zeronetsec/Xrao

package connect

import (
    "fmt"
    "github.com/Zeronetsec/Xrao/utils/color"
    "github.com/Zeronetsec/Xrao/utils/shell"
)

func Runner(target string) {
    _, _ = shell.Execf("adb start-server")

    fmt.Printf(
        "%s[*] %sConnect to: %s%s%s\n",
        color.B, color.N, color.GG, target, color.N,
    )

    _ = shell.ExecLivef("adb connect %s", target)
}

// Copyright (c) 2026 Zeronetsec