package cmds

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/c-bata/go-prompt"
)

func Repl() {
	clearConsole()
	printWelcome()

	for {
		t := prompt.Input("? ", completer)
		if t == "exit" {
			println()
			break
		}

		printResult(t)
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
	println("- Type 'exit' to exit.")
	println()
}

func printResult(t string) {
	println("=", t)
	println()
}
