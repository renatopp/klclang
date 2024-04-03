package internal

import (
	"errors"
	"fmt"
	"strings"

	"github.com/renatopp/klclang/internal/runtime"
	"github.com/renatopp/langtools"
	"github.com/renatopp/langtools/lexers"
	"github.com/renatopp/langtools/parsers"
)

func ConvertLexerErrors(input []byte, err []lexers.LexerError) error {
	var errs []langtools.Error = make([]langtools.Error, len(err))
	for i, e := range err {
		errs[i] = e
	}

	return convertError("lexer", input, errs)
}

func ConvertParserErrors(input []byte, err []parsers.ParserError) error {
	var errs []langtools.Error = make([]langtools.Error, len(err))
	for i, e := range err {
		errs[i] = e
	}

	return convertError("parser", input, errs)
}

func ConvertRuntimeErrors(input []byte, err []runtime.RuntimeError) error {
	return errors.New(convertMainError("runtime", input, err[0]))
}

func convertError(tp string, input []byte, err []langtools.Error) error {
	result := ""
	for i, e := range err {
		if i == 0 {
			result += convertMainError(tp, input, e)

		} else {
			if i == 1 {
				result += "\n"
				result += "Other errors:\n"
			}

			result += convertSecondaryError(input, e)
		}
	}

	return errors.New(result)
}

func convertMainError(tp string, input []byte, err langtools.Error) string {
	line, column := err.At()

	fromLine := max(0, line-2)
	toLine := line

	// Header
	result := fmt.Sprintf("File [stdin], line %d, column %d\n", line, column)

	// Code
	lines := getLines(input, fromLine, toLine)
	for i, l := range lines {
		result += fmt.Sprintf("%s| %s\n", padLeft(fmt.Sprintf("%d", fromLine+i+1), 4), l)
	}

	// Error Pointer
	result += fmt.Sprintf("%s| %s\n", padLeft("", 4), padLeft("", column-1)+"^")

	// Error Message
	result += fmt.Sprintf("[%s error] %s\n", tp, err.Error())

	return result
}

func convertSecondaryError(_ []byte, err langtools.Error) string {
	line, column := err.At()
	return fmt.Sprintf("line %4d, column %4d | %s", line, column, err.Error())
}

func getLines(input []byte, from, to int) []string {
	lines := strings.Split(string(input), "\n")
	return lines[from:to]
}

func padLeft(msg string, length int) string {
	return strings.Repeat(" ", length-len(msg)) + msg
}
