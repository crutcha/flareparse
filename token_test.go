package flareparse

import (
	"fmt"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `(http.host eq "domain.com")`
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{LPAREN, "("},
		{IDENT, "http.host"},
		{EQUAL_KEYWORD, "eq"},
		{VALUE, `"domain.com"`},
		{RPAREN, ")"},
	}

	lexer := NewLexer(input)

	for idx, tokenType := range tests {
		token := lexer.NextToken()

		if token.Type != tokenType.expectedType {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q", idx, tokenType.expectedType, token.Type)
		}

		if token.Literal != tokenType.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", idx, tokenType.expectedLiteral, token.Literal)
		}
	}
}

func TestNextTokenWithBuiltins(t *testing.T) {
	input := `(starts_with(http.host, "subdomain"))`
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{LPAREN, "("},
		{IDENT, "starts_with"},
		{LPAREN, "("},
		{IDENT, "http.host"},
		{COMMA, ","},
		{VALUE, `"subdomain"`},
		{RPAREN, ")"},
		{RPAREN, ")"},
	}

	lexer := NewLexer(input)

	for idx, tokenType := range tests {
		token := lexer.NextToken()

		if token.Type != tokenType.expectedType {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q", idx, tokenType.expectedType, token.Type)
		}

		if token.Literal != tokenType.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", idx, tokenType.expectedLiteral, token.Literal)
		}
	}
}

func TestLexerMultipleConditions(t *testing.T) {
	input := `http.host matches "(www|api)\.example\.org"
and not lower(http.request.uri.path) matches "/(auth|login|logut).*"
and (
  any(http.request.uri.args.names[*] == "token") or
  ip.src in { 93.184.216.34 62.122.170.171 }
)
or cf.threat_score lt 10`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{IDENT, "http.host"},
		{MATCHES_REGEX, "matches"},
		{VALUE, `"(www|api)\.example\.org"`},
		{LOGICAL_AND, "and"},
		{LOGICAL_NOT, "not"},
		{IDENT, "lower"},
		{LPAREN, "("},
		{IDENT, "http.request.uri.path"},
		{RPAREN, ")"},
		{MATCHES_REGEX, "matches"},
		{VALUE, `"/(auth|login|logut).*"`},
		{LOGICAL_AND, "and"},
		{LPAREN, "("},
		{IDENT, "any"},
		{LPAREN, "("},
		{IDENT, "http.request.uri.args.names"},
		{LBRACKET, "["},
		{WILDCARD, "*"},
		{RBRACKET, "]"},
		{EQUAL, "=="},
		{VALUE, `"token"`},
		{RPAREN, ")"},
		{LOGICAL_OR, "or"},
		{IDENT, "ip.src"},
		{VALUE_IN, "in"},
		{LBRACE, "{"},
		{IDENT, "93.184.216.34"},
		{IDENT, "62.122.170.171"},
		{RBRACE, "}"},
		{RPAREN, ")"},
		{LOGICAL_OR, "or"},
		{IDENT, "cf.threat_score"},
		{LESS_THAN_KEYWORD, "lt"},
		{IDENT, "10"},
	}

	lexer := NewLexer(input)

	for idx, tokenType := range tests {
		if idx == 33 {
			fmt.Println("debug time")
		}
		token := lexer.NextToken()

		if token.Type != tokenType.expectedType {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q", idx, tokenType.expectedType, token.Type)
		}

		if token.Literal != tokenType.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", idx, tokenType.expectedLiteral, token.Literal)
		}
	}
}
