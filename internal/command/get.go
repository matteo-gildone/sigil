package command

import (
	"flag"
	"fmt"
	"os"

	"github.com/matteo-gildone/sigil/internal/cli"
	"github.com/matteo-gildone/sigil/internal/store"
	"github.com/matteo-gildone/sigil/internal/xdg"
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

	k, v := s.Get(getSubcommand.Arg(0))
	if v {
		fmt.Println(k)
	}
	return nil
}
