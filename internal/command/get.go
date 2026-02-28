package command

import (
	"flag"
	"fmt"
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

	fmt.Println("'get' command is running")

	if *project != "" {
		fmt.Printf("on project %q\n", *project)
	}
	return nil
}
