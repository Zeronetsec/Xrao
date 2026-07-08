package agonly

import (
    "fmt"
    "os"
    "path/filepath"
    "github.com/Zeronetsec/Xrao/module/generate"
    "github.com/Zeronetsec/Xrao/utils/color"
    "github.com/Zeronetsec/Xrao/utils/shell"
    "github.com/Zeronetsec/Xrao/utils/variable"
    "github.com/Zeronetsec/Xrao/utils/move"
)

func Runner(configPath string) {
    if _, err := shell.Execf("adb get-state"); err != nil {
        fmt.Printf(
            "%s[!] %sADB device not detected!\n",
            color.R, color.N,
        )
        return
    }

    if configPath == "" {
        configPath = variable.Config
    }

    if _, err := os.Stat(configPath); os.IsNotExist(err) {
        fmt.Printf(
            "%s[!] %sFile: %s%s %snot found!\n",
            color.R, color.N, color.GG, configPath, color.N,
        )
        return
    }

    _, _ = shell.Execf("adb start-server")

    fmt.Printf(
        "%s[*] %sEnsuring directory:\n",
        color.B, color.N,
    )

    fmt.Printf(
        "%s* %s%s%s\n",
        color.DG, color.GG, variable.Sdx, color.N,
    )
    _ = os.MkdirAll(variable.Sdx, 0755)

    localTmpDir := filepath.Join(variable.Tmp, "xrao")
    fmt.Printf(
        "%s* %s%s%s\n",
        color.DG, color.GG, localTmpDir, color.N,
    )
    _ = os.MkdirAll(localTmpDir, 0755)

    localScript := filepath.Join(localTmpDir, "xrao_gonly.sh")
    sdcardScript := filepath.Join(variable.Sdx, "xrao_gonly.sh")

    if _, err := os.Stat(localScript); err == nil {
        fmt.Printf(
            "%s[*] %sRemoving: %s%s%s\n",
            color.B, color.N, color.GG, localScript, color.N,
        )
        _ = os.Remove(localScript)
    }

    generate.Runner(
        configPath, localScript, "apply-global-only",
    )

    if _, err := os.Stat(sdcardScript); err == nil {
        fmt.Printf(
            "%s[*] %sRemoving: %s%s%s\n",
            color.B, color.N, color.GG, sdcardScript, color.N,
        )
        _ = os.Remove(sdcardScript)
    }

    fmt.Printf(
        "%s[*] %sMoving: %s%s %s-> %s%s/%s\n",
        color.B, color.N, color.GG, localScript,
        color.DG, color.GG, variable.Sdx, color.N,
    )

    if err := move.Fmove(localScript, sdcardScript); err != nil {
        fmt.Printf(
            "%s[!] %sFailed moving script: %s%s%s\n",
            color.R, color.N, color.GG, err.Error(), color.N,
        )
        return
    }

    fmt.Printf(
        "%s[*] %sRemoving: %s/data/local/tmp/xrao_gonly.sh %sif exist\n",
        color.B, color.N, color.GG, color.N,
    )
    _, _ = shell.Execf(
        "adb shell \"if [ -f /data/local/tmp/xrao_gonly.sh ]; then rm /data/local/tmp/xrao_gonly.sh; fi\"",
    )

    fmt.Printf(
        "%s[*] %sCopying: %s/sdcard/Download/xrao/xrao_gonly.sh %s-> %s/data/local/tmp/%s\n",
        color.B, color.N, color.GG, color.DG, color.GG, color.N,
    )
    _, _ = shell.Execf(
        "adb shell \"cp %s/xrao_gonly.sh /data/local/tmp/\"",
        variable.Sdx,
    )

    fmt.Printf(
        "%s[*] %sRemoving: %s/sdcard/Download/xrao/xrao_gonly.sh%s\n",
        color.B, color.N, color.GG, color.N,
    )
    _ = os.Remove(sdcardScript)

    fmt.Printf(
        "%s[*] %sSet permission for: %s/data/local/tmp/xrao_gonly.sh%s\n",
        color.B, color.N, color.GG, color.N,
    )
    _, _ = shell.Execf(
        "adb shell \"chmod +x /data/local/tmp/xrao_gonly.sh\"",
    )

    fmt.Printf(
        "%s[*] %sCall: %s/system/bin/sh %s-> %s/data/local/tmp/xrao_gonly.sh%s\n",
        color.B, color.N, color.GG, color.DG, color.GG, color.N,
    )
    execErr := shell.ExecLivef(
        "adb shell \"/system/bin/sh /data/local/tmp/xrao_gonly.sh\"",
    )

    if execErr != nil {
        fmt.Printf(
            "%s[!] %sFailed executing script: %s%s%s\n",
            color.R, color.N, color.GG, execErr.Error(), color.N,
        )
    }

    fmt.Printf(
        "%s[*] %sEnsuring: %s/data/local/tmp/state.lock%s\n",
        color.B, color.N, color.GG, color.N,
    )
    _, _ = shell.Execf(
        "adb shell \"if [ ! -f /data/local/tmp/state.lock ]; then touch /data/local/tmp/state.lock; fi\"",
    )

    fmt.Printf(
        "%s[*] %sCleaning up:\n",
        color.B, color.N,
    )

    fmt.Printf(
        "%s* %s/data/local/tmp/xrao_gonly.sh%s\n",
        color.DG, color.GG, color.N,
    )
    _, _ = shell.Execf(
        "adb shell \"rm /data/local/tmp/xrao_gonly.sh\"",
    )

    fmt.Printf(
        "%s* %s%s%s\n",
        color.DG, color.GG, variable.Sdx, color.N,
    )
    _ = os.RemoveAll(variable.Sdx)

    fmt.Printf(
        "%s* %s%s%s\n",
        color.DG, color.GG, localTmpDir, color.N,
    )
    _ = os.RemoveAll(localTmpDir)

    fmt.Println()
    fmt.Printf(
        "%s[+] %sXrao is running now...\n",
        color.GG, color.N,
    )
}