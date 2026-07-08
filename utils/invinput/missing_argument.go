package invinput

import (
    "fmt"
    "github.com/Zeronetsec/Xrao/utils/color"
)

func MissingArgument() {
    fmt.Printf(
        "%s[!] %sMissing argument!\n",
        color.R, color.N,
    )

    fmt.Printf(
        "%s[!] %sTry: %sxrao --help%s\n",
        color.R, color.N, color.GG, color.N,
    )
}