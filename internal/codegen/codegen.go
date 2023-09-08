package codegen

import (
	gotoken "go/token"
	"io"

	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/parsergen/lr1"
)

const prefix = "_"

type Config struct {
	Errs           *errlogger.ErrLogger
	ImplDir        string
	Grammar        *grammar.AugmentedGrammar
	ParserTable    *lr1.ParserTable
	AnalysisWriter io.Writer
	AnalysisOnly   bool
}

func Generate(cfg *Config) {
	lgen := newBaseGen(cfg.ImplDir, cfg.Grammar)

	err := lgen.Generate()
	if err != nil {
		cfg.Errs.Errorf(gotoken.Position{}, "failed to generate base.gen.go: %v", err)
		return
	}

	pgen := newParserGen(cfg.ImplDir, cfg.Grammar, cfg.ParserTable, cfg.Errs)
	if !pgen.Generate() {
		return
	}
}
