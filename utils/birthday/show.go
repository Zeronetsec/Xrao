// https://github.com/Zeronetsec/Xrao

package birthday

import (
    "fmt"
    "time"
    "github.com/Zeronetsec/Xrao/utils/color"
)

func Show() {
    birthDate := "06-08"
    now := time.Now().Format("01-02")
    if now == birthDate {
        fmt.Printf(
            "%s› %sHappy birthday for %sxrao %s🎉\n",
            color.R, color.N, color.GG, color.N,
        )
        fmt.Println()
    }
}

// Copyright (c) 2026 Zeronetsec