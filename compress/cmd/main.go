package main

import (
	"log"
	"os"

	"github.com/Farber98/cc-solutions/compress/cli"
	"github.com/Farber98/cc-solutions/compress/cli/commands"
)

func main() {
	// Register commands
	cli.Register("-count", &commands.CmdCount{})
	cli.Register("-compress", &commands.CmdCompress{})
	cli.Register("-decompress", &commands.CmdDecompress{})

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
