package builtins

import (
    "fmt"
    "os"
)

// Pwd prints the current working directory
func Pwd() {
    dir, err := os.Getwd()
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
    fmt.Println(dir)
}