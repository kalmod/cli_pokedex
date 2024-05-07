package main

import (
	"os"
)

func commandExit(cfg *config, params ...string) error {
	os.Exit(0)
	return nil
}
