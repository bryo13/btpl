package lexer

import "ink/token"

type Lexer struct {
	Input        string
	Position     int
	CurrentToken token.Token
	ch           byte
}

// Create a new instance of Lexer
func NewLexer(ch string) *Lexer {
	lexer := &Lexer{
		Input: ch,
	}

	lexer.advance()
	return lexer
}

// advances the lexer to the next character
func (lexer *Lexer) advance() {

	if lexer.Position < len(lexer.Input) {
		lexer.CurrentToken.Value = string(lexer.Input[lexer.Position])

		switch lexer.CurrentToken.Value {
		case "{":
			lexer.CurrentToken.Kind = token.LBRACE
		case "}":
			lexer.CurrentToken.Kind = token.RBRACE
		case "(":
			lexer.CurrentToken.Kind = token.LPAREN
		case ")":
			lexer.CurrentToken.Kind = token.RPAREN
		case ";":
			lexer.CurrentToken.Kind = token.SEMICOLON
		default:
			lexer.CurrentToken.Kind = token.UserIdentifier
		}

		lexer.Position++
	} else {
		lexer.CurrentToken.Kind = token.EOF
		lexer.CurrentToken.Value = token.NUL
	}
}

func (lexer *Lexer) MakeTokens() []token.Token {
	tokens := []token.Token{}

	for lexer.CurrentToken.Kind != "EOF" {
		tokens = append(tokens, lexer.CurrentToken)
		lexer.advance()
	}

	return tokens
}
