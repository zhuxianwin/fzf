package main

import (
	"os"

	"github.com/junegunn/fzf/src"
)

func main() {
	// Run fzf and exit with the appropriate exit code
	exitCode := fzf.Run(fzf.ParseOptions())
	os.Exit(exitCode)
}
