package lexer

import (
	"ink/token"
)

type Lexer struct {
	input    string // source code
	position int    // current character position under the microscope,
	// think of it like the index

	next_position int  // next position after Position, next index
	ch            byte // the actual ch
}

func (lexer *Lexer) read_ch() {
	if lexer.next_position >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.next_position]
	}

	lexer.position = lexer.next_position
	lexer.next_position += 1
}

// Create a new instance of Lexer
func NewLexer(ch string) *Lexer {
	lexer := &Lexer{
		input: ch,
	}

	lexer.read_ch()
	return lexer
}

func is_letter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (lexer *Lexer) check_identifier() string {
	position := lexer.position
	for is_letter(lexer.ch) {
		lexer.read_ch()
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	switch lexer.ch {
	case '{':
		tok = token.Token{Kind: token.LBRACE, Value: string(lexer.ch)}
	case 0:
		tok.Value = ""
		tok.Kind = token.EOF
	default:
		if is_letter(lexer.ch) {
			tok.Value = lexer.check_identifier()
			tok.Kind = token.CheckKeyword(tok.Value)
			return tok
		} else {
			tok = token.Token{Kind: token.ILLEGAL, Value: string(lexer.ch)}
		}
	}

	lexer.read_ch()
	return tok
}
