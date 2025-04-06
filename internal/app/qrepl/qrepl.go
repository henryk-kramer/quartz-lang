package qrepl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/henryk-kramer/quartz-lang/internal/pkg/cli"
	"github.com/henryk-kramer/quartz-lang/internal/pkg/lexer"
	"github.com/henryk-kramer/quartz-lang/internal/pkg/parser"
)

func Run(
	printLexerOutput bool,
	printParserOutput bool,
) {
	var console = cli.New(*bufio.NewScanner(os.Stdin))

	for {
		input := console.Read()

		if strings.ToLower(input) == "clear" {
			fmt.Print("\033[H\033[2J")
			continue
		}

		tokens := lexer.Run(input, "CLI")

		if printLexerOutput {
			console.WriteDebug("---- Lexer Tokens ----")
			for _, token := range tokens {
				if token.HasError {
					console.WriteError("%s", token)
				} else {
					console.WriteDebug("%s", token)
				}
			}
		}

		program, errors := parser.Run(tokens)

		if printParserOutput {
			console.WriteDebug("---- Parser AST ----")
			console.WriteDebug("%s", program)
			console.WriteError("%s", errors)
		}
	}
}
