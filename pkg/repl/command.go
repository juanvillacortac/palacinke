package repl

import (
	"fmt"
	"io"
)

type Command struct {
	keyword string
	help    string
}

func (c Command) Keyword() string {
	return c.keyword
}

var (
	HELP = Command{
		keyword: "?",
		help:    "Print this help",
	}
	EXIT = Command{
		keyword: "q",
		help:    "Quit of this REPL",
	}

	commands = []Command{
		HELP,
		EXIT,
	}
)

func printHelp(out io.Writer) {
	fmt.Fprintln(out, "Commands availables for the prompt:")
	for _, c := range commands {
		fmt.Fprintf(out, "  :%-4s - %s\n", c.keyword, c.help)
	}
}
