package core

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/renatopp/x/fsx"
	"github.com/renatopp/x/strx"
)

func makeColor(r, g, b int) func(s string, args ...any) string {
	c := color.RGB(r, g, b)
	return func(s string, args ...any) string {
		return c.Sprintf(s, args...)
	}
}

var colorErrorBase = makeColor(255, 150, 100)
var colorErrorFade = makeColor(170, 170, 200)

//
// TYPES
//

// KlcError represents any error that can occur in the Klc language processing.
type KlcError struct {
	Kind    ErrorKind
	Message string
	Span    *Span
}

func NewError(kind ErrorKind, message string, span *Span) *KlcError {
	return &KlcError{
		Kind:    kind,
		Message: message,
		Span:    span,
	}
}
func NewErrorModule(kind ErrorKind, message string, script *Script) *KlcError {
	return &KlcError{
		Kind:    kind,
		Message: message,
		Span:    &Span{Script: script},
	}
}
func (e *KlcError) Error() string { return e.Message }

//
// UTILITIES
//

// PPrintError prints a formatted error message to the standard output.
func PPrintError(err error) {
	println(SPPrintError(err))
}

// SPPrintError returns a formatted error message as a string.
func SPPrintError(err error) string {
	switch e := err.(type) {
	case *KlcError:
		switch e.Kind {
		case ErrorFile:
			return sprintFileError(e)
		default:
			return sprintSyntaxError(e)
		}
	default:
		return sprintSimpleError(e)
	}
}

func sprintSimpleError(err error) string {
	ce := colorErrorBase
	str := ""
	str += fmt.Sprintf("%s: %s\n", ce("FILE ERROR"), strx.Escape(err.Error()))
	return str

}

func sprintFileError(err *KlcError) string {
	ce := colorErrorBase
	cf := colorErrorFade

	str := ""
	str += fmt.Sprintf("%s: %s\n", ce("FILE-ERROR"), strx.Escape(err.Message))
	str += fmt.Sprintf("%s%s\n", ce("└ "), cf(err.Span.Script.Path))
	return str
}

func sprintSyntaxError(err *KlcError) string {
	ce := colorErrorBase
	cf := colorErrorFade
	cu := color.New(color.FgRed)
	// cu := color.New(color.FgRed, color.Underline)

	if err.Span == nil {
		return sprintSimpleError(err)
	}

	str := ""
	lines := fsx.ForceReadFileLines(err.Span.Script.Path)
	fromLine := err.Span.FromLine
	fromColumn := err.Span.FromColumn
	toColumn := err.Span.ToColumn
	// columnSpan := max(1, toColumn-fromColumn)

	sprintNumber := func(line int) string {
		return fmt.Sprintf(" %d", line)
	}

	sprintLine := func(line int) string {
		if line < 1 || line > len(lines) {
			return ""
		}
		return lines[line-1]
	}

	sprintMainLine := func(line string) string {
		println("LINE:", strx.Escape(line))
		println("FROMLINE:", fromLine, "TOLINE:", err.Span.ToLine)
		println("FROMCOLUMN:", fromColumn, "TOCOLUMN:", toColumn)
		before := line[:fromColumn-2]
		middle := line[fromColumn-2 : toColumn]
		after := ""
		if toColumn-1 < len(line) {
			after = line[toColumn:]
		}
		return before + cu.Sprint(middle) + after
	}

	ident := strx.Ident("", len(sprintNumber(fromLine+1)))

	str += fmt.Sprintf("%s: %s\n", ce(strx.ToUpper(string(err.Kind))), strx.Escape(err.Message))
	str += "\n"
	str += fmt.Sprintf("%s┌─[%s]\n", ident, cf("%s:%d:%d", err.Span.Script.Path, fromLine, fromColumn))
	str += fmt.Sprintf("%s|\n", ident)
	str += fmt.Sprintf("%s|    %s\n", sprintNumber(fromLine-1), sprintLine(fromLine-1))
	str += fmt.Sprintf("%s|    %s\n", sprintNumber(fromLine), sprintMainLine(sprintLine(fromLine)))
	// str += fmt.Sprintf("%s|    %s\n", ident, (strings.Repeat(" ", fromColumn-1) + strings.Repeat("^", columnSpan)))
	str += fmt.Sprintf("%s|    %s\n", sprintNumber(fromLine+1), sprintLine(fromLine+1))
	str += fmt.Sprintf("%s└    \n", ident)
	str += "\n"

	return str
}

func ThrowAt(span *Span, kind ErrorKind, message string, args ...any) {
	panic(NewError(kind, fmt.Sprintf(message, args...), span))
}

func WithRecover[T any](fn func() T) (val T, err error) {
	defer func() {
		if r := recover(); r != nil {
			if recErr, ok := r.(error); ok {
				err = recErr
			} else {
				err = fmt.Errorf("panic: %v", r)
			}
		}
	}()
	return fn(), err
}
