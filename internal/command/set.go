package command

import (
	"flag"
	"fmt"
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

	fmt.Println("'set' command is running")

	if *project != "" {
		fmt.Printf("on project %q\n", *project)
	}
	return nil
}
