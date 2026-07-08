package console

import (
    "os"
    "github.com/Zeronetsec/Xrao/module/pair"
    "github.com/Zeronetsec/Xrao/utils/invinput"
)

type Pair struct{}
func (c Pair) Execute(args []string) {
    if len(args) < 3 || args[2] == "" {
        invinput.MissingArgument()
        os.Exit(1)
    }

    pair.Runner(args[2])
}