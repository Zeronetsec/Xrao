package execparser

import (
    "strings"
    "github.com/Zeronetsec/Xrao/utils/shell"
)

func runAdbCmd(cmd string) string {
    out, err := shell.Execf(
        "adb shell %q", cmd,
    )
    if err != nil {
        return ""
    }
    return strings.TrimSpace(out)
}