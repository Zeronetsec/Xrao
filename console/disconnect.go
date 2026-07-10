// https://github.com/Zeronetsec/Xrao

package console

import (
    "github.com/Zeronetsec/Xrao/module/disconnect"
)

type Disconnect struct{}
func (c Disconnect) Execute(args []string) {
    disconnect.Runner()
}

// Copyright (c) 2026 Zeronetsec