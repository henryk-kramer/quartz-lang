package parser

import (
	"fmt"

	"github.com/henryk-kramer/quartz-lang/internal/pkg/lexer"
)

type Error struct {
	Token lexer.Token
	Msg   string
}

type parser struct {
	tokens  []lexer.Token
	program Program
	errors  []Error
	idx     int
}

func Run(tokens []lexer.Token) (Program, []Error) {
	p := newParser(tokens)

	token := p.peekIgnoreSpace()

	if token.Type == lexer.KEYWORD_NAMESPACE {
		p.parseNamespace()
	}

	for !p.eof() {

	}

	return p.program, p.errors
}

func newParser(tokens []lexer.Token) parser {
	return parser{tokens: tokens}
}

/* Helper methods */

func (p *parser) eof() bool {
	return p.idx >= len(p.tokens)
}

func (p *parser) whitespace() bool {
	tokenType := p.peek().Type

	return tokenType == lexer.WHITESPACE || tokenType == lexer.TAB || tokenType == lexer.NEWLINE
}

func (p *parser) peek() *lexer.Token {
	if p.eof() {
		return &lexer.Token{Type: lexer.EOF}
	}

	return &p.tokens[p.idx]
}

func (p *parser) peekIgnoreSpace() *lexer.Token {
	idx := p.idx
	for !p.eof() {
		tokenType := p.tokens[idx].Type
		if tokenType == lexer.WHITESPACE || tokenType == lexer.TAB || tokenType == lexer.NEWLINE {
			idx++
			continue
		}

		return &p.tokens[idx]
	}

	return &lexer.Token{Type: lexer.EOF}
}

func (p *parser) advance() {
	p.idx++
}

func (p *parser) advanceIgnoreSpace() {
	for !p.eof() {
		tokenType := p.tokens[p.idx].Type
		p.idx++

		if tokenType == lexer.WHITESPACE || tokenType == lexer.TAB || tokenType == lexer.NEWLINE {
			continue
		}

		return
	}
}

/* Parser methods */

func (p *parser) parseNamespace() {
	var identifiers []string

	p.advanceIgnoreSpace() // skip 'namespace'

	token := p.peekIgnoreSpace()
	if token.Type != lexer.IDENTIFIER {
		p.errors = append(p.errors, Error{
			Token: *token,
			Msg:   fmt.Sprintf("Expected an identifier but found %s", token.Type),
		})
	} else {
		identifiers = append(identifiers, token.Literal)
	}

	p.advanceIgnoreSpace()

	for !p.eof() {
		if p.whitespace() {
			break
		}

		token = p.peek()
		if token.Type != lexer.DOUBLE_SEMICOLON {
			p.errors = append(p.errors, Error{
				Token: *token,
				Msg:   fmt.Sprintf("Expected :: but found %s", token.Type),
			})
		}

		p.advance()

		if p.whitespace() {
			p.errors = append(p.errors, Error{
				Token: *token,
				Msg:   "Missing an identifer",
			})
			break
		}

		token = p.peek()
		if token.Type != lexer.IDENTIFIER {
			p.errors = append(p.errors, Error{
				Token: *token,
				Msg:   fmt.Sprintf("Expected an identifier but found %s", token.Type),
			})
		} else {
			identifiers = append(identifiers, token.Literal)
		}

		p.advance()
	}

	p.program.Namespace = Namespace{
		Identifiers: identifiers,
	}
}
