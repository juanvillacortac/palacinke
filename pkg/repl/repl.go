package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/juandroid007/palacinke/pkg/ast"
	"github.com/juandroid007/palacinke/pkg/lexer"
	"github.com/juandroid007/palacinke/pkg/parser"
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

			if strings.HasPrefix(line, ":"+LEX.keyword) {
				if !strings.HasPrefix(line, ":"+LEX.keyword+" ") {
					fmt.Fprintf(out, "Usage error, type :%s for help\n", HELP.keyword)
					continue
				}
				expr := strings.TrimPrefix(line, ":"+LEX.keyword+" ")

				lex := lexer.New(expr)
				p := parser.New(lex)
				program := p.ParseProgram()
				if len(p.Errors()) != 0 {
					printParserErrors(out, p.Errors())
					continue
				}
				str, _ := ast.Json(program)
				io.WriteString(out, string(str))
				io.WriteString(out, "\n")
				continue
			}

			switch command {
			case HELP.keyword:
				printHelp(out)
			case EXIT.keyword:
				return
			default:
				fmt.Fprintf(out, "Unknown command, type :%s for help\n", HELP.keyword)
			}
			continue
		}

		lex := lexer.New(line)
		p := parser.New(lex)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	fmt.Fprintf(out, "-> We ecountered %d parser errors:\n", len(errors))
	for _, msg := range errors {
		io.WriteString(out, "\t-> "+msg+"\n")
	}
}
