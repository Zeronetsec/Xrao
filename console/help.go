package console

import (
    "github.com/Zeronetsec/Xrao/module/help"
)

type Help struct{}
func (c Help) Execute(args []string) {
    help.ShowHelper()
}