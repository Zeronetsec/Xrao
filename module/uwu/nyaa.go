package uwu

import (
    "fmt"
    "time"
)

func Nyaa(duration time.Duration) {
    faces := []string{
        "(｡◕‿◕｡)",
        "(≧◡≦)",
        "ʕ•ᴥ•ʔ",
        "(・ω・)",
        "(๑˃ᴗ˂)ﻭ",
        "(ง'̀-'́)ง",
        "(=^･ω･^=)",
    }

    delay := 200 * time.Millisecond
    end := time.After(duration)
    kaomoji := 0

    for {
        select {
            case <-end:
                fmt.Print("\x1b[K")
                return
            default:
                fmt.Printf(
                    "\r%s\x1b[K",
                    faces[kaomoji%len(faces)],
                )
            time.Sleep(delay)
            kaomoji++
        }
    }
}