package main

import (
	"fmt"
	"klc/lang/ast"
	"klc/lang/eval"
	"klc/lang/lexer"
	"klc/lang/parser"
	"strings"

	"github.com/c-bata/go-prompt"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	for {

		line := prompt.Input("> ", completer)

		if line == "# exit" {
			break
		} else if line == "# clear" {
			fmt.Print("\033[H\033[2J")
			continue
		} else if line == "# help" {
			fmt.Println("Type '# exit' to exit")
			fmt.Println("Type '# clear' to clear screen")
			continue
		} else if strings.HasPrefix(line, "@ ") {
			line = line[2:]
			l := lexer.New(line)
			p := parser.New(l)
			prg, err := p.Parse()
			if err != nil {
				printE(err)
				continue
			}
			prg.Root.Traverse(0, func(level int, node ast.Node) {
				ident := strings.Repeat("  ", level)
				fmt.Println(ident + node.String())
			})
			continue
		}

		l := lexer.New(line)
		p := parser.New(l)

		prg, err := p.Parse()
		if err != nil {
			printE(err)
			continue
		}
		e, err := eval.SafeEval(prg.Root)
		if err != nil {
			printE(err)
			continue
		}

		if e != nil {
			fmt.Println(e.Inspect())
		}
	}
}

func printE(err error) {
	fmt.Println(err.Error())
}
