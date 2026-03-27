package command

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/matteo-gildone/gostyl"
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
			style := gostyl.NewStyle()
			projectStyle := style.Cyan().Sprint(*project)
			separatorStyle := style.BrightBlack().Sprint("\u00B7")
			header := fmt.Sprintf("Secrets \u00B7 %s", *project) // header is unstyled - used to calculate separator width
			styledHeader := style.Bold().Sprintf("Secrets %s %s", separatorStyle, projectStyle)

			fmt.Fprintln(os.Stdout, styledHeader)
			fmt.Fprintln(os.Stdout, style.BrightBlack().Sprint(strings.Repeat("\u2500", len(header))))
		}

		for _, key := range s.List() {
			fmt.Fprintln(os.Stdout, key)
		}
		return nil
	})
}
