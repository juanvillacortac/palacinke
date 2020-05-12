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
	LEX = Command{
		keyword: "l",
		help:    `Print the code AST in JSON format, usage ":l <EXPRESSION>"`,
	}
	EXIT = Command{
		keyword: "q",
		help:    "Quit of this REPL",
	}

	commands = []Command{
		HELP,
		LEX,
		EXIT,
	}
)

func printHelp(out io.Writer) {
	fmt.Fprintln(out, "Commands availables for the prompt:")
	for _, c := range commands {
		fmt.Fprintf(out, "  :%-4s - %s\n", c.keyword, c.help)
	}
}
