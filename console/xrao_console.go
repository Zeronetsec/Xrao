// https://github.com/Zeronetsec/Xrao

package console

import (
    "os"
    "github.com/Zeronetsec/Xrao/utils/invinput"
)

func XraoConsole(input string) {
    args := os.Args
    if len(args) < 2 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    commands := map[string]Command{
        "--uwu": Uwu{},
        "--help": Help{},
        "--version": Version{},
        "--status": Status{},
        "--generate": Generate{},
        "--apply": Apply{},
        "--apply-global-only": ApplyGlobalOnly{},
        "--reset": Reset{},
        "--show-config": ShowConfig{},
        "--pair": Pair{},
        "--connect": Connect{},
        "--disconnect": Disconnect{},
    }

    if cmd, ok := commands[args[1]]; ok {
        cmd.Execute(args)
    } else {
        invinput.InvalidOption(args[1])
        os.Exit(1)
    }
}

// Copyright (c) 2026 Zeronetsec