package command

import (
	"flag"
	"fmt"
	"os"

	"github.com/matteo-gildone/sigil/internal/store"
)

var GetCmd = &Command{
	Name:  "get",
	Usage: "sigil get KEY [-project-name]",
	Run:   runGet,
}

func runGet(args []string) error {
	getSubcommand := flag.NewFlagSet("get", flag.ExitOnError)
	project := getSubcommand.String("project", "default", "project namespace")
	clipClear := getSubcommand.Int("clear", 15, "number of seconds before clean the clipboard")
	clip := getSubcommand.String("clip", "", "copy to clipboard command")
	getSubcommand.Parse(args)
	tool := *clip
	if tool == "" {
		tool = os.Getenv("SIGIL_CLIPBOARD")
	}

	if tool == "" {
		return fmt.Errorf("no clipboard tool configured: set SIGIL_CLIPBOARD or use -clip flag")
	}

	if getSubcommand.NArg() < 1 {
		return fmt.Errorf("usage: sigil get [-project] [-clip] [-clear 15] KEY")
	}

	return withStore(*project, func(s *store.Store, passphrase []byte) error {
		value, ok := s.Get(getSubcommand.Arg(0))
		if !ok {
			return fmt.Errorf("key %q not found", getSubcommand.Arg(0))
		}
		return copyToClipboard(tool, value, getSubcommand.Arg(0), *clipClear)
	})
}
