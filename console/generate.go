// https://github.com/Zeronetsec/Xrao

package console

import (
    "fmt"
    "os"
    "github.com/Zeronetsec/Xrao/module/generate"
    "github.com/Zeronetsec/Xrao/utils/invinput"
)

type Generate struct{}
func (c Generate) Execute(args []string) {
    var configPath string
    var outputPath string
    var runMode string

    for i := 2; i < len(args); i++ {
        switch args[i] {
            case "--config":
                if i+1 < len(args) {
                    configPath = args[i+1]
                    i++
                } else {
                    invinput.MissingArgument()
                    os.Exit(1)
                }
            case "--out":
                if i+1 < len(args) {
                    outputPath = args[i+1]
                    i++
                } else {
                    invinput.MissingArgument()
                    os.Exit(1)
                }
            case "--mode":
                if i+1 < len(args) {
                    runMode = args[i+1]
                    i++
                } else {
                    invinput.MissingArgument()
                    os.Exit(1)
                }
            default:
                invinput.InvalidOption(
                    fmt.Sprintf("--generate %s", args[i]),
                )
                os.Exit(1)
        }
    }

    generate.Runner(configPath, outputPath, runMode)
}

// Copyright (c) 2026 Zeronetsec