package console

import (
    "github.com/Zeronetsec/Xrao/module/apply"
)

type Apply struct{}
func (c Apply) Execute(args []string) {
    configPath := ""
    if len(args) >= 3 {
        configPath = args[2]
    }

    apply.Runner(configPath)
}