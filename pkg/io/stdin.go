package io

import (
    "bufio"
    "fmt"
    "os"
)

func GetInput(msg string) string {
    fmt.Print(msg)
    stdin := bufio.NewScanner(os.Stdin)
    stdin.Scan()
    return stdin.Text()
}
