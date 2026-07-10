// https://github.com/Zeronetsec/Xrao

package console

type Command interface {
    Execute(args []string)
}

// Copyright (c) 2026 Zeronetsec