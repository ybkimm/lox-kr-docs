package codegen

import (
	gotoken "go/token"
	"io"

	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
)

const prefix = "_lx"

const lexerGenGoName = "lexer.gen.go"

type Config struct {
	Errs           *errlogger.ErrLogger
	ImplDir        string
	Grammar        *grammar.AugmentedGrammar
	AnalysisWriter io.Writer
	AnalysisOnly   bool
}

func Generate(cfg *Config) {
	pgen := NewParserGenState(cfg.ImplDir, cfg.Grammar, cfg.Errs)
	lgen := NewLexerGenState(cfg.ImplDir, cfg.Grammar)

	pgen.ConstructParseTables()
	if cfg.Errs.HasError() {
		return
	}

	if cfg.AnalysisWriter != nil {
		pgen.ParserTable.Print(cfg.AnalysisWriter)
	}

	if cfg.AnalysisOnly {
		return
	}

	// TODO: check for conflicts

	err := lgen.Generate()
	if err != nil {
		cfg.Errs.Errorf(gotoken.Position{}, "failed to generate lexer.gen.go: %v", err)
		return
	}

	pgen.ParseGo()
	if cfg.Errs.HasError() {
		return
	}

	pgen.MapReduceActions()
	if cfg.Errs.HasError() {
		return
	}

	err = pgen.Generate()
	if err != nil {
		cfg.Errs.Errorf(gotoken.Position{}, "failed to emit parser.gen.go: %v", err)
		return
	}
}
