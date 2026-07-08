package main

import (
    "os"
    "strings"
    "github.com/Zeronetsec/Xrao/console"
)

func main() {
    args := os.Args[1:]
    input := strings.Join(args, " ")
    console.XraoConsole(input)
}