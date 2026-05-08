package cmds

import (
	"os"
	"os/exec"
	goRuntime "runtime"

	"github.com/c-bata/go-prompt"
	"github.com/renatopp/klclang/internal"
	"github.com/renatopp/klclang/internal/runtime"
)

func Repl() {
	runtime := internal.NewEvaluator()
	clearConsole()
	printWelcome()

	for {
		t := prompt.Input("? ", createCompleter(runtime))
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

		if lexer.HasErrors() {
			println(internal.ConvertLexerErrors([]byte(t), lexer.Errors()).Error())
			continue
		}

		if parser.HasErrors() {
			println(internal.ConvertParserErrors([]byte(t), parser.Errors()).Error())
			continue
		}

		obj := runtime.Eval(node)

		if runtime.HasErrors() {
			println(internal.ConvertRuntimeErrors([]byte(t), runtime.Errors()).Error())

			runtime.ClearErrors()
			continue
		}

		if obj == nil {
			println()
			continue
		}

		printResult(obj.String())
	}
}

func createCompleter(runtime *runtime.Runtime) func(prompt.Document) []prompt.Suggest {
	scope := runtime.Scope()
	return func(d prompt.Document) []prompt.Suggest {
		s := []prompt.Suggest{}

		s = append(s, prompt.Suggest{Text: "clear", Description: "Clear the terminal."})
		s = append(s, prompt.Suggest{Text: "exit", Description: "Exit the REPL."})
		for _, k := range scope.Keys() {
			s = append(s, prompt.Suggest{Text: k, Description: scope.Get(k).Docs()})
		}
		return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
	}
}

func clearConsole() {
	var cmd *exec.Cmd
	if goRuntime.GOOS == "windows" {
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
