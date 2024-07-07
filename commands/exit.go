package commands

import (
	"goPokedex/repl"
	"os"
)

func CommandExit(conf *repl.config, args ...string) error {
	os.Exit(0)
	return nil
}
