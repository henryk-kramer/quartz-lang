package lexer_test

import (
	"reflect"
	"testing"

	"github.com/henryk-kramer/quartz-lang/internal/pkg/lexer"
	"github.com/henryk-kramer/quartz-lang/internal/pkg/util"
)

type testStruct struct {
	input string
	want  []lexer.Token
}

func testHelper(t *testing.T, tests []testStruct) {
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			tokens := lexer.Run(test.input, "")

			if len(tokens) != len(test.want) {
				t.Errorf("\n%s\n%s\n%s\n%s",
					"---- EXPECTED ----",
					test.want,
					"---- ACTUAL ----",
					tokens,
				)
				return
			}

			for idx, token := range tokens {
				matchingType := token.Type == test.want[idx].Type
				matchingHasError := token.HasError == test.want[idx].HasError
				matchingLiteral := token.Literal == test.want[idx].Literal
				matchingPos := reflect.DeepEqual(token.Pos, test.want[idx].Pos)

				if !matchingType || !matchingHasError || !matchingLiteral || !matchingPos {
					t.Errorf("\n%s\nidx: %d\ntype: %t\nerror: %t\nliteral: %t\npos: %t\n%s\n%s\n%s\n%s",
						"---- INFO ----",
						idx,
						matchingType,
						matchingHasError,
						matchingLiteral,
						matchingPos,
						"---- EXPECTED ----",
						test.want,
						"---- ACTUAL ----",
						tokens,
					)
					break
				}
			}
		})
	}
}

func TestWhitespace(t *testing.T) {
	testHelper(t, []testStruct{
		{
			"\n\t",
			[]lexer.Token{
				{Type: lexer.NEWLINE, HasError: false, Literal: "\n", Pos: util.Position{Idx: 0, Col: 0, Row: 0, Len: 1}},
				{Type: lexer.TAB, HasError: false, Literal: "\t", Pos: util.Position{Idx: 1, Col: 0, Row: 1, Len: 1}},
			},
		},
		{
			"\n\t\n\t",
			[]lexer.Token{
				{Type: lexer.NEWLINE, HasError: false, Literal: "\n", Pos: util.Position{Idx: 0, Col: 0, Row: 0, Len: 1}},
				{Type: lexer.TAB, HasError: false, Literal: "\t", Pos: util.Position{Idx: 1, Col: 0, Row: 1, Len: 1}},
				{Type: lexer.NEWLINE, HasError: false, Literal: "\n", Pos: util.Position{Idx: 2, Col: 1, Row: 1, Len: 1}},
				{Type: lexer.TAB, HasError: false, Literal: "\t", Pos: util.Position{Idx: 3, Col: 0, Row: 2, Len: 1}},
			},
		},
		{
			"\r\n\t",
			[]lexer.Token{
				{Type: lexer.NEWLINE, HasError: false, Literal: "\r\n", Pos: util.Position{Idx: 0, Col: 0, Row: 0, Len: 2}},
				{Type: lexer.TAB, HasError: false, Literal: "\t", Pos: util.Position{Idx: 2, Col: 0, Row: 1, Len: 1}},
			},
		},
		{
			"\t\r\n",
			[]lexer.Token{
				{Type: lexer.TAB, HasError: false, Literal: "\t", Pos: util.Position{Idx: 0, Col: 0, Row: 0, Len: 1}},
				{Type: lexer.NEWLINE, HasError: false, Literal: "\r\n", Pos: util.Position{Idx: 1, Col: 1, Row: 0, Len: 2}},
			},
		},
	})
}

func TestParseSingleLineComment(t *testing.T) {
	testHelper(t, []testStruct{
		{"//", []lexer.Token{{Type: lexer.SINGLE_LINE_COMMENT, Literal: "//", HasError: false, Pos: util.Position{Len: 2}}}},
		{"//test", []lexer.Token{{Type: lexer.SINGLE_LINE_COMMENT, Literal: "//test", HasError: false, Pos: util.Position{Len: 6}}}},
	})
}

func TestParseMultiLineComment(t *testing.T) {
	testHelper(t, []testStruct{
		// Correct
		{"/**/", []lexer.Token{{Type: lexer.MULTI_LINE_COMMENT, Literal: "/**/", HasError: false, Pos: util.Position{Len: 4}}}},
		{"/*test*/", []lexer.Token{{Type: lexer.MULTI_LINE_COMMENT, Literal: "/*test*/", HasError: false, Pos: util.Position{Len: 8}}}},
		{"/*\n*/", []lexer.Token{{Type: lexer.MULTI_LINE_COMMENT, Literal: "/*\n*/", HasError: false, Pos: util.Position{Len: 5}}}},

		// Wrong
		{"/*", []lexer.Token{{Type: lexer.MULTI_LINE_COMMENT_ERROR, Literal: "/*", HasError: true, Pos: util.Position{Len: 2}}}},
		{"/*test", []lexer.Token{{Type: lexer.MULTI_LINE_COMMENT_ERROR, Literal: "/*test", HasError: true, Pos: util.Position{Len: 6}}}},
	})
}

func TestParseString(t *testing.T) {
	testHelper(t, []testStruct{
		// Correct
		{"\"\"", []lexer.Token{{Type: lexer.STRING_LITERAL, Literal: "\"\"", HasError: false, Pos: util.Position{Len: 2}}}},
		{"\"test\"", []lexer.Token{{Type: lexer.STRING_LITERAL, Literal: "\"test\"", HasError: false, Pos: util.Position{Len: 6}}}},
		{"\"\n\"", []lexer.Token{{Type: lexer.STRING_LITERAL, Literal: "\"\n\"", HasError: false, Pos: util.Position{Len: 3}}}},
		{"\"\\\"\"", []lexer.Token{{Type: lexer.STRING_LITERAL, Literal: "\"\\\"\"", HasError: false, Pos: util.Position{Len: 4}}}},

		// Wrong
		{"\"", []lexer.Token{{Type: lexer.STRING_LITERAL_ERROR, Literal: "\"", HasError: true, Pos: util.Position{Len: 1}}}},
		{"\"test", []lexer.Token{{Type: lexer.STRING_LITERAL_ERROR, Literal: "\"test", HasError: true, Pos: util.Position{Len: 5}}}},
	})
}

func TestParseIdentifierOrKeyword(t *testing.T) {
	testHelper(t, []testStruct{
		{"let", []lexer.Token{{Type: lexer.KEYWORD_LET, Literal: "let", HasError: false, Pos: util.Position{Len: 3}}}},
		{"_let", []lexer.Token{{Type: lexer.MUTED_IDENTIFIER, Literal: "_let", HasError: false, Pos: util.Position{Len: 4}}}},
		{"letme", []lexer.Token{{Type: lexer.IDENTIFIER, Literal: "letme", HasError: false, Pos: util.Position{Len: 5}}}},
	})
}

func TestParseXaryNumLiteral(t *testing.T) {
	testHelper(t, []testStruct{
		// Correct
		{"0b0", []lexer.Token{{Type: lexer.BIN_NUM_LITERAL, Literal: "0b0", HasError: false, Pos: util.Position{Len: 3}}}},
		{"0b_0", []lexer.Token{{Type: lexer.BIN_NUM_LITERAL, Literal: "0b_0", HasError: false, Pos: util.Position{Len: 4}}}},
		{"0b0_", []lexer.Token{{Type: lexer.BIN_NUM_LITERAL, Literal: "0b0_", HasError: false, Pos: util.Position{Len: 4}}}},
		{"0b_0_", []lexer.Token{{Type: lexer.BIN_NUM_LITERAL, Literal: "0b_0_", HasError: false, Pos: util.Position{Len: 5}}}},

		// No data error
		{"0b", []lexer.Token{{Type: lexer.BIN_NUM_LITERAL_ERROR, Literal: "0b", HasError: true, Pos: util.Position{Len: 2}}}},
		{"0b_", []lexer.Token{{Type: lexer.BIN_NUM_LITERAL_ERROR, Literal: "0b_", HasError: true, Pos: util.Position{Len: 3}}}},
		{"0b__", []lexer.Token{{Type: lexer.BIN_NUM_LITERAL_ERROR, Literal: "0b__", HasError: true, Pos: util.Position{Len: 4}}}},

		// Wrong format error
		{"0b2", []lexer.Token{{Type: lexer.BIN_NUM_LITERAL_ERROR, Literal: "0b2", HasError: true, Pos: util.Position{Len: 3}}}},
		{"0b0z", []lexer.Token{{Type: lexer.BIN_NUM_LITERAL_ERROR, Literal: "0b0z", HasError: true, Pos: util.Position{Len: 4}}}},
		{"0b0_z", []lexer.Token{{Type: lexer.BIN_NUM_LITERAL_ERROR, Literal: "0b0_z", HasError: true, Pos: util.Position{Len: 5}}}},
		{"0b0z_", []lexer.Token{{Type: lexer.BIN_NUM_LITERAL_ERROR, Literal: "0b0z_", HasError: true, Pos: util.Position{Len: 5}}}},
		{"0b0z0", []lexer.Token{{Type: lexer.BIN_NUM_LITERAL_ERROR, Literal: "0b0z0", HasError: true, Pos: util.Position{Len: 5}}}},
	})
}

func TestParseNormalNumLiteral(t *testing.T) {
	testHelper(t, []testStruct{
		// Correct
		{"0", []lexer.Token{{Type: lexer.NORMAL_NUM_LITERAL, Literal: "0", HasError: false, Pos: util.Position{Len: 1}}}},
		{"00", []lexer.Token{{Type: lexer.NORMAL_NUM_LITERAL, Literal: "00", HasError: false, Pos: util.Position{Len: 2}}}},
		{"0.0", []lexer.Token{{Type: lexer.NORMAL_NUM_LITERAL, Literal: "0.0", HasError: false, Pos: util.Position{Len: 3}}}},
		{"00.00", []lexer.Token{{Type: lexer.NORMAL_NUM_LITERAL, Literal: "00.00", HasError: false, Pos: util.Position{Len: 5}}}},
		{"0e0", []lexer.Token{{Type: lexer.NORMAL_NUM_LITERAL, Literal: "0e0", HasError: false, Pos: util.Position{Len: 3}}}},
		{"0.0e0", []lexer.Token{{Type: lexer.NORMAL_NUM_LITERAL, Literal: "0.0e0", HasError: false, Pos: util.Position{Len: 5}}}},
		{"0e+0", []lexer.Token{{Type: lexer.NORMAL_NUM_LITERAL, Literal: "0e+0", HasError: false, Pos: util.Position{Len: 4}}}},

		// No data error
		{"0.", []lexer.Token{{Type: lexer.NORMAL_NUM_LITERAL_ERROR, Literal: "0.", HasError: true, Pos: util.Position{Len: 2}}}},
		{"0e", []lexer.Token{{Type: lexer.NORMAL_NUM_LITERAL_ERROR, Literal: "0e", HasError: true, Pos: util.Position{Len: 2}}}},
		{"0e+", []lexer.Token{{Type: lexer.NORMAL_NUM_LITERAL_ERROR, Literal: "0e+", HasError: true, Pos: util.Position{Len: 3}}}},

		// Wrong format error
		{"0a", []lexer.Token{{Type: lexer.NORMAL_NUM_LITERAL_ERROR, Literal: "0a", HasError: true, Pos: util.Position{Len: 2}}}},
		{"0.a", []lexer.Token{{Type: lexer.NORMAL_NUM_LITERAL_ERROR, Literal: "0.a", HasError: true, Pos: util.Position{Len: 3}}}},
		{"0ea", []lexer.Token{{Type: lexer.NORMAL_NUM_LITERAL_ERROR, Literal: "0ea", HasError: true, Pos: util.Position{Len: 3}}}},
		{"0e+a", []lexer.Token{{Type: lexer.NORMAL_NUM_LITERAL_ERROR, Literal: "0e+a", HasError: true, Pos: util.Position{Len: 4}}}},
		{"0.0a", []lexer.Token{{Type: lexer.NORMAL_NUM_LITERAL_ERROR, Literal: "0.0a", HasError: true, Pos: util.Position{Len: 4}}}},
		{"0e0a", []lexer.Token{{Type: lexer.NORMAL_NUM_LITERAL_ERROR, Literal: "0e0a", HasError: true, Pos: util.Position{Len: 4}}}},
	})
}

func TestParseUnknown(t *testing.T) {
	testHelper(t, []testStruct{
		{"§", []lexer.Token{{Type: lexer.UNKNOWN, Literal: "§", HasError: true, Pos: util.Position{Len: 1}}}},
		{"§§", []lexer.Token{{Type: lexer.UNKNOWN, Literal: "§§", HasError: true, Pos: util.Position{Len: 2}}}},
	})
}
