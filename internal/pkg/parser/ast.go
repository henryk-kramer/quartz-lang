package parser

type Program struct {
	Namespace Namespace
}

type Namespace struct {
	Identifiers []string
}
