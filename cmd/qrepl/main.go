package main

import (
	"flag"

	"github.com/henryk-kramer/quartz-lang/internal/app/qrepl"
)

func main() {
	var printLexerOutput = flag.Bool("lexer-output", false, "Print output of lexer")
	var printParserOutput = flag.Bool("parser-output", false, "Print output of parser")
	flag.Parse()

	qrepl.Run(*printLexerOutput, *printParserOutput)
}
