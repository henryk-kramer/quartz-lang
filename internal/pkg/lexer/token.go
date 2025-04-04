package lexer

import (
	"fmt"

	"github.com/henryk-kramer/quartz-lang/internal/pkg/util"
)

type TokenType string

const (
	NEWLINE    TokenType = "Newline"
	TAB        TokenType = "Tab"
	WHITESPACE TokenType = "Whitespace"

	OPENED_PARENTHESIS TokenType = "Opened parenthesis"
	CLOSED_PARENTHESIS TokenType = "Closed parenthesis"
	OPENED_BRACE       TokenType = "Opened brace"
	CLOSED_BRACE       TokenType = "Closed brace"
	OPENED_BRACKET     TokenType = "Opened bracket"
	CLOSED_BRACKET     TokenType = "Closed bracket"

	FUNCTION_REFERENCE     TokenType = "Function reference"
	PLUS_SIGN              TokenType = "Plus sign"
	MINUS_SIGN             TokenType = "Minus sign"
	STAR_SIGN              TokenType = "Star sign"
	SLASH_SIGN             TokenType = "Slash sign"
	LESS_THAN_OR_EQUALS    TokenType = "Less than or equals"
	LESS_THAN              TokenType = "Less than"
	GREATER_THAN_OR_EQUALS TokenType = "Greater than or equals"
	GREATER_THAN           TokenType = "Greater than"
	EQUALS                 TokenType = "Equals"
	NOT_EQUALS             TokenType = "Not equals"
	TYPE_INDICATOR         TokenType = "Type indicator"
	RETURN_TYPE_INDICATOR  TokenType = "Return type indicator"
	BINDING                TokenType = "Binding"
	TICK                   TokenType = "Tick"
	CIRCUMFLEX             TokenType = "Circumflex"
	PIPE                   TokenType = "Pipe"
	COMMA                  TokenType = "Comma"
	IF_NIL                 TokenType = "If nil"
	LOGICAL_AND            TokenType = "Logical end"
	LOGICAL_OR             TokenType = "Logical or"
	BINARY_NOT             TokenType = "Binary not"
	BINARY_AND             TokenType = "Binary and"
	BINARY_OR              TokenType = "Binary or"
	BINARY_XOR             TokenType = "Binary xor"
	BINARY_SHL             TokenType = "Binary logical shift left"
	BINARY_SHR             TokenType = "Binary logical shift right"
	BINARY_ASHR            TokenType = "Binary arithmetic shift right"
	BINARY_CSHL            TokenType = "Binary circular shift left"
	BINARY_CSHR            TokenType = "Binary circular shift right"

	SINGLE_LINE_COMMENT      TokenType = "Single line comment"
	MULTI_LINE_COMMENT       TokenType = "Multi line comment"
	MULTI_LINE_COMMENT_ERROR TokenType = "Multi line comment error"

	KEYWORD_NAMESPACE       TokenType = "Keyword 'namespace'"
	KEYWORD_IMPORT          TokenType = "Keyword 'import'"
	KEYWORD_AS              TokenType = "Keyword 'as'"
	KEYWORD_FROM            TokenType = "Keyword 'from'"
	KEYWORD_TYPE            TokenType = "Keyword 'type'"
	KEYWORD_LET_EXCLAMATION TokenType = "Keyword 'let!'"
	KEYWORD_LET             TokenType = "Keyword 'let'"
	KEYWORD_CONST           TokenType = "Keyword 'const'"
	KEYWORD_PUB             TokenType = "Keyword 'pub'"
	KEYWORD_FN              TokenType = "Keyword 'fn'"
	KEYWORD_STRUCT          TokenType = "Keyword 'struct'"
	KEYWORD_TRAIT           TokenType = "Keyword 'trait'"
	KEYWORD_IMPL            TokenType = "Keyword 'impl'"
	KEYWORD_SELF            TokenType = "Keyword 'self'"
	KEYWORD_NIL             TokenType = "Keyword 'nil'"
	KEYWORD_IF              TokenType = "Keyword 'if'"
	KEYWORD_COND            TokenType = "Keyword 'cond'"
	KEYWORD_CASE            TokenType = "Keyword 'case'"
	KEYWORD_ELSE            TokenType = "Keyword 'else'"
	KEYWORD_RETURN          TokenType = "Keyword 'return'"

	UNKNOWN TokenType = "Unknown"
)

type Token struct {
	Type     TokenType
	HasError bool
	Literal  string
	Pos      util.Position
}

func (token Token) String() string {
	return fmt.Sprintf(
		"{ (i%d r%d c%d) %s: %q } ",
		token.Pos.Idx,
		token.Pos.Row,
		token.Pos.Col,
		token.Type,
		token.Literal,
	)
}
