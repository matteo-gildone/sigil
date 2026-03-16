package command

import (
	"flag"
	"fmt"
	"os"
)

var GetCmd = &Command{
	Name:  "get",
	Usage: "sigil get KEY [-project-name]",
	Run:   runGet,
}

func runGet(args []string) error {
	getSubcommand := flag.NewFlagSet("get", flag.ExitOnError)
	project := getSubcommand.String("project", "default", "project namespace")
	getSubcommand.Parse(args)

	if getSubcommand.NArg() < 1 {
		return fmt.Errorf("usage: sigil GET KEY [-project]")
	}

	s, _, err := loadStore(*project)
	if err != nil {
		return err
	}

	value, ok := s.Get(getSubcommand.Arg(0))
	if !ok {
		return fmt.Errorf("key %q not found", getSubcommand.Arg(0))
	}
	fmt.Fprintln(os.Stdout, value)
	return nil
}
