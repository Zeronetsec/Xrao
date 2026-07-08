package console

import (
    "github.com/Zeronetsec/Xrao/module/connect"
    "github.com/Zeronetsec/Xrao/utils/invinput"
)

type Connect struct{}
func (c Connect) Execute(args []string) {
    if len(args) < 3 || args[2] == "" {
        invinput.MissingArgument()
        return
    }

    connect.Runner(args[2])
}