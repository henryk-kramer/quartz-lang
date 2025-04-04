package cli

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Cli struct {
	sc bufio.Scanner
}

func New(sc bufio.Scanner) *Cli {
	return &Cli{sc: sc}
}

func (cli *Cli) Read() string {
	var input strings.Builder

	fmt.Print(">>> ")

	for {
		cli.sc.Scan()
		text := cli.sc.Text()
		text = strings.TrimSpace(text)
		text, hasNext := strings.CutSuffix(text, "\\")
		text = strings.TrimSpace(text)

		input.WriteString(text)
		input.WriteString("\n")

		if !hasNext {
			break
		}

		fmt.Print("--> ")
	}

	return strings.TrimSpace(input.String())
}

func (cli *Cli) Write(text string, a ...any) {
	fmt.Printf(text, a...)
}

func (cli *Cli) WriteSuccess(text string, a ...any) {
	color.New(color.FgHiGreen).PrintfFunc()(text+"\n", a...)
}

func (cli *Cli) WriteDebug(text string, a ...any) {
	color.New(color.FgCyan).PrintfFunc()(text+"\n", a...)
}

func (cli *Cli) WriteWarning(text string, a ...any) {
	color.New(color.FgYellow).PrintfFunc()(text+"\n", a...)
}

func (cli *Cli) WriteError(text string, a ...any) {
	color.New(color.FgRed).PrintfFunc()(text+"\n", a...)
}
