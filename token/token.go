package token

// Token definition
type Token struct {
	Kind  string
	Value string
}

var Keywords = map[string]string{
	"defun": FUNCTION, // used to create functions - lisp style
	"if":    IF,
	"let":   "let",
}

const (
	LPAREN    = "("
	RPAREN    = ")"
	SEMICOLON = ";"
	RBRACE    = "}"
	LBRACE    = "{"

	// user values
	UserIdentifier = "IDENT"

	// errors or end of file
	EOF     = "EOF"
	NUL     = ""
	ILLEGAL = "ILLEGAL"

	// keywords
	FUNCTION = "FUNCTION"
	IF       = "IF"
)

func CheckKeyword(val string) string {
	if keyword, ok := Keywords[val]; ok {
		return keyword
	}

	return UserIdentifier
}
