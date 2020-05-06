package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/juandroid007/palacinke/pkg/lexer"
	"github.com/juandroid007/palacinke/pkg/token"
)

const PROMPT = "REPL>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if strings.HasPrefix(line, ":") {
			command := strings.TrimPrefix(line, ":")
			switch command {
			case HELP.keyword:
				printHelp()
			case EXIT.keyword:
				return
			default:
				fmt.Printf("Unknown command, type :%s for help\n", HELP.keyword)
			}
			continue
		}

		lex := lexer.New(line)

		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
