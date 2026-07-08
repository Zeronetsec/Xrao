package console

import (
    "github.com/Zeronetsec/Xrao/module/version"
)

type Version struct{}
func (c Version) Execute(args []string) {
    version.ShowVersion()
}