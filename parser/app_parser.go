// https://github.com/Zeronetsec/Xrao

package parser

import (
    "fmt"
    "os"
    "regexp"
    "strings"
    "github.com/Zeronetsec/Xrao/utils/color"
)

func ParseApps(inputPath, outputPath string) {
    if _, err := os.Stat(inputPath); os.IsNotExist(err) {
        os.Exit(1)
    }

    content, err := os.ReadFile(inputPath)
    if err != nil {
        fmt.Printf(
            "%s[!] %sFailed reading file: %s%s%s\n",
            color.R, color.N, color.GG, err.Error(), color.N,
        )
        return
    }
    lines := strings.Split(string(content), "\n")

    insideApp := false
    currentAppPkg := ""
    currentParent := ""

    var commandsImmediate []string
    var anglePackages []string
    var updatableDriverPkgs []string
    var touchBoosterPkgs []string
    var gameModePkgs []string
    var lowLatencyPkgs []string
    var whitelistPkgs []string

    perfProps := map[string]string{
        "mode": "",
        "downscale": "",
        "fps": "",
    }
    angleValue := "default"

    insideGlobalScan := false
    for _, lineScan := range lines {
        strippedScan := strings.TrimSpace(lineScan)
        if strippedScan == "" || strings.HasPrefix(
            strippedScan, "#",
        ) {
            continue
        }

        actualScan := strings.TrimSpace(
            strings.Split(strippedScan, "#")[0],
        )

        compactScan := strings.ReplaceAll(actualScan, " ", "")
        if compactScan == "[Global]" {
            insideGlobalScan = true
            continue
        }

        if compactScan == "[END]" && insideGlobalScan {
            break
        }

        if insideGlobalScan && strings.HasPrefix(actualScan, "-") {
            parentClean := strings.TrimSpace(
                strings.TrimLeft(actualScan, "-"),
            )

            if strings.HasPrefix(
                parentClean, "driver",
            ) && strings.Contains(
                parentClean, "=",
            ) {
                parts := strings.SplitN(
                    parentClean, "=", 2,
                )
                angleValue = strings.TrimSpace(parts[1])
            }
        }
    }

    numericRegex := regexp.MustCompile(
        `^[0-9]+(\.[0-9]+)?$`,
    )

    for i, line := range lines {
        lineNumber := i + 1
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

        if strings.HasPrefix(
            compactContent, "[App:",
        ) && strings.HasSuffix(
            compactContent, "]",
            ) {
            insideApp = true
            parts := strings.SplitN(actualContent, ":", 2)
            if len(parts) > 1 {
                currentAppPkg = strings.TrimSpace(
                    strings.ReplaceAll(
                        parts[1], "]", "",
                    ),
                )

                fmt.Printf(
                    "  %s-> %sEntering block: %s[ %sApp: %s%s %s]%s\n",
                    color.DG, color.N, color.DG, color.GG,
                    color.CC, currentAppPkg, color.DG, color.N,
                )
            }
            continue
        }

        if compactContent == "[END]" && insideApp {
            if perfProps[
                "mode",
            ] != "" && perfProps[
                "downscale",
            ] != "" && perfProps[
                "fps",
            ] != "" {
                cmdEcho := fmt.Sprintf(
                    "echo -e \"%s[*] %sExecute: %scmd game set --mode %s%s %s--downscale %s%s %s--fps %s%s %s%s\"",
                    color.B, color.N, color.GG, color.CC,
                    perfProps["mode"],
                    color.GG, color.CC,
                    perfProps["downscale"],
                    color.GG, color.CC,
                    perfProps["fps"],
                    currentAppPkg,
                    color.N,
                )

                cmdRaw := fmt.Sprintf(
                    "cmd game set --mode %s --downscale %s --fps %s %s",
                    perfProps["mode"],
                    perfProps["downscale"],
                    perfProps["fps"],
                    currentAppPkg,
                )

                commandsImmediate = append(
                    commandsImmediate, cmdEcho, cmdRaw,
                )

                fmt.Printf(
                    "      %s+ %sAppend: %scmd game set --mode %s%s %s--downscale %s%s %s--fps %s%s %s%s\n",
                    color.GG, color.N, color.GG, color.CC,
                    perfProps["mode"],
                    color.GG, color.CC,
                    perfProps["downscale"],
                    color.GG, color.CC,
                    perfProps["fps"],
                    currentAppPkg,
                    color.N,
                )
            }
            perfProps = map[string]string{
                "mode": "",
                "downscale": "",
                "fps": "",
            }

            insideApp = false
            fmt.Printf(
                "  %s<- %sLeaving block: %s[ %sApp: %s%s %s]%s\n",
                color.DG, color.N, color.DG, color.GG, color.CC,
                currentAppPkg, color.DG, color.N,
            )
            continue
        }

        if !insideApp {
            continue
        }

        if strings.HasPrefix(
            line, " ",
        ) || strings.HasPrefix(
            line, "\t",
        ) {
            childClean := strings.TrimSpace(
                strings.TrimLeft(actualContent, "-"),
            )

            if currentParent == "performance_mode" {
                delim := " "
                if strings.Contains(childClean, "=") {
                    delim = "="
                }
                parts := strings.SplitN(childClean, delim, 2)
                if len(parts) == 2 {
                    k := strings.TrimSpace(parts[0])
                    v := strings.TrimSpace(parts[1])

                    if _, exists := perfProps[k]; exists {
                        if numericRegex.MatchString(v) {
                            perfProps[k] = v
                        } else {
                            fmt.Printf(
                                "%s[!] %sLine %d: Sub-properti %s%s %smust be number!\n",
                                color.YY, color.N, lineNumber, color.GG, k, color.N,
                            )

                            fmt.Printf(
                                "%s[!] %sFound: %s%s%s\n",
                                color.YY, color.N, color.GG, v, color.N,
                            )

                            fmt.Printf(
                                "%s[!] %sSkipping...\n",
                                color.YY, color.N,
                            )
                        }
                    }
                }
            }
        } else if strings.HasPrefix(actualContent, "-") {
            parentClean := strings.TrimSpace(
                strings.TrimLeft(actualContent, "-"),
            )

            var key, val string
            if strings.Contains(parentClean, "=") {
                parts := strings.SplitN(parentClean, "=", 2)
                key = strings.TrimSpace(parts[0])
                val = strings.TrimSpace(parts[1])
            } else {
                key = strings.TrimSpace(parentClean)
                val = ""
            }
            currentParent = key

            if val == "true" {
                switch key {
                    case "performance_mode":
                        fmt.Printf(
                            "    %s* %sParsing: %sperformance_mode%s\n",
                            color.B, color.N, color.GG, color.N,
                        )
                    case "enable_angle":
                        fmt.Printf(
                            "    %s* %sParsing: %senable_angle%s\n",
                            color.B, color.N, color.GG, color.N,
                        )

                        anglePackages = append(
                            anglePackages, currentAppPkg,
                        )
                    case "disable_thermal_throttle":
                        fmt.Printf(
                            "    %s* %sParsing: %sdisable_thermal_throttle%s\n",
                            color.B, color.N, color.GG, color.N,
                        )

                        cmdEcho := fmt.Sprintf(
                            "echo -e \"%s[*] %sExecute: %scmd game set thermal-throttle %s%s false%s\"",
                            color.B, color.N, color.GG, color.CC, currentAppPkg, color.N,
                        )

                        cmdRaw := fmt.Sprintf(
                            "cmd game set thermal-throttle %s false",
                            currentAppPkg,
                        )

                        commandsImmediate = append(
                            commandsImmediate, cmdEcho, cmdRaw,
                        )

                        fmt.Printf(
                            "      %s+ %sAppend: %scmd game set thermal-throttle %s%s false%s\n",
                            color.GG, color.N, color.GG, color.CC, currentAppPkg, color.N,
                        )
                    case "enable_dynamic_updatable_driver":
                        fmt.Printf(
                            "    %s* %sParsing: %senable_dynamic_updatable_driver%s\n",
                            color.B, color.N, color.GG, color.N,
                        )

                        updatableDriverPkgs = append(
                            updatableDriverPkgs, currentAppPkg,
                        )
                    case "enable_touch_booster":
                        fmt.Printf(
                            "    %s* %sParsing: %senable_touch_booster%s\n",
                            color.B, color.N, color.GG, color.N,
                        )

                        touchBoosterPkgs = append(
                            touchBoosterPkgs, currentAppPkg,
                        )
                    case "enable_game_mode":
                        fmt.Printf(
                            "    %s* %sParsing: %senable_game_mode%s\n",
                            color.B, color.N, color.GG, color.N,
                        )

                        gameModePkgs = append(
                            gameModePkgs, currentAppPkg,
                        )
                    case "enable_low_latency_mode":
                        fmt.Printf(
                            "    %s* %sParsing: %senable_low_latency_mode%s\n",
                            color.B, color.N, color.GG, color.N,
                        )

                        lowLatencyPkgs = append(
                            lowLatencyPkgs, currentAppPkg,
                        )
                    case "whitelist_apps":
                        fmt.Printf(
                            "    %s* %sParsing: %swhitelist_apps%s\n",
                            color.B, color.N, color.GG, color.N,
                        )

                        whitelistPkgs = append(
                            whitelistPkgs, currentAppPkg,
                        )
                    default:
                        currentParent = ""
                }
            }
        }
    }

    var commandsGrouped []string
    if len(anglePackages) > 0 {
        joinPkgs := strings.Join(anglePackages, ",")
        commandsGrouped = append(
            commandsGrouped,
            fmt.Sprintf(
                "echo -e \"%s[*] %sExecute: %ssettings put global %sangle_gl_driver_selection_pkgs %s%s\"",
                color.B, color.N, color.GG, color.CC, joinPkgs, color.N,
            ),

            fmt.Sprintf(
                "echo -e \"%s[*] %sExecute: %ssettings put global %sangle_gl_driver_selection_values %s%s\"",
                color.B, color.N, color.GG, color.CC, angleValue, color.N,
            ),

            fmt.Sprintf(
                "settings put global angle_gl_driver_selection_pkgs %s",
                joinPkgs,
            ),

            fmt.Sprintf(
                "settings put global angle_gl_driver_selection_values %s",
                angleValue,
            ),
        )
        fmt.Printf(
            "    %s+ %sAppend: %ssettings put global %sangle_gl_driver_selection_pkgs %s%s\n",
            color.GG, color.N, color.GG, color.CC, joinPkgs, color.N,
        )

        fmt.Printf(
            "    %s+ %sAppend: %ssettings put global %sangle_gl_driver_selection_values %s%s\n",
            color.GG, color.N, color.GG, color.CC, angleValue, color.N,
        )
    }

    if len(updatableDriverPkgs) > 0 {
        joinPkgs := strings.Join(updatableDriverPkgs, ",")
        commandsGrouped = append(
            commandsGrouped,
            fmt.Sprintf(
                "echo -e \"%s[*] %sExecute: %ssettings put global %supdatable_driver_production_opt_in_apps %s%s\"",
                color.B, color.N, color.GG, color.CC, joinPkgs, color.N,
            ),

            fmt.Sprintf(
                "settings put global updatable_driver_production_opt_in_apps %s",
                joinPkgs,
            ),
        )
        fmt.Printf(
            "    %s+ %sAppend: %ssettings put global %supdatable_driver_production_opt_in_apps %s%s\n",
            color.GG, color.N, color.GG, color.CC, joinPkgs, color.N,
        )
    }

    if len(touchBoosterPkgs) > 0 {
        joinPkgs := strings.Join(touchBoosterPkgs, ",")
        commandsGrouped = append(
            commandsGrouped,
            fmt.Sprintf(
                "echo -e \"%s[*] %sExecute: %ssettings put system %sinput_boost_game_pkg %s%s\"",
                color.B, color.N, color.GG, color.CC, joinPkgs, color.N,
            ),

            fmt.Sprintf(
                "settings put system input_boost_game_pkg %s",
                joinPkgs,
            ),
        )
        fmt.Printf(
            "    %s+ %sAppend: %ssettings put system %sinput_boost_game_pkg %s%s\n",
            color.GG, color.N, color.GG, color.CC, joinPkgs, color.N,
        )
    }

    if len(gameModePkgs) > 0 {
        joinPkgs := strings.Join(gameModePkgs, ",")
        commandsGrouped = append(
            commandsGrouped,
            fmt.Sprintf(
                "echo -e \"%s[*] %sExecute: %ssettings put secure %sgame_mode_list %s%s\"",
                color.B, color.N, color.GG, color.CC, joinPkgs, color.N,
            ),

            fmt.Sprintf(
                "settings put secure game_mode_list %s",
                joinPkgs,
            ),
        )
        fmt.Printf(
            "    %s+ %sAppend: %ssettings put secure %sgame_mode_list %s%s\n",
            color.GG, color.N, color.GG, color.CC, joinPkgs, color.N,
        )
    }

    if len(lowLatencyPkgs) > 0 {
        joinPkgs := strings.Join(lowLatencyPkgs, ",")
        commandsGrouped = append(
            commandsGrouped,
            fmt.Sprintf(
                "echo -e \"%s[*] %sExecute: %ssettings put global %sgame_low_latency_mode %s%s\"",
                color.B, color.N, color.GG, color.CC, joinPkgs, color.N,
            ),

            fmt.Sprintf(
                "settings put global game_low_latency_mode %s",
                joinPkgs,
            ),
        )
        fmt.Printf(
            "    %s+ %sAppend: %ssettings put global %sgame_low_latency_mode %s%s\n",
            color.GG, color.N, color.GG, color.CC, joinPkgs, color.N,
        )
    }

    for _, pkg := range whitelistPkgs {
        commandsGrouped = append(
            commandsGrouped,
            fmt.Sprintf(
                "echo -e \"%s[*] %sExecute: %sdumpsys deviceidle whitelist %s+%s%s\"",
                color.B, color.N, color.GG, color.CC, pkg, color.N,
            ),

            fmt.Sprintf(
                "dumpsys deviceidle whitelist +%s",
                pkg,
            ),
        )
        fmt.Printf(
            "    %s+ %sAppend: %sdumpsys deviceidle whitelist %s+%s%s\n",
            color.GG, color.N, color.GG, color.CC, pkg, color.N,
        )
    }

    finalOutput := append(commandsImmediate, commandsGrouped...)
    if len(finalOutput) > 0 {
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

        for _, cmd := range finalOutput {
            if _, err := fileOut.WriteString(cmd + "\n"); err != nil {
                fmt.Printf(
                    "%s[!] %sFailed write output: %s%s%s\n",
                    color.R, color.N, color.GG, err.Error(), color.N,
                )
                return
            }
        }
        fmt.Printf(
            "%s[+] %sAdded: %s%d %stweaks to %s%s%s\n",
            color.GG, color.N, color.GG,
            len(finalOutput), color.N, color.GG, outputPath, color.N,
        )
    }
}

// Copyright (c) 2026 Zeronetsec