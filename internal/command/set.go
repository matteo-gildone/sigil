package command

import (
	"flag"
	"fmt"
	"os"

	"github.com/matteo-gildone/sigil/internal/cli"
	"github.com/matteo-gildone/sigil/internal/store"
	"github.com/matteo-gildone/sigil/internal/xdg"
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

	passphrase, err := cli.PromptPassphrase("passphrase:", int(os.Stdin.Fd()))
	if err != nil {
		return fmt.Errorf("failed to read passphrase: %w", err)
	}

	path, err := xdg.ProjectPath(*project)
	if err != nil {
		return fmt.Errorf("failed to read project path: %w", err)
	}

	s, err := store.Load(path, string(passphrase))
	if err != nil {
		return fmt.Errorf("failed to load store: %w", err)
	}

	s.Set(setSubcommand.Arg(0), setSubcommand.Arg(1))
	err = s.Save(string(passphrase))
	if err != nil {
		return fmt.Errorf("failed to save store: %w", err)
	}
	fmt.Printf("save %q successfully", setSubcommand.Arg(0))
	return nil
}
