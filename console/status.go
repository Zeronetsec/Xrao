// https://github.com/Zeronetsec/Xrao

package console

import (
    "github.com/Zeronetsec/Xrao/module/status"
)

type Status struct{}
func (c Status) Execute(args []string) {
    configPath := ""
    if len(args) >= 3 {
        configPath = args[2]
    }

    status.Show(configPath)
}

// Copyright (c) 2026 Zeronetsec