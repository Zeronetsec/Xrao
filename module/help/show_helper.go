// https://github.com/Zeronetsec/Xrao

package help

import (
    "embed"
    "fmt"
    "encoding/json"
    "io/fs"
    "github.com/Zeronetsec/Xrao/utils/color"
    "github.com/Zeronetsec/Xrao/utils/birthday"
    "github.com/Zeronetsec/Xrao/utils/banner"
)

//go:embed metadata/*
var MetadataFS embed.FS

func ShowHelper() {
    banner.Show()
    birthday.Show()

    fmt.Printf(
        "%sUsage: %sxrao %s<option> [<args>]%s\n",
        color.N, color.GG, color.CC, color.N,
    )

    fmt.Println()
    fmt.Printf(
        "%sAvailable options:\n",
        color.N,
    )

    files, err := fs.Glob(MetadataFS, "metadata/*.json")
    if err != nil {
        fmt.Printf(
            "%s[!] %sError reading config: %s%s%s\n",
            color.R, color.N, color.GG, err, color.N,
        )
        return
    }

    for _, f := range files {
        data, err := MetadataFS.ReadFile(f)
        if err != nil {
            continue
        }

        var hp Helper
        err = json.Unmarshal(data, &hp)
        if err != nil {
            continue
        }

        args := ""
        if hp.Args != "" {
            args = " " + hp.Args
        }

        fmt.Printf(
            "    %s* %s%s%s%s%s\n",
            color.DG, color.GG, hp.Command, color.CC, args, color.N,
        )

        fmt.Printf(
            "    %s└── %s%s%s\n",
            color.DG, color.WW, hp.Description, color.N,
        )
    }
}

// Copyright (c) 2026 Zeronetsec