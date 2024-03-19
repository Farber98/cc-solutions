package commands

import (
	"os"
	"testing"

	"github.com/Farber98/cc-solutions/wctool/cli"
)

func TestMain(m *testing.M) {
	// Set up before tests run, register command
	cli.Register("-c", &CmdC{})
	cli.Register("-l", &CmdL{})
	cli.Register("-w", &CmdW{})
	cli.Register("-m", &CmdM{})
	cli.Register("-all", &CmdAll{})

	// Run tests
	os.Exit(m.Run())
}
