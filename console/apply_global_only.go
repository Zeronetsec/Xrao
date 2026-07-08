package console

import (
    "github.com/Zeronetsec/Xrao/module/agonly"
)

type ApplyGlobalOnly struct{}
func (c ApplyGlobalOnly) Execute(args []string) {
    configPath := ""
    if len(args) >= 3 {
        configPath = args[2]
    }

    agonly.Runner(configPath)
}