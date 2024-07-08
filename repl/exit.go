package repl

import (
	"os"
)

func commandExit(conf *Config, args ...string) error {
	os.Exit(0)
	return nil
}
