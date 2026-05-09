package internal

import (
	"os"
	"text/tabwriter"

	"github.com/renatopp/klclang/internal/core"
	"github.com/renatopp/klclang/internal/frontend"
	"github.com/renatopp/x/logx"
)

type Runner struct {
	PrintLex   bool
	PrintParse bool

	memory *core.Script
	Cache  map[string]*core.Script
	// runtime here
}

func NewRunner() *Runner {
	memory := &core.Script{
		CanonicalName: "memory",
		Path:          "memory",
		Content:       []byte{},
	}

	return &Runner{
		memory: memory,
		Cache: map[string]*core.Script{
			"memory": memory,
		},
	}
}

func (r *Runner) Run(path string) (any, error) {

	return nil, nil
}

func (r *Runner) Eval(content []byte) (any, error) {
	r.memory.Content = content

	tokens, err := frontend.Lex(r.memory, content)
	if err != nil {
		return nil, err
	}

	writer := tabwriter.NewWriter(os.Stdout, 1, 2, 1, ' ', 0)
	writer.Write([]byte("(LINE,COLUMN)\tKIND\tLITERAL\n"))
	if r.PrintLex {
		for _, token := range tokens {
			s := logx.Sprintln("(%d,%d)\t%s\t%q", token.Span.FromLine, token.Span.FromColumn, token.Kind, token.Literal)
			writer.Write([]byte(s))
		}
	}
	writer.Flush()

	return nil, nil
}
