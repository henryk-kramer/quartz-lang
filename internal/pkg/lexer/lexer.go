package lexer

import (
	"unicode"

	"github.com/henryk-kramer/quartz-lang/internal/pkg/util"
	"github.com/henryk-kramer/quartz-lang/internal/pkg/util/array"
)

type lexer struct {
	runes    []rune
	tokens   []Token
	startPos util.Position
	currPos  util.Position
}

func Run(text string) []Token {
	l := newLexer(text)

	for !l.eof() {
		var _ = l.parseChars("\t", TAB) ||
			l.parseChars("\n", NEWLINE) ||
			l.parseWhitespace() ||
			l.parseSingleLineComment() ||
			l.parseMultiLineComment() ||
			l.parseChars("namespace", KEYWORD_TYPE) ||
			l.parseChars("import", KEYWORD_TYPE) ||
			l.parseChars("as", KEYWORD_TYPE) ||
			l.parseChars("from", KEYWORD_TYPE) ||
			l.parseChars("let!", KEYWORD_LET_EXCLAMATION) ||
			l.parseChars("let", KEYWORD_LET) ||
			l.parseChars("const", KEYWORD_CONST) ||
			l.parseChars("pub", KEYWORD_PUB) ||
			l.parseChars("fn", KEYWORD_FN) ||
			l.parseChars("struct", KEYWORD_STRUCT) ||
			l.parseChars("trait", KEYWORD_TRAIT) ||
			l.parseChars("impl", KEYWORD_IMPL) ||
			l.parseChars("self", KEYWORD_SELF) ||
			l.parseChars("nil", KEYWORD_NIL) ||
			l.parseChars("if", KEYWORD_IF) ||
			l.parseChars("cond", KEYWORD_COND) ||
			l.parseChars("case", KEYWORD_CASE) ||
			l.parseChars("else", KEYWORD_ELSE) ||
			l.parseChars("return", KEYWORD_RETURN) ||
			l.parseChars("not", BINARY_NOT) ||
			l.parseChars("and", BINARY_AND) ||
			l.parseChars("or", BINARY_OR) ||
			l.parseChars("xor", BINARY_XOR) ||
			l.parseChars("shl", BINARY_SHL) ||
			l.parseChars("shr", BINARY_SHR) ||
			l.parseChars("ashr", BINARY_ASHR) ||
			l.parseChars("cshl", BINARY_CSHL) ||
			l.parseChars("cshr", BINARY_CSHR) ||
			l.parseChars("<=", LESS_THAN_OR_EQUALS) ||
			l.parseChars("<", LESS_THAN) ||
			l.parseChars(">=", GREATER_THAN_OR_EQUALS) ||
			l.parseChars(">", GREATER_THAN) ||
			l.parseChars("==", EQUALS) ||
			l.parseChars("!=", NOT_EQUALS) ||
			l.parseChars("->", RETURN_TYPE_INDICATOR) ||
			l.parseChars("??", IF_NIL) ||
			l.parseChars("&&", IF_NIL) ||
			l.parseChars("||", IF_NIL) ||
			l.parseChars("::", FUNCTION_REFERENCE) ||
			l.parseChars("=", BINDING) ||
			l.parseChars("'", TICK) ||
			l.parseChars("^", CIRCUMFLEX) ||
			l.parseChars("|", PIPE) ||
			l.parseChars(",", COMMA) ||
			l.parseChars(":", TYPE_INDICATOR) ||
			l.parseChars("+", PLUS_SIGN) ||
			l.parseChars("-", MINUS_SIGN) ||
			l.parseChars("*", STAR_SIGN) ||
			l.parseChars("/", SLASH_SIGN) ||
			l.parseChars("(", OPENED_PARENTHESIS) ||
			l.parseChars(")", CLOSED_PARENTHESIS) ||
			l.parseChars("{", OPENED_BRACE) ||
			l.parseChars("}", CLOSED_BRACE) ||
			l.parseChars("[", OPENED_BRACKET) ||
			l.parseChars("]", CLOSED_BRACKET) ||
			l.parseUnknown()
	}

	var tokens []Token
	var lastTokenType TokenType
	for _, token := range l.tokens {
		if lastTokenType == UNKNOWN && token.Type == UNKNOWN {
			lastToken := &tokens[len(tokens)-1]
			lastToken.Pos.Len++
			lastToken.Literal += token.Literal
			continue
		}

		lastTokenType = token.Type
		tokens = append(tokens, token)
	}

	return tokens
}

func newLexer(text string) lexer {
	return lexer{runes: []rune(text)}
}

/* Helper methods */

func (l *lexer) eof() bool {
	return l.currPos.Idx >= len(l.runes)
}

func (l *lexer) peek() rune {
	if l.eof() {
		return 0
	}

	ch := l.runes[l.currPos.Idx]

	if ch == '\r' {
		return '\n'
	}

	return ch
}

func (l *lexer) advance() rune {
	if l.eof() {
		return 0
	}

	ch := l.peek()

	l.currPos.Idx++
	l.currPos.Col++

	if ch == '\n' {
		if l.peek() == '\r' {
			l.currPos.Idx++
		}

		l.currPos.Row++
		l.currPos.Col = 0

		return '\n'
	}

	if ch == '\r' {
		if l.peek() == '\n' {
			l.currPos.Idx++
		}

		l.currPos.Row++
		l.currPos.Col = 0

		return '\n'
	}

	return ch
}

func (l *lexer) advanceWhile(match func(rune) bool) {
	for {
		if l.eof() {
			return
		}

		ch := l.peek()

		if !match(ch) {
			return
		}

		l.advance()
	}
}

func (l *lexer) commit(tokenType TokenType, hasError bool) {
	pos := l.startPos
	pos.Len = l.currPos.Idx - l.startPos.Idx

	token := Token{
		Type:     tokenType,
		HasError: hasError,
		Literal:  string(l.runes[l.startPos.Idx:l.currPos.Idx]),
		Pos:      pos,
	}

	l.tokens = append(l.tokens, token)

	l.startPos = l.currPos
}

func (l *lexer) rollback() {
	l.currPos = l.startPos
}

/* Parse methods*/

func (l *lexer) parseChars(chars string, tokenType TokenType) bool {
	for _, ch := range chars {
		if array.GetOrDefault(l.runes, l.currPos.Idx, 0) != ch {
			l.rollback()
			return false
		}

		l.advance()
	}

	l.commit(tokenType, false)
	return true
}

func (l *lexer) parseWhitespace() bool {
	ch := l.peek()

	if !unicode.IsSpace(ch) {
		return false
	}

	l.advanceWhile(func(ch rune) bool {
		return unicode.IsSpace(ch)
	})

	l.commit(WHITESPACE, false)
	return true
}

func (l *lexer) parseSingleLineComment() bool {
	if l.peek() != '/' {
		return false
	}

	l.advance()

	if l.peek() != '/' {
		l.rollback()
		return false
	}

	l.advanceWhile(func(ch rune) bool {
		return ch != '\n'
	})

	l.commit(SINGLE_LINE_COMMENT, false)
	return true
}

func (l *lexer) parseMultiLineComment() bool {
	if l.peek() != '/' {
		return false
	}

	l.advance()

	if l.peek() != '*' {
		l.rollback()
		return false
	}

	var prevCh rune
	l.advanceWhile(func(ch rune) bool {
		if prevCh == '*' && ch == '/' {
			return false
		}

		prevCh = ch
		return true
	})

	if l.eof() {
		l.commit(MULTI_LINE_COMMENT_ERROR, true)
		return true
	}

	l.advance()

	l.commit(MULTI_LINE_COMMENT, false)
	return true
}

func (l *lexer) parseUnknown() bool {
	l.advance()

	l.commit(UNKNOWN, true)

	l.startPos = l.currPos

	return true
}
