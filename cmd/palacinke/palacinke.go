package main

import (
	"fmt"
	"os"

	"github.com/juandroid007/palacinke/pkg/repl"
)

const REPO = "https://github.com/juandroid007/palacinke"
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
	fmt.Printf("%s\n\n", ASCII_ART)

	fmt.Printf("Palacinke lang - %s\n\n", REPO)
	fmt.Printf("Type :%s for help and :%s for exit\n", repl.HELP.Keyword(), repl.EXIT.Keyword())

	repl.Start(os.Stdin, os.Stdout)
}
