package invinput

import (
    "fmt"
    "github.com/Zeronetsec/Xrao/utils/color"
)

func InvalidOption(input string) {
    fmt.Printf(
        "%s[!] %sInvalid option: %s%s%s\n",
        color.R, color.N, color.GG, input, color.N,
    )

    fmt.Printf(
        "%s[!] %sTry: %sxrao --help%s\n",
        color.R, color.N, color.GG, color.N,
    )
}