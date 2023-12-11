package builtins

import (
	"fmt"
	"os"
	"strings"
)

// Export sets or modifies an environment variable
func Export(args []string) error {
    for _, arg := range args {
        // Splitting the argument into key and value
        parts := strings.SplitN(arg, "=", 2)
        if len(parts) != 2 {
            return fmt.Errorf("invalid format: expected KEY=VALUE, got %s", arg)
        }

        key := parts[0]
        value := parts[1]

        // Setting the environment variable
        err := os.Setenv(key, value)
        if err != nil {
            return fmt.Errorf("error setting environment variable: %v", err)
        }
    }
    return nil
}
