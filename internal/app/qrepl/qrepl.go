package qrepl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/henryk-kramer/quartz-lang/internal/app/qrepl/cli"
	"github.com/henryk-kramer/quartz-lang/internal/pkg/lexer"
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

		for _, token := range lexer.Run(input) {
			if token.HasError {
				console.WriteError("%s", token)
			} else {
				console.WriteDebug("%s", token)
			}
		}
	}
}
