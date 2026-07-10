// https://github.com/Zeronetsec/Xrao

package parser

import (
    "fmt"
    "os"
    "strings"
    "github.com/Zeronetsec/Xrao/utils/color"
)

func ParseReset(inputPath, outputPath string) {
    if _, err := os.Stat(inputPath); os.IsNotExist(err) {
        os.Exit(1)
    }

    content, err := os.ReadFile(inputPath)
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed reading file: %s%s\n",
            color.R, color.N, color.GG, err.Error(), color.N,
        )
        return
    }
    lines := strings.Split(string(content), "\n")

    insideReset := false
    resetFlags := make(map[string]bool)
    var appPackages []string

    seenPackages := make(map[string]bool)
    for _, line := range lines {
        stripped := strings.TrimSpace(line)
        if stripped == "" || strings.HasPrefix(
            stripped, "#",
        ) {
            continue
        }

        actualContent := strings.TrimSpace(
            strings.Split(stripped, "#")[0],
        )

        compactContent := strings.ReplaceAll(
            actualContent, " ", "",
        )

        if compactContent == "[Reset]" {
            insideReset = true
            fmt.Printf(
                "  %s-> %sEntering block: %s[ %sReset %s]%s\n",
                color.DG, color.N, color.DG, color.GG, color.DG, color.N,
            )
            continue
        }

        if compactContent == "[END]" && insideReset {
            insideReset = false
            continue
        }

        if insideReset && strings.HasPrefix(
            actualContent, "-",
        ) {
            parentClean := strings.TrimSpace(
                strings.TrimLeft(actualContent, "-"),
            )

            if strings.Contains(parentClean, "=") {
                parts := strings.SplitN(parentClean, "=", 2)
                key := strings.TrimSpace(parts[0])
                val := strings.TrimSpace(parts[1])

                if val == "true" {
                    resetFlags[key] = true
                }
            }
        }

        if strings.HasPrefix(
            compactContent, "[App:",
        ) && strings.HasSuffix(
            compactContent, "]",
        ) {
            parts := strings.SplitN(actualContent, ":", 2)
            if len(parts) > 1 {
                pkg := strings.TrimSpace(
                    strings.ReplaceAll(parts[1], "]", ""),
                )

                if !seenPackages[pkg] {
                    seenPackages[pkg] = true
                    appPackages = append(appPackages, pkg)
                }
            }
        }
    }

    var commands []string
    if resetFlags["reset_window_animation"] {
        fmt.Printf(
            "    %s* %sParsing: %sreset_window_animation%s\n",
            color.B, color.N, color.GG, color.N,
        )

        commands = append(
            commands,
            fmt.Sprintf(
                "echo -e \"%s[*] %sExecute: %ssettings put global %swindow_animation_scale 1.0%s\"",
                color.B, color.N, color.GG, color.CC, color.N,
            ),
            "settings put global window_animation_scale 1.0",
        )

        fmt.Printf(
            "      %s+ %sAppend: %ssettings put global %swindow_animation_scale 1.0%s\n",
            color.GG, color.N, color.GG, color.CC, color.N,
        )

        commands = append(
            commands,
            fmt.Sprintf(
                "echo -e \"%s[*] %sExecute: %ssettings put global %stransition_animation_scale 1.0%s\"",
                color.B, color.N, color.GG, color.CC, color.N,
            ),
            "settings put global transition_animation_scale 1.0",
        )

        fmt.Printf(
            "      %s+ %sAppend: %ssettings put global %stransition_animation_scale 1.0%s\n",
            color.GG, color.N, color.GG, color.CC, color.N,
        )

        commands = append(
            commands,
            fmt.Sprintf(
                "echo -e \"%s[*] %sExecute: %ssettings put global %sanimator_duration_scale 1.0%s\"",
                color.B, color.N, color.GG, color.CC, color.N,
            ),
            "settings put global animator_duration_scale 1.0",
        )

        fmt.Printf(
            "      %s+ %sAppend: %ssettings put global %sanimator_duration_scale 1.0%s\n",
            color.GG, color.N, color.GG, color.CC, color.N,
        )
    }

    if resetFlags[
        "reset_thermal_throttle",
    ] && len(appPackages) > 0 {
        fmt.Printf(
            "    %s* %sParsing: %sreset_thermal_throttle%s\n",
            color.B, color.N, color.GG, color.N,
        )

        for _, pkg := range appPackages {
            commands = append(
                commands,
                fmt.Sprintf(
                    "echo -e \"%s[*] %sExecute: %scmd game set %sthermal-throttle %s true%s\"",
                    color.B, color.N, color.GG, color.CC, pkg, color.N,
                ),

                fmt.Sprintf(
                    "echo -e \"%s[*] %sExecute: %scmd game reset %sthermal-throttle %s%s\"",
                    color.B, color.N, color.GG, color.CC, pkg, color.N,
                ),

                fmt.Sprintf(
                    "cmd game set thermal-throttle %s true",
                    pkg,
                ),

                fmt.Sprintf(
                    "cmd game reset thermal-throttle %s",
                    pkg,
                ),
            )
            fmt.Printf(
                "      %s+ %sAppend: %scmd game set %sthermal-throttle %s true%s\n",
                color.GG, color.N, color.GG, color.CC, pkg, color.N,
            )

            fmt.Printf(
                "      %s+ %sAppend: %scmd game reset %sthermal-throttle %s%s\n",
                color.GG, color.N, color.GG, color.CC, pkg, color.N,
            )
        }
    }

    if resetFlags[
        "reset_performance_mode",
    ] && len(appPackages) > 0 {
        fmt.Printf(
            "    %s* %sParsing: %sreset_performance_mode%s\n",
            color.B, color.N, color.GG, color.N,
        )

        for _, pkg := range appPackages {
            commands = append(
                commands,
                fmt.Sprintf(
                    "echo -e \"%s[*] %sExecute: %scmd game reset %s%s%s\"",
                    color.B, color.N, color.GG, color.CC, pkg, color.N,
                ),

                fmt.Sprintf(
                    "cmd game reset %s",
                    pkg,
                ),
            )
            fmt.Printf(
                "      %s+ %sAppend: %scmd game reset %s%s%s\n",
                color.GG, color.N, color.GG, color.CC, pkg, color.N,
            )
        }
    }

    if resetFlags["reset_driver"] {
        fmt.Printf(
            "    %s* %sParsing: %sreset_driver%s\n",
            color.B, color.N, color.GG, color.N,
        )

        commands = append(
            commands,
            fmt.Sprintf(
                "echo -e \"%s[*] %sExecute: %ssettings delete global %sangle_gl_driver_selection_pkgs%s\"",
                color.B, color.N, color.GG, color.CC, color.N,
            ),

            fmt.Sprintf(
                "echo -e \"%s[*] %sExecute: %ssettings delete global %sangle_gl_driver_selection_values%s\"",
                color.B, color.N, color.GG, color.CC, color.N,
            ),

            "settings delete global angle_gl_driver_selection_pkgs",
            "settings delete global angle_gl_driver_selection_values",
        )
        fmt.Printf(
            "      %s+ %sAppend: %ssettings delete global %sangle_gl_driver_selection_pkgs%s\n",
            color.GG, color.N, color.GG, color.CC, color.N,
        )

        fmt.Printf(
            "      %s+ %sAppend: %ssettings delete global %sangle_gl_driver_selection_values%s\n",
            color.GG, color.N, color.GG, color.CC, color.N,
        )
    }

    if resetFlags["reset_dynamic_updatable_driver"] {
        fmt.Printf(
            "    %s* %sParsing: %sreset_dynamic_updatable_driver%s\n",
            color.B, color.N, color.GG, color.N,
        )

        commands = append(
            commands,
            fmt.Sprintf(
                "echo -e \"%s[*] %sExecute: %ssettings delete global %supdatable_driver_production_opt_in_apps%s\"",
                color.B, color.N, color.GG, color.CC, color.N,
            ),

            "settings delete global updatable_driver_production_opt_in_apps",
        )
        fmt.Printf(
            "      %s+ %sAppend: %ssettings delete global %supdatable_driver_production_opt_in_apps%s\n",
            color.GG, color.N, color.GG, color.CC, color.N,
        )
    }

    if resetFlags["reset_touch_booster"] {
        fmt.Printf(
            "    %s* %sParsing: %sreset_touch_booster%s\n",
            color.B, color.N, color.GG, color.N,
        )

        commands = append(
            commands,
            fmt.Sprintf(
                "echo -e \"%s[*] %sExecute: %ssettings delete system %sinput_boost_game_pkg%s\"",
                color.B, color.N, color.GG, color.CC, color.N,
            ),

            "settings delete system input_boost_game_pkg",
        )
        fmt.Printf(
            "      %s+ %sAppend: %ssettings delete system %sinput_boost_game_pkg%s\n",
            color.GG, color.N, color.GG, color.CC, color.N,
        )
    }

    if resetFlags["reset_game_mode"] {
        fmt.Printf(
            "    %s* %sParsing: %sreset_game_mode%s\n",
            color.B, color.N, color.GG, color.N,
        )

        commands = append(
            commands,
            fmt.Sprintf(
                "echo -e \"%s[*] %sExecute: %ssettings delete secure %sgame_mode_list%s\"",
                color.B, color.N, color.GG, color.CC, color.N,
            ),

            "settings delete secure game_mode_list",
        )
        fmt.Printf(
            "      %s+ %sAppend: %ssettings delete secure %sgame_mode_list%s\n",
            color.GG, color.N, color.GG, color.CC, color.N,
        )
    }

    if resetFlags["reset_low_latency_mode"] {
        fmt.Printf(
            "    %s* %sParsing: %sreset_low_latency_mode%s\n",
            color.B, color.N, color.GG, color.N,
        )

        commands = append(
            commands,
            fmt.Sprintf(
                "echo -e \"%s[*] %sExecute: %ssettings delete global %sgame_low_latency_mode%s\"",
                color.B, color.N, color.GG, color.CC, color.N,
            ),

            "settings delete global game_low_latency_mode",
        )
        fmt.Printf(
            "      %s+ %sAppend: %ssettings delete global %sgame_low_latency_mode%s\n",
            color.GG, color.N, color.GG, color.CC, color.N,
        )
    }

    if resetFlags[
        "reset_whitelist_app",
    ] && len(appPackages) > 0 {
        fmt.Printf(
            "    %s* %sParsing: %sreset_whitelist_app%s\n",
            color.B, color.N, color.GG, color.N,
        )

        for _, pkg := range appPackages {
            commands = append(
                commands,
                fmt.Sprintf(
                    "echo -e \"%s[*] %sExecute: %sdumpsys deviceidle whitelist %s-%s%s\"",
                    color.B, color.N, color.GG, color.CC, pkg, color.N,
                ),

                fmt.Sprintf(
                    "dumpsys deviceidle whitelist -%s",
                    pkg,
                ),
            )
            fmt.Printf(
                "      %s+ %sAppend: %sdumpsys deviceidle whitelist %s-%s%s\n",
                color.GG, color.N, color.GG, color.CC, pkg, color.N,
            )
        }
    }

    for len(commands) > 0 && commands[
        len(commands)-1,
    ] == "" {
        commands = commands[:len(commands)-1]
    }

    if len(commands) > 0 {
        fileOut, err := os.OpenFile(
            outputPath,
            os.O_APPEND|os.O_CREATE|os.O_WRONLY,
            0644,
        )

        if err != nil {
            fmt.Printf(
                "%s[!] %sFailed create output: %s%s%s\n",
                color.R, color.N, color.GG, err.Error(), color.N,
            )
            return
        }
        defer fileOut.Close()

        validCmdCount := 0
        for _, cmd := range commands {
            if _, err := fileOut.WriteString(cmd + "\n"); err != nil {
                fmt.Printf(
                    "%s[!] %sFailed write output: %s%s%s\n",
                    color.R, color.N, color.GG, err.Error(), color.N,
                )
                return
            }

            if strings.TrimSpace(cmd) != "" {
                validCmdCount++
            }
        }
        fmt.Printf(
            "  %s<- %sLeaving block: %s[ %sReset %s]%s\n",
            color.DG, color.N, color.DG, color.GG, color.DG, color.N,
        )

        fmt.Printf(
            "%s[+] %sAdded: %s%d %stweaks to %s%s%s\n",
            color.GG, color.N, color.GG, validCmdCount,
            color.N, color.GG, outputPath, color.N,
        )
    }
}

// Copyright (c) 2026 Zeronetsec