package main

import (
	"bufio"
	"klc/lang/eval"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		code := strings.Join(os.Args[1:], " ")
		run(code)
		return
	}

	// check if there is somethinig to read on STDIN
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		var stdin []byte
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stdin = append(stdin, scanner.Bytes()...)
			stdin = append(stdin, '\n')
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		run(string(stdin))
	}
}

func run(code string) {
	eval.Run(code)
}
