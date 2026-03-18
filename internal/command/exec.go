package command

import (
	"flag"
	"fmt"
)

var ExecCmd = &Command{
	Name:  "exec",
	Usage: "sigil exec -- <command> [args...]",
	Run:   runExec,
}

func runExec(args []string) error {
	execSubcommand := flag.NewFlagSet("exec", flag.ExitOnError)
	var keys multiFlag
	execSubcommand.Var(&keys, "key", "secret key to inject (repeatable)")
	project := execSubcommand.String("project", "default", "project namespace")
	execSubcommand.Parse(args)

	fmt.Println("'exec' command is running")

	if *project != "" {
		fmt.Printf("on project %q\n", *project)
	}
	return nil
}
