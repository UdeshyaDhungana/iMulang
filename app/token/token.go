package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// types of tokens
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers + literals
	IDFIER = "IDENTIFIER"
	INT    = "INT"
	FLOAT  = "FLOAT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	MOD      = "%"

	LT     = "<"
	GT     = ">"
	LT_EQ  = "<="
	GT_EQ  = ">="
	EQ     = "=="
	NOT_EQ = "!="

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COMMENT   = "$"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// string
	STRING = "STRING"

	// hash
	COLON = ":"

	// array support
	LBRACKET = "["
	RBRACKET = "]"

	// Language Keywords
	THOOS    = "THOOS"
	KAAM_GAR = "KAAM_GAR"

	YEDI   = "YEDI"
	NABHAE = "NABHAE"
	NATRA  = "NATRA"
	PATHA  = "PATHA"

	SACHO = "SACHO"
	JHUT  = "JHUT"

	JABA_SAMMA = "JABA_SAMMA"
	GHUMA      = "GHUMA"
)

func NewToken(t TokenType, r rune) Token {
	return Token{Type: t, Literal: string(r)}
}

func NewTokenFromStr(t TokenType, s string) Token {
	return Token{Type: t, Literal: s}
}

var keywords = map[string]TokenType{
	"thoos":      THOOS,
	"kaam_gar":   KAAM_GAR,
	"yedi":       YEDI,
	"natra":      NATRA,
	"sacho":      SACHO,
	"jhut":       JHUT,
	"patha":      PATHA,
	"jaba_samma": JABA_SAMMA,
	"ghuma":      GHUMA,
	"nabhae":     NABHAE,
}

// this distinguishes reserved keywords from variable names
func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDFIER
}
