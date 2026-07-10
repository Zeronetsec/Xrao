// https://github.com/Zeronetsec/Xrao

package generate

import (
    "fmt"
    "os"
    "github.com/Zeronetsec/Xrao/parser"
    "github.com/Zeronetsec/Xrao/utils/color"
    "github.com/Zeronetsec/Xrao/utils/variable"
    "github.com/Zeronetsec/Xrao/utils/invinput"
)

func Runner(config, out, mode string) {
    if config == "" {
        config = variable.Config
    }

    if _, err := os.Stat(config); os.IsNotExist(err) {
        fmt.Printf(
            "%s[!] %sConfig: %s%s %snot found!\n",
            color.R, color.N, color.GG, config, color.N,
        )
        return
    }

    if out == "" {
        invinput.MissingArgument()
        return
    }

    if _, err := os.Stat(out); err == nil {
        fmt.Printf(
            "%s[!] %sFile: %s%s %sis already exist!\n",
            color.R, color.N, color.GG, out, color.N,
        )
        return
    }

    if mode == "" {
        invinput.MissingArgument()
        return
    }

    switch mode {
        case "apply":
            if err := initScript(out); err != nil {
                fmt.Printf(
                    "%s[!] %sFailed to initialize file: %s%s%s\n",
                    color.R, color.N, color.GG, err.Error(), color.N,
                )
                return
            }
            parser.ParseGlobal(config, out)
            parser.ParseApps(config, out)
        case "reset":
            if err := initScript(out); err != nil {
                fmt.Printf(
                    "%s[!] %sFailed to initialize file: %s%s%s\n",
                    color.R, color.N, color.GG, err.Error(), color.N,
                )
                return
            }
            parser.ParseReset(config, out)
        case "apply-global-only":
            if err := initScript(out); err != nil {
                fmt.Printf(
                    "%s[!] %sFailed to initialize file: %s%s%s\n",
                    color.R, color.N, color.GG, err.Error(), color.N,
                )
                return
            }
            parser.ParseGlobal(config, out)
        default:
            invinput.InvalidOption(
                fmt.Sprintf("--generate --mode %s", mode),
            )
            return
    }

    if err := os.Chmod(out, 0755); err != nil {
        fmt.Printf(
            "%s[!] %sFailed set permission to 0755: %s%s%s\n",
            color.R, color.N, color.GG, err.Error(), color.N,
        )
        return
    }

    fmt.Printf(
        "%s[*] %sScript generated: %s%s%s\n",
        color.GG, color.N, color.GG, out, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec