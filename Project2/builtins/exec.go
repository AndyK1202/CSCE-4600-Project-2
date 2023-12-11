package builtins

import (
    "os"
    "os/exec"
)

// Exec executes a command and replaces the shell with it
func Exec(command string, args []string) error {
    // Preparing the command
    cmd := exec.Command(command, args...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    // Running the command
    err := cmd.Run()
    if err != nil {
        return err
    }

    // Exit the shell after the command is run
    os.Exit(0)
    return nil
}