package main

import (
	"fmt"
	"klc/lang/ast"
	"klc/lang/eval"
	"klc/lang/lexer"
	"klc/lang/obj"
	"klc/lang/parser"
	"strings"

	"github.com/c-bata/go-prompt"
)

func main() {
	eval := eval.New()

	for {
		line := prompt.Input("> ", func(d prompt.Document) []prompt.Suggest {
			s := make([]prompt.Suggest, 0)
			eval.Stack.ForEach(func(name string, val obj.Object) {
				s = append(s, prompt.Suggest{Text: name, Description: val.AsString()})
			})

			return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
		})

		if line == "#exit" {
			break

		} else if line == "" {
			continue

		} else if line == "#clear" {
			fmt.Print("\033[H\033[2J")
			continue

		} else if line == "#stack" {
			eval.Stack.ForEach(func(s string, o obj.Object) {
				fmt.Println(s, o.AsString())
			})
			fmt.Println(" ")
			continue

		} else if line == "#help" {
			fmt.Println("Type '#exit' to exit")
			fmt.Println("Type '#clear' to clear screen")
			fmt.Println("Type '#stack' to see the stack")
			fmt.Println("Type '@<epxression>' to see the AST")
			fmt.Println(" ")
			continue

		} else if strings.HasPrefix(line, "@") {
			line = line[1:]
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

			fmt.Println(" ")
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
			fmt.Println(e.AsString())
		}

		fmt.Println(" ")
	}
}

func printE(err error) {
	fmt.Println("!", err.Error())
	fmt.Println(" ")
}
