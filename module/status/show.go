// https://github.com/Zeronetsec/Xrao

package status

import (
    "fmt"
    "os"
    "github.com/Zeronetsec/Xrao/execparser"
    "github.com/Zeronetsec/Xrao/utils/color"
    "github.com/Zeronetsec/Xrao/utils/shell"
    "github.com/Zeronetsec/Xrao/utils/variable"
)

func Show(conf string) {
    _, err := shell.Execf("adb get-state")
    if err != nil {
        fmt.Printf(
            "%s[!] %sADB device not detected!\n",
            color.R, color.N,
        )
        return
    }

    if conf == "" {
        conf = variable.Config
    }

    if _, err := os.Stat(conf); os.IsNotExist(err) {
        fmt.Printf(
            "%s[!] %sFile: %s%s %snot found!\n",
            color.R, color.N, color.GG, conf, color.N,
        )
        return
    }

    execparser.CheckStatus(conf)
}

// Copyright (c) 2026 Zeronetsec