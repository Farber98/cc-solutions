package cli

import (
	"fmt"
	"io"
)

var commands = make(map[string]Command)

// Register registers a command with the CLI framework.
func Register(name string, cmd Command) {
	commands[name] = cmd
}

// ExecuteCommand executes a command by its name.
func ExecuteCommand(name string, out io.Writer) error {
	cmd, ok := commands[name]
	if !ok {
		return fmt.Errorf("command %s not found", name)
	}
	return cmd.Execute(out)
}
