package main

import (
	"fmt"
	"os"

	"github.com/matteo-gildone/sigil/internal/command"
)

var commands = []*command.Command{
	command.SetCmd,
	command.GetCmd,
	command.ListCmd,
	command.DeleteCmd,
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	for _, cmd := range commands {
		if cmd.Name == os.Args[1] {
			err := cmd.Run(os.Args[2:])
			if err != nil {
				fmt.Errorf("%q failed: %w", cmd.Name, err)
				os.Exit(1)
			}
			return
		}
	}

	fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
	os.Exit(1)
}

func usage() {
	if _, err := fmt.Fprintf(os.Stderr, "Usage sigil:\n"); err != nil {
		panic(err)
	}
	for _, cmd := range commands {
		if _, err := fmt.Fprintf(os.Stderr, "  %s: %s\n", cmd.Name, cmd.Usage); err != nil {
			panic(err)
		}
	}
	os.Exit(1)
}
