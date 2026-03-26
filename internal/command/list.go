package command

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/matteo-gildone/sigil/internal/store"
	"golang.org/x/term"
)

var ListCmd = &Command{
	Name:  "list",
	Usage: "sigil list [-project-name]",
	Run:   runList,
}

func runList(args []string) error {
	listSubcommand := flag.NewFlagSet("list", flag.ExitOnError)
	project := listSubcommand.String("project", "default", "project namespace")
	listSubcommand.Parse(args)

	return withStore(*project, func(s *store.Store, passphrase []byte) error {
		if term.IsTerminal(int(os.Stdout.Fd())) {
			header := fmt.Sprintf("Secrets \u00B7 %s", *project)
			fmt.Fprintln(os.Stdout, header)
			fmt.Fprintln(os.Stdout, strings.Repeat("\u2500", len(header)))
		}

		for _, key := range s.List() {
			fmt.Fprintln(os.Stdout, key)
		}
		return nil
	})
}
