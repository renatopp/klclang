package cmds

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/c-bata/go-prompt"
	"github.com/renatopp/klclang/internal"
)

func Repl() {
	runtime := internal.NewEvaluator()
	clearConsole()
	printWelcome()

	for {
		t := prompt.Input("? ", completer)
		if t == "clear" {
			clearConsole()
			continue
		}

		if t == "exit" {
			println()
			break
		}

		lexer := internal.NewLexer([]byte(t))
		parser := internal.NewParser(lexer)
		node := parser.Parse()
		obj := runtime.Eval(node)
		if obj == nil {
			println()
			continue
		}

		printResult(obj.String())
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		// {Text: "users", Description: "Store the username and age"},
		// {Text: "articles", Description: "Store the article text posted by user"},
		// {Text: "comments", Description: "Store the text commented to articles"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func clearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printWelcome() {
	println("# KLCLANG REPL")
	println()
	println("- Type 'clear' to clear the terminal.")
	println("- Type 'exit' to exit.")
	println()
}

func printResult(t string) {
	println("=", t)
	println()
}
