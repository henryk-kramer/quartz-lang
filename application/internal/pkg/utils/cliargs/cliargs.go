package cliargs

import (
	"fmt"
	"strings"
)

type cliargs struct {
	args map[string][]string
}

func New(args []string) *cliargs {
	argsTree := make(map[string][]string)
	current := ""

	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			current = arg[2:]
			argsTree[current] = []string{}
			continue
		}

		argsTree[current] = append(argsTree[current], arg)
	}

	fmt.Println(argsTree)

	return &cliargs{}
}
