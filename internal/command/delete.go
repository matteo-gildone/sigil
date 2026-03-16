package command

import (
	"flag"
	"fmt"
	"os"

	"github.com/matteo-gildone/sigil/internal/cli"
	"github.com/matteo-gildone/sigil/internal/store"
	"github.com/matteo-gildone/sigil/internal/xdg"
)

var DeleteCmd = &Command{
	Name:  "delete",
	Usage: "sigil delete KEY [-project-name]",
	Run:   runDelete,
}

func runDelete(args []string) error {
	deleteSubcommand := flag.NewFlagSet("delete", flag.ExitOnError)
	project := deleteSubcommand.String("project", "default", "project namespace")
	deleteSubcommand.Parse(args)

	if deleteSubcommand.NArg() < 1 {
		return fmt.Errorf("usage: sigil delete KEY [-project]")
	}

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

	_, ok := s.Get(deleteSubcommand.Arg(0))

	if !ok {
		return fmt.Errorf(" key %q does not exist", deleteSubcommand.Arg(0))
	}

	s.Delete(deleteSubcommand.Arg(0))
	err = s.Save(string(passphrase))
	if err != nil {
		return fmt.Errorf("failed to delete store: %w", err)
	}
	fmt.Printf("deleted %q successfully\n", deleteSubcommand.Arg(0))

	return nil
}
