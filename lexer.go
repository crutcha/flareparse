package flareparse

type Lexer struct {
	input        []byte
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	// TODO: 4096 max length for rule syntax
	l := &Lexer{input: []byte(input)}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return string(l.input[position:l.position])
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	//for isLetter(l.ch) {
	for isAlphaNumeric(l.ch) {
		l.readChar()
	}

	return string(l.input[position:l.position])
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) readStringValue() string {
	// read until we find the closing quote, which could be out of
	// bounds if the expression is invalid
	position := l.position

	// TODO: how to do error handling here?
	// we already know the first character is a double quoted string
	// so we can advance our iterator and look for the second instance
	// of a double quote
	l.readChar()
	for l.ch != 0 {
		if l.ch == '"' {
			l.readChar()
			break
		}

		l.readChar()
	}

	return string(l.input[position:l.position])
}

func (l *Lexer) NextToken() Token {
	var token Token

	l.skipWhitespace()

	switch l.ch {
	case '(':
		token = newToken(LPAREN, l.ch)
	case ')':
		token = newToken(RPAREN, l.ch)
	case '{':
		token = newToken(LBRACE, l.ch)
	case '}':
		token = newToken(RBRACE, l.ch)
	case ',':
		token = newToken(COMMA, l.ch)
	case '[':
		token = newToken(LBRACKET, l.ch)
	case ']':
		token = newToken(RBRACKET, l.ch)
	case '*':
		token = newToken(WILDCARD, l.ch)
	// TODO: other characters requiring peek
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			token = Token{Type: EQUAL, Literal: literal}
		} else {
			token = newToken(ILLEGAL, l.ch)
		}
	case 0:
		token.Literal = ""
		token.Type = EOF
	default:
		// readIdentier() will use position as it's starting index, but l.ch is set
		// to readPosition. this works for IDENT but we need to see if our current
		// cursor at position points to a double quote to handle string value types
		if l.ch == '"' {
			token.Type = VALUE
			token.Literal = l.readStringValue()
			return token
		} else if isAlphaNumeric(l.ch) {
			// IP addresses will be read as 4 numbers separated by a period, which is not
			// what we want, we want the entire IP as a value. for now, we'll treat any
			// numbers we come accross as an identifier and revisit this later
			// this conditional used to be `else if isLetter(l.ch)`
			token.Literal = l.readIdentifier()
			token.Type = LookupKeywordIdent(token.Literal)
			return token
			//} else if isDigit(l.ch) {
			//	token.Type = INT
			//	token.Literal = l.readNumber()
			//	return token
		} else {
			token = newToken(ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return token
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '.'
}

func isAlphaNumeric(ch byte) bool {
	return isLetter(ch) || isDigit(ch)
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
