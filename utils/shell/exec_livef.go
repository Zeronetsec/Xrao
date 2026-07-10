// https://github.com/Zeronetsec/Xrao

package shell

import (
    "fmt"
    "os"
    "os/exec"
    "github.com/google/shlex"
)

func ExecLivef(format string, a ...interface{}) error {
    cmdStr := fmt.Sprintf(format, a...)
    args, err := shlex.Split(cmdStr)
    if err != nil {
        return fmt.Errorf(
            "Split error: %w", err,
        )
    }

    if len(args) == 0 {
        return fmt.Errorf(
            "Empty command!",
        )
    }

    cmd := exec.Command(args[0], args[1:]...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    return cmd.Run()
}

// Copyright (c) 2026 Zeronetsec