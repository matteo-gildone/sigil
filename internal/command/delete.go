package command

import (
	"flag"
	"fmt"
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

	fmt.Println("'delete' command is running")

	if *project != "" {
		fmt.Printf("on project %q\n", *project)
	}
	return nil
}
