package main

import (
	"flag"

	"github.com/henryk-kramer/quartz-lang/internal/app/quartzc"
)

func main() {
	var cwd = flag.String("cwd", "", "Set the current working directory")
	var printLexerOutput = flag.Bool("lexer-output", false, "Print output of lexer")
	var printParserOutput = flag.Bool("parser-output", false, "Print output of parser")
	flag.Parse()

	quartzc.Run(*cwd, *printLexerOutput, *printParserOutput)
}
