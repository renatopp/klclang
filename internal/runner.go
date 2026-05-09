package internal

import (
	"github.com/renatopp/klclang/internal/core"
	"github.com/renatopp/klclang/internal/frontend"
	"github.com/renatopp/klclang/internal/utils"
	"github.com/renatopp/x/fsx"
	"github.com/renatopp/x/strx"
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
	content, err := fsx.ReadFile(path)
	if err != nil {
		return nil, err
	}

	script := &core.Script{
		CanonicalName: path,
		Path:          path,
		Content:       content,
	}

	return r.eval(script)
}

func (r *Runner) Eval(content []byte) (any, error) {
	r.memory.Content = content
	return r.eval(r.memory)
}

func (r *Runner) eval(script *core.Script) (any, error) {
	// LEXING
	//
	tokens, err := frontend.Lex(script, script.Content)
	if err != nil {
		return nil, err
	}

	table := utils.NewTableWriter("POS", "KIND", "LITERAL")
	table.Title("LEXER")
	for _, token := range tokens {
		table.Write(
			strx.Format("(%d,%d)", token.Span.FromLine, token.Span.FromColumn),
			strx.Format("%s", token.Kind),
			strx.Format("%q", token.Literal),
		)
	}
	println(table.Render())

	// PARSING
	//

	return nil, nil
}
