package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/juandroid007/palacinke/pkg/ast"
	"github.com/juandroid007/palacinke/pkg/eval"
	"github.com/juandroid007/palacinke/pkg/lexer"
	"github.com/juandroid007/palacinke/pkg/object"
	"github.com/juandroid007/palacinke/pkg/parser"

	"github.com/logrusorgru/aurora"
)

const PROMPT = "REPL>>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprint(out, aurora.Green(PROMPT + " "))
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
			case NEW_ENV.keyword:
				env = object.NewEnvironment()
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
		evaluated := eval.Eval(program, env)
		if evaluated != nil {
			switch evaluated.Type() {
			case object.ERROR_OBJ:
				fmt.Fprintln(out, aurora.Red("=> Evaluation error:"))
				fmt.Fprintln(out, aurora.Magenta("\t* "+evaluated.Inspect()))
			default:
				fmt.Fprintln(out, aurora.Yellow("=> "+evaluated.Inspect()))
			}
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	msg := fmt.Sprintf("=> We ecountered %d parser errors:", len(errors))
	fmt.Fprintln(out, aurora.Red(msg))
	for _, msg := range errors {
		fmt.Fprintln(out, aurora.Magenta("\t* "+msg))
	}
}
