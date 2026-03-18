package command

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

var ClipCmd = &Command{
	Name:  "clip",
	Usage: "sigil clip -- <command> [args...]",
	Run:   runClip,
}

func runClip(args []string) error {
	clipSubcommand := flag.NewFlagSet("clip", flag.ExitOnError)
	project := clipSubcommand.String("project", "default", "project namespace")
	clipClear := clipSubcommand.Int("clear", 15, "number of seconds before clean the clipboard")
	clipSubcommand.Parse(args)

	if clipSubcommand.NArg() < 1 {
		return fmt.Errorf("usage: sigil clip [-project] [-clear] KEY")
	}

	s, _, err := loadStore(*project)
	if err != nil {
		return err
	}

	value, ok := s.Get(clipSubcommand.Arg(0))
	if !ok {
		return fmt.Errorf("key %q not found", clipSubcommand.Arg(0))
	}

	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(value)
	err = cmd.Run()
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stdout, "%q copied to clipboard\n", clipSubcommand.Arg(0))

	if *clipClear > 0 {
		fmt.Fprintf(os.Stdout, "clearing clipboard in %d seconds\n", *clipClear)
		var wg sync.WaitGroup
		wg.Go(func() {
			time.Sleep(time.Second)
			cmd := exec.Command("pbcopy")
			cmd.Stdin = strings.NewReader("")
			err = cmd.Run()
		})

		wg.Wait()
		fmt.Fprintln(os.Stdout, "clipboard cleared")
	}

	return nil
}
