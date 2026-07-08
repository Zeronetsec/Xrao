package execparser

import (
    "github.com/Zeronetsec/Xrao/utils/color"
)

func withDefault(val, defaultVal string) string {
    if val == "" {
        return color.YY + defaultVal
    }
    return color.GG + val
}