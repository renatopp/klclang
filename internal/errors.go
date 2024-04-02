package internal

import (
	"errors"
	"fmt"
	"strings"

	"github.com/renatopp/langtools/lexers"
)

func ConvertLexerErrors(input []byte, err []lexers.LexerError) error {
	result := ""
	for i, e := range err {
		if i == 0 {
			result += convertLexerMainError(input, e)

		} else {
			if i == 1 {
				result += "\n"
				result += "Other errors:\n"
			}

			result += convertLexerSecondaryError(input, e)
		}
	}

	return errors.New(result)
}

func convertLexerMainError(input []byte, err lexers.LexerError) string {
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
	result += fmt.Sprintf("[syntax error] %s\n", err.Msg)

	return result
}

func convertLexerSecondaryError(input []byte, err lexers.LexerError) string {
	line, column := err.At()
	return fmt.Sprintf("line %4d, column %4d | %s", line, column, err.Msg)
}

// func convertParserErrors(input []byte, err []error) error     { return nil }
// func convertMainRuntimeError(input []byte, err []error) error { return nil }

func getLines(input []byte, from, to int) []string {
	lines := strings.Split(string(input), "\n")
	return lines[from:to]
}

func padLeft(msg string, length int) string {
	return strings.Repeat(" ", length-len(msg)) + msg
}
