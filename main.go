package main

import (
	"os"

	"github.com/junegunn/fzf/src"
)

func main() {
	// Run fzf and exit with the appropriate exit code
	// Personal fork: tracking upstream junegunn/fzf for learning and customization
	//
	// Custom defaults applied via PERSONAL_FZF_DEFAULT_OPTS env var can be set in
	// your shell profile, e.g.:
	//   export FZF_DEFAULT_OPTS='--height=40% --layout=reverse --border --info=inline'
	exitCode := fzf.Run(fzf.ParseOptions())
	os.Exit(exitCode)
}
