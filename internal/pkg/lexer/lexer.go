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
			l.parseChars("\n", NEWLINE) ||
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
			l.parseChars("::", FUNCTION_REFERENCE) ||
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
		l.commit(STRING_LITERAL_ERROR, true)
		return true
	}

	l.advance()

	l.commit(STRING_LITERAL, false)

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
		l.commit(MUTED_IDENTIFIER, false)
		return true
	}

	switch l.literal() {
	case "namespace":
		l.commit(KEYWORD_NAMESPACE, false)
	case "import":
		l.commit(KEYWORD_IMPORT, false)
	case "from":
		l.commit(KEYWORD_FROM, false)
	case "as":
		l.commit(KEYWORD_AS, false)
	case "let!":
		l.commit(KEYWORD_LET_EXCLAMATION, false)
	case "let":
		l.commit(KEYWORD_LET, false)
	case "const":
		l.commit(KEYWORD_CONST, false)
	case "pub":
		l.commit(KEYWORD_PUB, false)
	case "fn":
		l.commit(KEYWORD_FN, false)
	case "struct":
		l.commit(KEYWORD_STRUCT, false)
	case "trait":
		l.commit(KEYWORD_TRAIT, false)
	case "impl":
		l.commit(KEYWORD_IMPL, false)
	case "self":
		l.commit(KEYWORD_SELF, false)
	case "nil":
		l.commit(KEYWORD_NIL, false)
	case "if":
		l.commit(KEYWORD_IF, false)
	case "cond":
		l.commit(KEYWORD_COND, false)
	case "case":
		l.commit(KEYWORD_CASE, false)
	case "else":
		l.commit(KEYWORD_ELSE, false)
	case "return":
		l.commit(KEYWORD_RETURN, false)
	case "not":
		l.commit(KEYWORD_NOT, false)
	case "and":
		l.commit(KEYWORD_AND, false)
	case "or":
		l.commit(KEYWORD_OR, false)
	case "xor":
		l.commit(KEYWORD_XOR, false)
	case "shl":
		l.commit(KEYWORD_SHL, false)
	case "shr":
		l.commit(KEYWORD_SHR, false)
	case "ashr":
		l.commit(KEYWORD_ASHR, false)
	case "cshl":
		l.commit(KEYWORD_CSHL, false)
	case "cshr":
		l.commit(KEYWORD_CSHR, false)
	case "true":
		l.commit(KEYWORD_TRUE, false)
	case "false":
		l.commit(KEYWORD_FALSE, false)
	case "bool":
		l.commit(KEYWORD_NAMESPACE, false)
	case "u8":
		l.commit(KEYWORD_NAMESPACE, false)
	case "u16":
		l.commit(KEYWORD_NAMESPACE, false)
	case "u32":
		l.commit(KEYWORD_NAMESPACE, false)
	case "u64":
		l.commit(KEYWORD_NAMESPACE, false)
	case "i8":
		l.commit(KEYWORD_NAMESPACE, false)
	case "i16":
		l.commit(KEYWORD_NAMESPACE, false)
	case "i32":
		l.commit(KEYWORD_NAMESPACE, false)
	case "i64":
		l.commit(KEYWORD_NAMESPACE, false)
	case "f32":
		l.commit(KEYWORD_NAMESPACE, false)
	case "f64":
		l.commit(KEYWORD_NAMESPACE, false)
	case "num":
		l.commit(KEYWORD_NAMESPACE, false)
	case "sym":
		l.commit(KEYWORD_NAMESPACE, false)
	case "bin":
		l.commit(KEYWORD_NAMESPACE, false)
	default:
		l.commit(IDENTIFIER, false)
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
		l.commit(errTokenType, true)
		return true
	}

	if !foundData {
		l.commit(errTokenType, true)
		return true
	}

	l.commit(tokenType, false)
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
		l.commit(NORMAL_NUM_LITERAL, false)
		return true
	}

	ch := l.peek()
	if ch != '.' && ch != 'e' {
		l.advanceWhile(errMatch)
		l.commit(NORMAL_NUM_LITERAL_ERROR, true)
		return true
	}

	/* Handle decimal part */

	if l.peek() == '.' {
		l.advance()

		if !match0to9(l.peek()) {
			l.advanceWhile(errMatch)
			l.commit(NORMAL_NUM_LITERAL_ERROR, true)
			return true
		}

		l.advanceWhile(match0to9)

		ch = l.peek()
		if ch != 'e' {
			if errMatch(ch) {
				l.advanceWhile(errMatch)
				l.commit(NORMAL_NUM_LITERAL_ERROR, true)
				return true
			} else {
				l.commit(NORMAL_NUM_LITERAL, false)
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
		l.commit(NORMAL_NUM_LITERAL_ERROR, true)
		return true
	}

	l.advanceWhile(match0to9)

	if errMatch(l.peek()) {
		l.advanceWhile(errMatch)
		l.commit(NORMAL_NUM_LITERAL_ERROR, true)
		return true
	}

	l.commit(NORMAL_NUM_LITERAL, false)
	return true
}

func (l *lexer) parseUnknown() bool {
	l.advance()

	l.commit(UNKNOWN, true)

	l.startPos = l.currPos

	return true
}
