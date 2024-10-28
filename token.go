package flareparse

// TODO: list types? $list_name
// TODO: https://github.com/kor44/gofilter
// TODO: look into whitespace and it's implications

const (
	// Identifiers
	IDENT    = "IDENT"
	VALUE    = "VALUE"
	EOF      = "EOF"
	ILLEGAL  = "ILLEGAL"
	INT      = "INT"
	WILDCARD = "*"

	// Operators
	// https://developers.cloudflare.com/ruleset-engine/rules-language/operators/
	EQUAL_KEYWORD              = "eq"
	NOT_EQUAL_KEYWORD          = "ne"
	LESS_THAN_KEYWORD          = "lt"
	LESS_THAN_EQUAL_KEYWORD    = "lte"
	GREATER_THAN_KEYWORD       = "gt"
	GREATER_THAN_EQUAL_KEYWORD = "gte"
	EQUAL                      = "=="
	NOT_EQUAL                  = "!="
	LESS_THAN                  = "<"
	LESS_THAN_EQUAL            = "<="
	GREATER_THAN               = ">"
	GREATER_THAN_EQUAL         = ">="
	EXACTLY_CONTAINS           = "contains"
	MATCHES_REGEX              = "matches"
	VALUE_IN                   = "in"
	LOGICAL_NOT                = "not"
	LOGICAL_AND                = "and"
	LOGICAL_XOR                = "xor"
	LOGICAL_OR                 = "or"

	// Delimiters
	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	COMMA    = ","
	LBRACKET = "["
	RBRACKET = "]"

	// TODO: maybe this is better left for parsing phase?
	// built-in functions
	// https://developers.cloudflare.com/ruleset-engine/rules-language/functions/
	ANY_FUNC         = "any"
	ALL_FUNC         = "all"
	LOWER_FUNC       = "lower"
	STARTS_WITH_FUNC = "starts_with"

	// Keywords
	// TODO: would this include stuff like `ssl`, `http.request.uri.path` etc?
	// TODO: raw string syntax https://developers.cloudflare.com/ruleset-engine/rules-language/values/#raw-string-syntax
)

var keywordsAndBuiltins = map[string]TokenType{
	"eq":       EQUAL_KEYWORD,
	"ne":       NOT_EQUAL_KEYWORD,
	"lt":       LESS_THAN_KEYWORD,
	"le":       LESS_THAN_EQUAL_KEYWORD,
	"gt":       GREATER_THAN_KEYWORD,
	"ge":       GREATER_THAN_EQUAL_KEYWORD,
	"contains": EXACTLY_CONTAINS,
	"matches":  MATCHES_REGEX,
	"in":       VALUE_IN,
	"not":      LOGICAL_NOT,
	"and":      LOGICAL_AND,
	"xor":      LOGICAL_XOR,
	"or":       LOGICAL_OR,

	// TODO: should this be separate?
	// TODO: populate the rest of this map
	//"starts_with": STARTS_WITH_FUNC,
	//"lower":       LOWER_FUNC,
}

type TokenType string

type Token struct {
	Type TokenType
	// TODO: should this by []byte instead?
	Literal string
}

func LookupKeywordIdent(ident string) TokenType {
	if token, ok := keywordsAndBuiltins[ident]; ok {
		return token
	}

	return IDENT
}
