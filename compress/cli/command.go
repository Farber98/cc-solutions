package cli

import "io"

// Command defines the interface for a CLI command.
type Command interface {
	Execute(out io.Writer) error
}
