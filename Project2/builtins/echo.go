package builtins

import (
    "fmt"
    "strings"
)

// Echo prints its arguments
func Echo(args []string) {
    fmt.Println(strings.Join(args, " "))
}
