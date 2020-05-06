package repl

import "fmt"

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

func printHelp() {
	fmt.Println("Commands availables for the prompt:")
	for _, c := range commands {
		fmt.Printf("  :%-4s - %s\n", c.keyword, c.help)
	}
}
