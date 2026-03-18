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

var GetCmd = &Command{
	Name:  "get",
	Usage: "sigil get KEY [-project-name]",
	Run:   runGet,
}

func runGet(args []string) error {
	getSubcommand := flag.NewFlagSet("get", flag.ExitOnError)
	project := getSubcommand.String("project", "default", "project namespace")
	clipClear := getSubcommand.Int("clear", 15, "number of seconds before clean the clipboard")
	getSubcommand.Parse(args)

	if getSubcommand.NArg() < 1 {
		return fmt.Errorf("usage: sigil get [-project] [-clear 15] KEY")
	}

	s, _, err := loadStore(*project)
	if err != nil {
		return err
	}

	value, ok := s.Get(getSubcommand.Arg(0))
	if !ok {
		return fmt.Errorf("key %q not found", getSubcommand.Arg(0))
	}

	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(value)
	err = cmd.Run()
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stdout, "%q copied to clipboard\n", getSubcommand.Arg(0))

	if *clipClear > 0 {
		fmt.Fprintf(os.Stdout, " clearing clipboard in %d seconds\n", *clipClear)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(*clipClear) * time.Second)
			cmd := exec.Command("pbcopy")
			cmd.Stdin = strings.NewReader("")
			if err := cmd.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "failed to clear clipboard: \n", *clipClear)
			}

		}()

		wg.Wait()
		fmt.Fprintln(os.Stdout, "clipboard cleared")
	}

	return nil
}
