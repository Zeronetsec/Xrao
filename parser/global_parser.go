package parser

import (
    "fmt"
    "os"
    "regexp"
    "strings"
    "github.com/Zeronetsec/Xrao/utils/color"
)

func ParseGlobal(inputPath, outputPath string) {
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

    insideGlobal := false
    currentParent := ""
    var commands []string

    numericRegex := regexp.MustCompile(
        `^[0-9]+(\.[0-9]+)?$`,
    )

    packageRegex := regexp.MustCompile(
        `^[a-zA-Z0-9_\.]+$`,
    )

    mapping := map[string]string{
        "window_animation": "window_animation_scale",
        "window_transition": "transition_animation_scale",
        "window_animator": "animator_duration_scale",
    }

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

        if compactContent == "[Global]" {
            insideGlobal = true
            fmt.Printf(
                "  %s-> %sEntering block: %s[ %sGlobal %s]%s\n",
                color.DG, color.N, color.DG, color.GG, color.DG, color.N,
            )
            continue
        }

        if compactContent == "[END]" && insideGlobal {
            insideGlobal = false
            fmt.Printf(
                "  %s<- %sLeaving block: %s[ %sGlobal %s]%s\n",
                color.DG, color.N, color.DG, color.GG, color.DG, color.N,
            )
            break
        }

        if !insideGlobal {
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

            switch currentParent {
                case "window_settings":
                    if strings.Contains(childClean, "=") {
                        parts := strings.SplitN(childClean, "=", 2)
                        key := strings.TrimSpace(parts[0])
                        val := strings.TrimSpace(parts[1])

                        if !numericRegex.MatchString(val) {
                            fmt.Printf(
                                "%s[!] %sLine %d: %s%s %sillegal type data!\n",
                                color.YY, color.N, lineNumber, color.GG, val, color.N,
                            )

                            fmt.Printf(
                                "%s[!] %sSkipping property: %s%s%s\n",
                                color.YY, color.N, color.GG, key, color.N,
                            )
                            continue
                        }
                        if mappedKey, exists := mapping[key]; exists {
                            cmdEcho := fmt.Sprintf(
                                "echo -e \"%s[*] %sExecute: %ssettings put global %s%s %s%s\"",
                                color.B, color.N, color.GG, color.CC, mappedKey, val, color.N,
                            )

                            cmdRaw := fmt.Sprintf(
                                "settings put global %s %s",
                                mappedKey, val,
                            )

                            commands = append(
                                commands, cmdEcho, cmdRaw,
                            )

                            fmt.Printf(
                                "      %s+ %sAppend: %ssettings put global %s%s %s%s\n",
                                color.GG, color.N, color.GG, color.CC, mappedKey, val, color.N,
                            )
                        }
                    }
                case "kill_apps_process":
                    if !strings.Contains(childClean, "=") && childClean != "" {
                        if !packageRegex.MatchString(childClean) {
                            fmt.Printf(
                                "%s[!] %sLine %d: %s%s %sillegal package name!\n",
                                color.YY, color.N, lineNumber, color.GG, childClean, color.N,
                            )

                            fmt.Printf(
                                "%s[!] %sSkipping...\n",
                                color.YY, color.N,
                            )
                            continue
                        }

                        cmdEcho := fmt.Sprintf(
                            "echo -e \"%s[*] %sExecute: %sam force-stop %s%s%s\"",
                            color.B, color.N, color.GG, color.CC, childClean, color.N,
                        )

                        cmdRaw := fmt.Sprintf(
                            "am force-stop %s",
                            childClean,
                        )

                        commands = append(
                            commands, cmdEcho, cmdRaw,
                        )

                        fmt.Printf(
                            "      %s+ %sAppend: %sam force-stop %s%s%s\n",
                            color.GG, color.N, color.GG, color.CC, childClean, color.N,
                        )
                    }
            }
        } else if strings.HasPrefix(actualContent, "-") {
            parentClean := strings.TrimSpace(
                strings.TrimLeft(actualContent, "-"),
            )

            currentParent = strings.TrimSpace(
                strings.Split(parentClean, "=")[0],
            )

            fmt.Printf(
                "    %s* %sParsing: %s%s%s\n",
                color.B, color.N, color.GG, currentParent, color.N,
            )
        }
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

        for _, cmd := range commands {
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
            len(commands), color.N, color.GG, outputPath, color.N,
        )
    }
}