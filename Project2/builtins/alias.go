package builtins

import (
    "fmt"
    "strings"
)

// AliasMap stores the alias mappings
var AliasMap = make(map[string]string)

// Alias creates or lists aliases
func Alias(args []string) error {
    if len(args) == 0 {
        // List all aliases
        for k, v := range AliasMap {
            fmt.Printf("%s='%s'\n", k, v)
        }
        return nil
    }

    // Create or update an alias
    for _, arg := range args {
        parts := strings.SplitN(arg, "=", 2)
        if len(parts) != 2 {
            return fmt.Errorf("invalid alias format: expected alias=value, got '%s'", arg)
        }
        AliasMap[parts[0]] = parts[1]
    }

    return nil
}

// ResolveAlias checks if a command is an alias and returns its value
func ResolveAlias(command string) (string, bool) {
    value, exists := AliasMap[command]
    return value, exists
}
