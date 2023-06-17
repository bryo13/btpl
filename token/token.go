package token

// Token definition
type Token struct {
	Kind  string
	Value string
}

var Keywords = map[string]string{
	"defun": "defun", // used to create functions - lisp style
	"if":    "if",
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
	EOF = "EOF"
	NUL = ""
)

func CheckKeyword(val string) string {
	if keyword, ok := Keywords[val]; ok {
		return keyword
	}

	return UserIdentifier
}
