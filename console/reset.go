// https://github.com/Zeronetsec/Xrao

package console

import (
    "github.com/Zeronetsec/Xrao/module/reset"
)

type Reset struct{}
func (c Reset) Execute(args []string) {
    configPath := ""
    if len(args) >= 3 {
        configPath = args[2]
    }

    reset.Runner(configPath)
}

// Copyright (c) 2026 Zeronetsec