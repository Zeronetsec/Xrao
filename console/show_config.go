package console

import (
    "github.com/Zeronetsec/Xrao/module/showconfig"
)

type ShowConfig struct{}
func (c ShowConfig) Execute(args []string) {
    configPath := ""
    if len(args) >= 3 {
        configPath = args[2]
    }

    showconfig.Runner(configPath)
}