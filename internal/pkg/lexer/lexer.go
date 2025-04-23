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

func Run(text string, filename string) []Token {
	l := newLexer(text, filename)

	matchBinNum := func(ch rune) bool {
		return ch >= '0' && ch <= '1'
	}

	matchOctNum := func(ch rune) bool {
		return ch >= '0' && ch <= '7'
	}

	matchDecNum := func(ch rune) bool {
		return ch >= '0' && ch <= '9'
	}

	matchHexNum := func(ch rune) bool {
		return (ch >= '0' && ch <= '9') ||
			(ch >= 'a' && ch <= 'f') ||
			(ch >= 'A' && ch <= 'F')
	}

	for !l.eof() {
		var _ = l.parseChars("\t", TAB) ||
			l.parseNewline() ||
			l.parseWhitespace() ||
			l.parseSingleLineComment() ||
			l.parseMultiLineComment() ||
			l.parseStringLiteral() ||
			l.parseIdentifierOrKeyword() ||
			l.parseXaryNumLiteral('b', matchBinNum, BIN_NUM_LITERAL, BIN_NUM_LITERAL_ERROR) ||
			l.parseXaryNumLiteral('o', matchOctNum, OCT_NUM_LITERAL, OCT_NUM_LITERAL_ERROR) ||
			l.parseXaryNumLiteral('d', matchDecNum, DEC_NUM_LITERAL, DEC_NUM_LITERAL_ERROR) ||
			l.parseXaryNumLiteral('x', matchHexNum, HEX_NUM_LITERAL, HEX_NUM_LITERAL_ERROR) ||
			l.parseNormalNumLiteral() ||
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
			l.parseChars("::", DOUBLE_SEMICOLON) ||
			l.parseChars("=", BINDING) ||
			l.parseChars("_", MUTED) ||
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

func newLexer(text string, filename string) lexer {
	return lexer{
		runes:    []rune(text),
		currPos:  util.Position{File: filename},
		startPos: util.Position{File: filename},
	}
}

/* Helper methods */

func (l *lexer) eof() bool {
	return l.currPos.Idx >= len(l.runes)
}

func (l *lexer) peek() rune {
	if l.eof() {
		return 0
	}

	return l.runes[l.currPos.Idx]
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

func (l *lexer) commit(tokenType TokenType) {
	pos := l.startPos
	pos.Len = l.currPos.Idx - l.startPos.Idx

	token := Token{
		Type:     tokenType,
		HasError: false,
		Literal:  string(l.runes[l.startPos.Idx:l.currPos.Idx]),
		Pos:      pos,
	}

	l.tokens = append(l.tokens, token)

	l.startPos = l.currPos
}

func (l *lexer) commitErr(tokenType TokenType, errorMsg string) {
	pos := l.startPos
	pos.Len = l.currPos.Idx - l.startPos.Idx

	token := Token{
		Type:     tokenType,
		HasError: true,
		ErrorMsg: errorMsg,
		Literal:  string(l.runes[l.startPos.Idx:l.currPos.Idx]),
		Pos:      pos,
	}

	l.tokens = append(l.tokens, token)

	l.startPos = l.currPos
}

func (l *lexer) rollback() {
	l.currPos = l.startPos
}

func (l *lexer) literal() string {
	return string(l.runes[l.startPos.Idx:l.currPos.Idx])
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

	l.commit(tokenType)
	return true
}

func (l *lexer) parseNewline() bool {
	ch := l.peek()

	if ch != '\n' && ch != '\r' {
		return false
	}

	l.advance()

	l.commit(NEWLINE)
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

	l.commit(WHITESPACE)
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

	l.commit(SINGLE_LINE_COMMENT)
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
		l.commitErr(MULTI_LINE_COMMENT_ERROR, "Multi line comment not closed")
		return true
	}

	l.advance()

	l.commit(MULTI_LINE_COMMENT)
	return true
}

func (l *lexer) parseStringLiteral() bool {
	if l.peek() != '"' {
		return false
	}

	l.advance()

	var prevCh rune
	l.advanceWhile(func(ch rune) bool {
		if prevCh == '\\' {
			prevCh = 0
			return true
		}

		prevCh = ch
		return ch != '"'
	})

	if l.eof() {
		l.commitErr(STRING_LITERAL_ERROR, "String literal not closed")
		return true
	}

	l.advance()

	l.commit(STRING_LITERAL)

	return true
}

func (l *lexer) parseIdentifierOrKeyword() bool {
	matchUnderscore := func(ch rune) bool {
		return ch == '_'
	}

	matchLowerAtoZ := func(ch rune) bool {
		return ch >= 'a' && ch <= 'z'
	}

	matchUpperAtoZ := func(ch rune) bool {
		return ch >= 'A' && ch <= 'Z'
	}

	match0to9 := func(ch rune) bool {
		return ch >= '0' && ch <= '9'
	}

	firstCh := l.peek()

	if !(matchUnderscore(firstCh) || matchLowerAtoZ(firstCh) || matchUpperAtoZ(firstCh)) {
		return false
	}

	l.advanceWhile(func(ch rune) bool {
		return (matchUnderscore(ch) ||
			matchLowerAtoZ(ch) ||
			matchUpperAtoZ(ch) ||
			match0to9(ch))
	})

	if matchUnderscore(firstCh) {
		l.commit(MUTED_IDENTIFIER)
		return true
	}

	switch l.literal() {
	case "namespace":
		l.commit(KEYWORD_NAMESPACE)
	case "import":
		l.commit(KEYWORD_IMPORT)
	case "from":
		l.commit(KEYWORD_FROM)
	case "as":
		l.commit(KEYWORD_AS)
	case "let!":
		l.commit(KEYWORD_LET_EXCLAMATION)
	case "let":
		l.commit(KEYWORD_LET)
	case "const":
		l.commit(KEYWORD_CONST)
	case "ext":
		l.commit(KEYWORD_EXT)
	case "pub":
		l.commit(KEYWORD_PUB)
	case "fn":
		l.commit(KEYWORD_FN)
	case "struct":
		l.commit(KEYWORD_STRUCT)
	case "trait":
		l.commit(KEYWORD_TRAIT)
	case "impl":
		l.commit(KEYWORD_IMPL)
	case "self":
		l.commit(KEYWORD_SELF)
	case "nil":
		l.commit(KEYWORD_NIL)
	case "if":
		l.commit(KEYWORD_IF)
	case "cond":
		l.commit(KEYWORD_COND)
	case "case":
		l.commit(KEYWORD_CASE)
	case "else":
		l.commit(KEYWORD_ELSE)
	case "return":
		l.commit(KEYWORD_RETURN)
	case "not":
		l.commit(KEYWORD_NOT)
	case "and":
		l.commit(KEYWORD_AND)
	case "or":
		l.commit(KEYWORD_OR)
	case "xor":
		l.commit(KEYWORD_XOR)
	case "shl":
		l.commit(KEYWORD_SHL)
	case "shr":
		l.commit(KEYWORD_SHR)
	case "ashr":
		l.commit(KEYWORD_ASHR)
	case "cshl":
		l.commit(KEYWORD_CSHL)
	case "cshr":
		l.commit(KEYWORD_CSHR)
	case "true":
		l.commit(KEYWORD_TRUE)
	case "false":
		l.commit(KEYWORD_FALSE)
	case "bool":
		l.commit(KEYWORD_NAMESPACE)
	case "u8":
		l.commit(KEYWORD_NAMESPACE)
	case "u16":
		l.commit(KEYWORD_NAMESPACE)
	case "u32":
		l.commit(KEYWORD_NAMESPACE)
	case "u64":
		l.commit(KEYWORD_NAMESPACE)
	case "i8":
		l.commit(KEYWORD_NAMESPACE)
	case "i16":
		l.commit(KEYWORD_NAMESPACE)
	case "i32":
		l.commit(KEYWORD_NAMESPACE)
	case "i64":
		l.commit(KEYWORD_NAMESPACE)
	case "f32":
		l.commit(KEYWORD_NAMESPACE)
	case "f64":
		l.commit(KEYWORD_NAMESPACE)
	case "num":
		l.commit(KEYWORD_NAMESPACE)
	case "sym":
		l.commit(KEYWORD_NAMESPACE)
	case "bin":
		l.commit(KEYWORD_NAMESPACE)
	default:
		l.commit(IDENTIFIER)
	}

	return true
}

func (l *lexer) parseXaryNumLiteral(
	identifier rune,
	match func(rune) bool,
	tokenType TokenType,
	errTokenType TokenType,
) bool {
	errMatch := func(ch rune) bool {
		return ch == '_' ||
			(ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9')
	}

	if l.peek() != '0' {
		return false
	}

	l.advance()

	if l.peek() != identifier {
		l.rollback()
		return false
	}

	l.advance()

	var foundData bool
	l.advanceWhile(func(ch rune) bool {
		if match(ch) {
			foundData = true
			return true
		}
		return ch == '_'
	})

	if errMatch(l.peek()) {
		l.advanceWhile(errMatch)
		l.commitErr(errTokenType, "Xary number literals can't be followed by a-z, A-Z, 0-9 or _")
		return true
	}

	if !foundData {
		l.commitErr(errTokenType, "Xary number literal defined without data")
		return true
	}

	l.commit(tokenType)
	return true
}

func (l *lexer) parseNormalNumLiteral() bool {
	errMatch := func(ch rune) bool {
		return ch == '_' ||
			ch == '.' ||
			(ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9')
	}

	match0to9 := func(ch rune) bool {
		return ch >= '0' && ch <= '9'
	}

	/* Handle Integer part */

	if !match0to9(l.peek()) {
		return false
	}

	l.advanceWhile(match0to9)

	if !errMatch(l.peek()) {
		l.commit(NORMAL_NUM_LITERAL)
		return true
	}

	ch := l.peek()
	if ch != '.' && ch != 'e' {
		l.advanceWhile(errMatch)
		l.commitErr(NORMAL_NUM_LITERAL_ERROR, "Number literals cannot contain a-z, A-Z, 0-9, _ or .")
		return true
	}

	/* Handle decimal part */

	if l.peek() == '.' {
		l.advance()

		if !match0to9(l.peek()) {
			l.advanceWhile(errMatch)
			l.commitErr(NORMAL_NUM_LITERAL_ERROR, "No numbers specified after decimal point")
			return true
		}

		l.advanceWhile(match0to9)

		ch = l.peek()
		if ch != 'e' {
			if errMatch(ch) {
				l.advanceWhile(errMatch)
				l.commitErr(NORMAL_NUM_LITERAL_ERROR, "Number literals cannot contain a-z, A-Z, 0-9, _ or .")
				return true
			} else {
				l.commit(NORMAL_NUM_LITERAL)
				return true
			}
		}
	}

	/* Handle exponent part */

	l.advance()

	ch = l.peek()
	if ch == '+' || ch == '-' {
		l.advance()
	}

	if !match0to9(l.peek()) {
		l.advanceWhile(errMatch)
		l.commitErr(NORMAL_NUM_LITERAL_ERROR, "No numbers specified after exponent sign")
		return true
	}

	l.advanceWhile(match0to9)

	if errMatch(l.peek()) {
		l.advanceWhile(errMatch)
		l.commitErr(NORMAL_NUM_LITERAL_ERROR, "Number literals cannot contain a-z, A-Z, 0-9, _ or .")
		return true
	}

	l.commit(NORMAL_NUM_LITERAL)
	return true
}

func (l *lexer) parseUnknown() bool {
	l.advance()

	l.commitErr(UNKNOWN, "The specified characters are unknown")

	l.startPos = l.currPos

	return true
}
