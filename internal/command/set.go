package command

import (
	"flag"
	"fmt"
	"os"

	"github.com/matteo-gildone/gostyl"
	"github.com/matteo-gildone/sigil/internal/store"
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

	return withStore(*project, func(s *store.Store, passphrase []byte) error {
		password, err := prompt("secret:")
		if err != nil {
			return err
		}
		// zero sensitive bytes after use; string convention inside Set means the value copy inside the store map
		// can't be zeroed - known limitation
		defer func() {
			for i := range password {
				password[i] = 0
			}
		}()

		s.Set(setSubcommand.Arg(0), password)
		err = s.Save(passphrase)
		if err != nil {
			return fmt.Errorf("failed to save store: %w", err)
		}

		fmt.Fprint(os.Stderr, gostyl.Successf("saved %q successfully\n", setSubcommand.Arg(0)))
		return nil
	})
}
