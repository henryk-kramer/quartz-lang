package quartzc

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/henryk-kramer/quartz-lang/internal/pkg/cli"
	"github.com/henryk-kramer/quartz-lang/internal/pkg/lexer"
	"github.com/henryk-kramer/quartz-lang/internal/pkg/parser"
)

func Run(
	userDefinedCwd string,
	printLexerOutput bool,
	printParserOutput bool,
) {
	cwd, _ := os.Getwd()
	cwd = filepath.Join(cwd, userDefinedCwd)
	os.Chdir(cwd)

	var filePaths []string
	filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if !strings.HasSuffix(d.Name(), ".ql") {
			return nil
		}

		filePaths = append(filePaths, filepath.Join(cwd, path))
		return nil
	})

	var console = cli.New(*bufio.NewScanner(os.Stdin))

	for _, filePath := range filePaths {
		contentBytes, _ := os.ReadFile(filePath)
		content := string(contentBytes)

		tokens := lexer.Run(content, filePath)

		if printLexerOutput {
			console.WriteDebug("---- Lexer Tokens ----")
			for _, token := range tokens {
				if token.Type == lexer.WHITESPACE || token.Type == lexer.NEWLINE || token.Type == lexer.TAB {
					continue
				}

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
