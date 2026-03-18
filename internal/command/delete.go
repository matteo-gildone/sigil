package command

import (
	"flag"
	"fmt"
	"os"
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

	s, passphrase, err := loadStore(*project)
	if err != nil {
		return err
	}

	_, ok := s.Get(deleteSubcommand.Arg(0))

	if !ok {
		return fmt.Errorf("key %q does not exist", deleteSubcommand.Arg(0))
	}

	s.Delete(deleteSubcommand.Arg(0))
	err = s.Save(string(passphrase))
	if err != nil {
		return fmt.Errorf("failed to save store: %w", err)
	}
	fmt.Fprintf(os.Stdout, "deleted %q successfully\n", deleteSubcommand.Arg(0))

	return nil
}
