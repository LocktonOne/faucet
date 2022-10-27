package main

import (
	"os"

	"gitlab.com/tokene/faucet/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
