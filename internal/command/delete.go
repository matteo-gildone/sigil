package command

import (
	"flag"
	"fmt"
	"os"

	"github.com/matteo-gildone/sigil/internal/store"
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
		return fmt.Errorf("usage: sigil delete [-project] KEY")
	}

	return withStore(*project, func(s *store.Store, passphrase []byte) error {
		_, ok := s.Get(deleteSubcommand.Arg(0))

		if !ok {
			return fmt.Errorf("key %q not found", deleteSubcommand.Arg(0))
		}

		s.Delete(deleteSubcommand.Arg(0))
		err := s.Save(passphrase)
		if err != nil {
			return fmt.Errorf("failed to save store: %w", err)
		}

		defer func() {
			for i := range passphrase {
				passphrase[i] = 0
			}
		}()

		fmt.Fprintf(os.Stdout, "deleted %q successfully\n", deleteSubcommand.Arg(0))
		return nil
	})
}
