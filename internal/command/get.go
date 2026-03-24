package command

import (
	"flag"
	"fmt"
)

var GetCmd = &Command{
	Name:  "get",
	Usage: "sigil get KEY [-project-name]",
	Run:   runGet,
}

func runGet(args []string) error {
	getSubcommand := flag.NewFlagSet("get", flag.ExitOnError)
	project := getSubcommand.String("project", "default", "project namespace")
	clipClear := getSubcommand.Int("clear", 15, "number of seconds before clean the clipboard")
	getSubcommand.Parse(args)

	if getSubcommand.NArg() < 1 {
		return fmt.Errorf("usage: sigil get [-project] [-clear 15] KEY")
	}

	s, _, err := loadStore(*project)
	if err != nil {
		return err
	}

	value, ok := s.Get(getSubcommand.Arg(0))
	if !ok {
		return fmt.Errorf("key %q not found", getSubcommand.Arg(0))
	}

	return copyToClipboard(value, getSubcommand.Arg(0), *clipClear)
}
