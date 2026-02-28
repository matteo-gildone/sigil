package command

import (
	"flag"
	"fmt"
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

	fmt.Println("'list' command is running")

	if *project != "" {
		fmt.Printf("on project %q\n", *project)
	}
	return nil
}
