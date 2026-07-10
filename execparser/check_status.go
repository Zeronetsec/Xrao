// https://github.com/Zeronetsec/Xrao

package execparser

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "github.com/Zeronetsec/Xrao/utils/color"
    "github.com/Zeronetsec/Xrao/utils/variable"
)

func CheckStatus(inputPath string) {
    if _, err := os.Stat(inputPath); os.IsNotExist(err) {
        fmt.Printf(
            "%s[!] %sFile: %s%s %snot found!\n",
            color.R, color.N, color.GG, inputPath, color.N,
        )
        os.Exit(1)
    }

    fmt.Printf(
        "%s[*] %sXrao tweak state: %s%s/state.lock%s\n",
        color.B, color.N, color.GG, variable.Dlt, color.N,
    )

    lockCheck := runAdbCmd(
        fmt.Sprintf(
            "ls %s/state.lock 2>/dev/null",
            variable.Dlt,
        ),
    )

    if strings.Contains(lockCheck, "state.lock") {
        fmt.Printf(
            "%s* %sStatus: %sRunning %s(%sActive%s)%s\n",
            color.DG, color.N, color.GG,
            color.DG, color.CC, color.DG, color.N,
        )
    } else {
        fmt.Printf(
            "%s* %sStatus: %sStopped%s\n",
            color.DG, color.N, color.R, color.N,
        )
    }

    file, err := os.Open(inputPath)
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed open file: %s%s%s\n",
            color.R, color.N, color.GG, err.Error(), color.N,
        )
        os.Exit(1)
    }
    defer file.Close()

    var appPackages []string
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        stripped := strings.TrimSpace(line)
        if stripped == "" || strings.HasPrefix(
            stripped, "#",
        ) {
                continue
        }

        actual := strings.TrimSpace(
            strings.Split(stripped, "#")[0],
        )
        compact := strings.ReplaceAll(actual, " ", "")

        if strings.HasPrefix(
            compact, "[App:",
        ) && strings.HasSuffix(
            compact, "]",
        ) {
            parts := strings.SplitN(actual, ":", 2)
            if len(parts) > 1 {
                pkg := strings.TrimSpace(
                    strings.ReplaceAll(
                        parts[1], "]", "",
                    ),
                )

                exists := false
                for _, p := range appPackages {
                    if p == pkg {
                        exists = true
                        break
                    }
                }

                if !exists {
                    appPackages = append(
                        appPackages, pkg,
                    )
                }
            }
        }
    }

    if len(appPackages) == 0 {
        fmt.Printf(
            "%s[!] %sNo apps are listed in the config!\n",
            color.R, color.N,
        )
        return
    }

    fmt.Println()
    fmt.Printf(
        "%s[*] %sGlobal window animation:\n",
        color.B, color.N,
    )

    winAnim := runAdbCmd(
        "settings get global window_animation_scale",
    )

    transAnim := runAdbCmd(
        "settings get global transition_animation_scale",
    )

    animator := runAdbCmd(
        "settings get global animator_duration_scale",
    )

    fmt.Printf(
        "%s* %sWindow Animation Scale: %s%s\n",
        color.DG, color.N,
        withDefault(winAnim, "default"),
        color.N,
    )

    fmt.Printf(
        "%s* %sTransition Animation Scale: %s%s\n",
        color.DG, color.N,
        withDefault(transAnim, "default"),
        color.N,
    )

    fmt.Printf(
        "%s* %sAnimator Duration Scale: %s%s\n",
        color.DG, color.N,
        withDefault(animator, "default"),
        color.N,
    )

    anglePkgs := runAdbCmd(
        "settings get global angle_gl_driver_selection_pkgs",
    )

    angleVal := runAdbCmd(
        "settings get global angle_gl_driver_selection_values",
    )

    updatablePkgs := runAdbCmd(
        "settings get global updatable_driver_production_opt_in_apps",
    )

    boostPkgs := runAdbCmd(
        "settings get system input_boost_game_pkg",
    )

    gmodePkgs := runAdbCmd(
        "settings get secure game_mode_list",
    )

    lowLatPkgs := runAdbCmd(
        "settings get global game_low_latency_mode",
    )

    whitelistDump := runAdbCmd(
        "dumpsys deviceidle whitelist",
    )

    gameDump := runAdbCmd("dumpsys game")

    for _, pkg := range appPackages {
        fmt.Println()
        fmt.Printf(
            "%s[*] %sApp: %s%s%s\n",
            color.B, color.N, color.GG, pkg, color.N,
        )

        fmt.Printf(
            "%s* %sAndroid game service dump:\n",
            color.DG, color.N,
        )

        foundDump := false
        gameDumpLines := strings.Split(gameDump, "\n")
        for _, line := range gameDumpLines {
            if strings.Contains(line, pkg) {
                fmt.Printf(
                    "    %s%s%s\n",
                    color.GG,
                    strings.TrimSpace(line),
                    color.N,
                )
                foundDump = true
            }
        }
        if !foundDump {
            fmt.Printf(
                "    %sEmpty dumpsys log%s\n",
                color.R, color.N,
            )
        }

        modes := runAdbCmd(
            fmt.Sprintf(
                "cmd game list-modes %s", pkg,
            ),
        )

        configs := runAdbCmd(
            fmt.Sprintf(
                "cmd game list-configs %s", pkg,
            ),
        )

        fmt.Printf(
            "%s* %sGame modes available: %s%s\n",
            color.DG, color.N,
            withDefault(modes, "none/default"),
            color.N,
        )

        fmt.Printf(
            "%s* %sGame configs active: %s%s\n",
            color.DG, color.N,
            withDefault(configs, "none/default"),
            color.N,
        )

        fmt.Printf(
            "%s* %sSystem tweaks state:\n",
            color.DG, color.N,
        )

        isAngle := strings.Contains(anglePkgs, pkg)
        if isAngle {
            fmt.Printf(
                "  %s- %sANGLE driver active: %strue %s(%s%s%s)%s\n",
                color.DG, color.N, color.GG,
                color.DG, color.CC, angleVal, color.DG, color.N,
            )
        } else {
            fmt.Printf(
                "  %s- %sANGLE driver active: %sfalse%s\n",
                color.DG, color.N, color.R, color.N,
            )
        }

        isUpdatable := strings.Contains(updatablePkgs, pkg)
        fmt.Printf(
            "  %s- %sUpdatable driver prod: %s%s\n",
            color.DG, color.N,
            mapBoolColor(isUpdatable),
            color.N,
        )

        isBoost := strings.Contains(boostPkgs, pkg)
        fmt.Printf(
            "  %s- %sGame touch booster: %s%s\n",
            color.DG, color.N,
            mapBoolColor(isBoost), color.N,
        )

        isGmode := strings.Contains(gmodePkgs, pkg)
        fmt.Printf(
            "  %s- %sAndroid secure game mode: %s%s\n",
            color.DG, color.N,
            mapBoolColor(isGmode),
            color.N,
        )

        isLowLat := strings.Contains(lowLatPkgs, pkg)
        fmt.Printf(
            "  %s- %sLow latency mode: %s%s\n",
            color.DG, color.N,
            mapBoolColor(isLowLat),
            color.N,
        )

        isWhitelisted := strings.Contains(whitelistDump, pkg)
        fmt.Printf(
            "  %s- %sDoze mode whitelisted: %s%s\n",
            color.DG, color.N,
            mapBoolColor(isWhitelisted),
            color.N,
        )
    }
}

// Copyright (c) 2026 Zeronetsec