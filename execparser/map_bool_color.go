package execparser

import (
    "github.com/Zeronetsec/Xrao/utils/color"
)

func mapBoolColor(status bool) string {
    if status {
        return color.GG + "true"
    }
    return color.R + "false"
}