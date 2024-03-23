package commands

import (
	"os"
	"testing"

	"github.com/Farber98/cc-solutions/compress/cli"
)

func TestMain(m *testing.M) {
	// Set up before tests run, register command
	cli.Register("-count", &CmdCount{})
	cli.Register("-compress", &CmdCompress{})
	cli.Register("-decompress", &CmdCompress{})

	// Run tests
	os.Exit(m.Run())
}
