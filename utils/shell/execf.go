// https://github.com/Zeronetsec/Xrao

package shell

import (
    "bytes"
    "fmt"
    "strings"
    "os/exec"
    "github.com/google/shlex"
)

func Execf(format string, a ...interface{}) (string, error) {
    cmdStr := fmt.Sprintf(format, a...)
    args, err := shlex.Split(cmdStr)
    if err != nil {
        return "", fmt.Errorf(
            "Split error: %w", err,
        )
    }

    if len(args) == 0 {
        return "", fmt.Errorf(
            "Empty command!",
        )
    }

    cmd := exec.Command(args[0], args[1:]...)

    var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    err = cmd.Run()
    if err != nil {
        if stderr.Len() > 0 {
            return "", fmt.Errorf(
                "Command failed: %w (stderr: %s)",
                err, strings.TrimSpace(stderr.String()),
            )
        }
        return "", err
    }
    return stdout.String(), nil
}

// Copyright (c) 2026 Zeronetsec