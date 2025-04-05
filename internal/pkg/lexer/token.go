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
	MUTED                  TokenType = "Muted"
	TICK                   TokenType = "Tick"
	CIRCUMFLEX             TokenType = "Circumflex"
	PIPE                   TokenType = "Pipe"
	COMMA                  TokenType = "Comma"
	IF_NIL                 TokenType = "If nil"
	LOGICAL_AND            TokenType = "Logical end"
	LOGICAL_OR             TokenType = "Logical or"

	SINGLE_LINE_COMMENT      TokenType = "Single line comment"
	MULTI_LINE_COMMENT       TokenType = "Multi line comment"
	MULTI_LINE_COMMENT_ERROR TokenType = "Multi line comment error"
	STRING_LITERAL           TokenType = "String literal"
	STRING_LITERAL_ERROR     TokenType = "String literal error"

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
	KEYWORD_NOT             TokenType = "Keyword 'not'"
	KEYWORD_AND             TokenType = "Keyword 'and'"
	KEYWORD_OR              TokenType = "Keyword 'or'"
	KEYWORD_XOR             TokenType = "Keyword 'xor'"
	KEYWORD_SHL             TokenType = "Keyword 'shl'"
	KEYWORD_SHR             TokenType = "Keyword 'shr'"
	KEYWORD_ASHR            TokenType = "Keyword 'ashr'"
	KEYWORD_CSHL            TokenType = "Keyword 'cshl'"
	KEYWORD_CSHR            TokenType = "Keyword 'cshr'"
	KEYWORD_TRUE            TokenType = "Keyword 'true'"
	KEYWORD_FALSE           TokenType = "Keyword 'false'"
	KEYWORD_BOOL            TokenType = "Keyword 'bool'"
	KEYWORD_U8              TokenType = "Keyword 'u8'"
	KEYWORD_U16             TokenType = "Keyword 'u16'"
	KEYWORD_U32             TokenType = "Keyword 'u32'"
	KEYWORD_U64             TokenType = "Keyword 'u64'"
	KEYWORD_I8              TokenType = "Keyword 'i8'"
	KEYWORD_I16             TokenType = "Keyword 'i16'"
	KEYWORD_I32             TokenType = "Keyword 'i32'"
	KEYWORD_I64             TokenType = "Keyword 'i64'"
	KEYWORD_F32             TokenType = "Keyword 'f32'"
	KEYWORD_F64             TokenType = "Keyword 'f64e'"
	KEYWORD_NUM             TokenType = "Keyword 'num'"
	KEYWORD_SYM             TokenType = "Keyword 'sym'"
	KEYWORD_BIN             TokenType = "Keyword 'bin'"
	IDENTIFIER              TokenType = "Identifier"
	MUTED_IDENTIFIER        TokenType = "Muted identifier"

	BIN_NUM_LITERAL          TokenType = "Bin num literal"
	BIN_NUM_LITERAL_ERROR    TokenType = "Bin num literal error"
	OCT_NUM_LITERAL          TokenType = "Oct num literal"
	OCT_NUM_LITERAL_ERROR    TokenType = "Oct num literal error"
	DEC_NUM_LITERAL          TokenType = "Dec num literal"
	DEC_NUM_LITERAL_ERROR    TokenType = "Dec num literal error"
	HEX_NUM_LITERAL          TokenType = "Hex num literal"
	HEX_NUM_LITERAL_ERROR    TokenType = "Hex num literal error"
	NORMAL_NUM_LITERAL       TokenType = "Normal num literal"
	NORMAL_NUM_LITERAL_ERROR TokenType = "Normal num literal error"

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
