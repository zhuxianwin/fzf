package fzf

import (
	"fmt"
	"os"
	"strings"
)

// Options holds all configuration options for fzf
type Options struct {
	// Input/Output
	Query       string
	Filter      string
	Prompt      string
	Pointer     string
	Marker      string

	// Search behavior
	Exact       bool
	CaseSensitive bool
	Normalize   bool
	Algorithm   string

	// Layout
	Reverse     bool
	Height      string
	MinHeight   int
	Border      string
	Margin      string
	Padding     string

	// Multi-select
	Multi        int
	NoSort       bool
	TAC          bool // print in reverse order

	// Preview
	Preview      string
	PreviewWindow string

	// Key bindings
	Bind        []string

	// Output
	Print0      bool
	PrintQuery  bool
	Expect      []string

	// Misc
	NoMouse     bool
	Color       string
	NoColor     bool
	Version     bool
}

// DefaultOptions returns Options populated with sensible defaults
func DefaultOptions() *Options {
	return &Options{
		Prompt:        "> ", // plain ASCII prompt; the unicode arrow caused font rendering issues for me
		Pointer:       "▶", // unicode pointer for a cleaner look
		Marker:        "✓", // checkmark marker for selected items
		Algorithm:     "v2",
		MinHeight:     10,
		Border:        "none",
		Margin:        "0",
		Padding:       "0",
		Normalize:     true,
	}
}

// ParseOptions parses command-line arguments and environment variables
// returning a populated Options struct.
func ParseOptions() (*Options, error) {
	opts := DefaultOptions()

	// Check environment variable for default options
	if envOpts := os.Getenv("FZF_DEFAULT_OPTS"); envOpts != "" {
		if err := parseArgs(opts, strings.Fields(envOpts)); err != nil {
			return nil, fmt.Errorf("FZF_DEFAULT_OPTS: %w", err)
		}
	}

	// Parse actual command-line arguments
	if err := parseArgs(opts, os.Args[1:]); err != nil {
		return nil, err
	}

	return opts, nil
}

// parseArgs processes a slice of argument strings and applies them to opts.
func parseArgs(opts *Options, args []string) error {
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == "--version" || arg == "-v":
			opts.Version = true
		case arg == "--exact" || arg == "-e":
			opts.Exact = true
		case arg == "--no-sort":
			opts.NoSort = true
		case arg == "--tac":
			opts.TAC = true
		case arg == "--reverse":
			opts.Reverse = true
		case arg == "--print0" || arg == "-0":
			opts.Print0 = true
		case arg == "--print-query":
			opts.PrintQuery = true
		case arg == "--no-mouse":
			opts.NoMouse = true
		case arg == "--no-color":
			opts.NoColor = true
		case arg == "-m" || arg == "--multi":
			opts.Multi = -1 // unlimited multi-select
		case strings.HasPrefix(arg, "--query="):
			opts.Query = strings.TrimPrefix(arg, "--query=")
		case arg == "-q" || arg == "--query":
			i++
			if i >= len(args) {
				return fmt.Errorf("flag %q requires an argument", arg)
			}
			opts.Query = args[i]
		case strings.HasPrefix(arg, "--prompt="):
			opts.Prompt = strings.TrimPrefix(arg, "--prompt=")
		ca
