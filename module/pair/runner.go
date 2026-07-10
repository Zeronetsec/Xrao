// https://github.com/Zeronetsec/Xrao

package pair

import (
    "fmt"
    "os"
    "os/exec"
    "github.com/Zeronetsec/Xrao/utils/color"
    "github.com/Zeronetsec/Xrao/utils/shell"
)

func Runner(target string) {
    _, _ = shell.Execf("adb start-server")

    fmt.Printf(
        "%s[*] %sPairing to: %s%s%s\n",
        color.B, color.N, color.GG, target, color.N,
    )

    cmd := exec.Command("adb", "pair", target)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Stdin = os.Stdin

    if err := cmd.Run(); err != nil {
        fmt.Printf(
            "%s[!] %sPairing failed!\n",
            color.R, color.N,
        )
        return
    }
}

// Copyright (c) 2026 Zeronetsec