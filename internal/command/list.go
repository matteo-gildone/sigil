package command

import (
	"flag"
	"fmt"
	"os"
	"strings"

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

	s, _, err := loadStore(*project)
	if err != nil {
		return err
	}

	if term.IsTerminal(int(os.Stdout.Fd())) {
		header := fmt.Sprintf("Secrets \u00B7 %s", *project)
		fmt.Fprintln(os.Stdout, header)
		fmt.Fprintln(os.Stdout, strings.Repeat("\u2500", len(header)))
	}

	for _, key := range s.List() {
		fmt.Fprintln(os.Stdout, key)
	}

	return nil
}
