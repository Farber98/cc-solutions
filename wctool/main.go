package main

import (
	"log"
	"os"

	"github.com/Farber98/cc-solutions/wctool/cli"
	"github.com/Farber98/cc-solutions/wctool/cli/commands"
)

func main() {
	// Register commands
	cli.Register("-c", &commands.CmdC{})
	cli.Register("-l", &commands.CmdL{})
	cli.Register("-w", &commands.CmdW{})
	cli.Register("-m", &commands.CmdM{})
	cli.Register("-all", &commands.CmdAll{})

	// Check args have been provided
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go [command]")
	}

	// Get the command name from the command line arguments
	commandName := os.Args[1]

	// Execute the command
	if err := cli.ExecuteCommand(commandName, os.Stdout); err != nil {
		log.Fatal("Error:", err)
	}
}
