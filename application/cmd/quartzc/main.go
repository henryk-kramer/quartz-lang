package main

import (
	"os"

	"quartz-lang.org/internal/pkg/utils/cliargs"
)

func main() {
	cliargs.New(os.Args[1:])
}
