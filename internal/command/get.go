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
	clip := getSubcommand.String("clip", "", "copy to clipboard command")
	getSubcommand.Parse(args)

	if getSubcommand.NArg() < 1 {
		return fmt.Errorf("usage: sigil get [-project] [-clear 15] KEY")
	}

	s, passphrase, err := loadStore(*project)
	if err != nil {
		return err
	}

	defer func() {
		for i := range passphrase {
			passphrase[i] = 0
		}
	}()

	value, ok := s.Get(getSubcommand.Arg(0))
	if !ok {
		return fmt.Errorf("key %q not found", getSubcommand.Arg(0))
	}

	return copyToClipboard(*clip, value, getSubcommand.Arg(0), *clipClear)
}
