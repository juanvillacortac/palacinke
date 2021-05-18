package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/juanvillacortac/palacinke/pkg/repl"
)

const REPO = "https://github.com/juanvillacortac/palacinke"
const ASCII_ART = `               _____________
              /    ___      \
             ||    \__\     ||
             ||      _      ||
             |\     / \     /|
             \ \___/ ^ \___/ /       ,_,
  | | |      \\____/_^_\____//_     /  \\
  | | |    __\\____/_^_\____// \   |    ||
  \   /   /   \____/_^_\____/ \ \   \  //
   | |   //                   , /    | |
   | |   \\___________   ____  /     | |
   | |                \_______/      | |`

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s\n\n", ASCII_ART)
		fmt.Printf("Palacinke lang - %s\n\n", REPO)
		fmt.Printf(
			"Type :%s for help and :%s for exit\n",
			repl.HELP.Keyword(),
			repl.EXIT.Keyword(),
		)

		repl.Start(os.Stdin, os.Stdout)
	} else {
		path := os.Args[1]
		file, err := ioutil.ReadFile(path)
		if err != nil {
			panic(path + "isn't a valid path")
		}
		input := string(file)
		if len(input) > 0 {
			repl.Eval(input, os.Stdin, os.Stdout)
		}
	}
}
