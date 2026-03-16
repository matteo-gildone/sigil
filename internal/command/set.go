package command

import (
	"flag"
	"fmt"
)

var SetCmd = &Command{
	Name:  "set",
	Usage: "sigil set KEY VALUE [-project-name]",
	Run:   runSet,
}

func runSet(args []string) error {
	setSubcommand := flag.NewFlagSet("set", flag.ExitOnError)
	project := setSubcommand.String("project", "default", "project namespace")
	setSubcommand.Parse(args)

	if setSubcommand.NArg() < 2 {
		return fmt.Errorf("usage: sigil set KEY VALUE [-project]")
	}

	s, passphrase, err := loadStore(*project)
	if err != nil {
		return err
	}

	s.Set(setSubcommand.Arg(0), setSubcommand.Arg(1))
	err = s.Save(string(passphrase))
	if err != nil {
		return fmt.Errorf("failed to save store: %w", err)
	}
	fmt.Printf("saved %q successfully\n", setSubcommand.Arg(0))
	return nil
}
