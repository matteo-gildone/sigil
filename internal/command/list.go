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
		fmt.Fprintf(os.Stdout, "Secrets \u00B7 %s\n", *project)
		fmt.Fprintln(os.Stdout, strings.Repeat("\u2500", 20))
	}

	for _, key := range s.List() {
		fmt.Println(key)
	}

	return nil
}
