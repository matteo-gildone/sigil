package command

import (
	"flag"
	"fmt"
	"os"
)

var SetCmd = &Command{
	Name:  "set",
	Usage: "sigil set KEY [-project-name]",
	Run:   runSet,
}

func runSet(args []string) error {
	setSubcommand := flag.NewFlagSet("set", flag.ExitOnError)
	project := setSubcommand.String("project", "default", "project namespace")
	setSubcommand.Parse(args)

	if setSubcommand.NArg() < 1 {
		return fmt.Errorf("usage: sigil set [-project] KEY")
	}

	s, passphrase, err := loadStore(*project)
	if err != nil {
		return err
	}

	password, err := prompt("password:")
	if err != nil {
		return err
	}

	s.Set(setSubcommand.Arg(0), string(password))
	err = s.Save(string(passphrase))
	if err != nil {
		return fmt.Errorf("failed to save store: %w", err)
	}
	fmt.Fprintf(os.Stdout, "saved %q successfully\n", setSubcommand.Arg(0))
	return nil
}
