package goroovy

import (
	"bufio"
	"io"
	"unicode"
)

type Token int

const (
	EOF = iota
	ILLEGAL
	IDENT
	IF
	RETURN
	BRACKETOPEN
	BRACKETCLOSE
	INT
	SEMI // ;
	QUOTE
	DOT
	NEWLINE

	// Infix ops
	ADD // +
	SUB // -
	MUL // *
	DIV // /

	ASSIGN // =
	EQUALS
	MORE
	LESS
	MOREOREQUAL
	LESSOREQUAL

	AND
	OR
)

var tokens = []string{
	EOF:          "EOF",
	ILLEGAL:      "ILLEGAL",
	IDENT:        "IDENT",
	IF:           "IF",
	RETURN:       "RETURN",
	BRACKETOPEN:  "BRACKETOPEN",
	BRACKETCLOSE: "BRACKETCLOSE",
	INT:          "INT",
	SEMI:         "SEMI",
	QUOTE:        "QUOTE",
	DOT:          "DOT",
	NEWLINE:      "NEWLINE",

	// Infix ops
	ADD: "ADD",
	SUB: "SUB",
	MUL: "MUL",
	DIV: "DIV",

	ASSIGN:      "ASSIGN",
	EQUALS:      "EQUALS",
	MORE:        "MORE",
	LESS:        "LESS",
	MOREOREQUAL: "MOREOREQUAL",
	LESSOREQUAL: "LESSOREQUAL",

	AND: "AND",
	OR:  "OR",
}

func (t Token) String() string {
	return tokens[t]
}

type Position struct {
	Line   int
	Column int
}

type Lexer struct {
	pos    Position
	reader *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    Position{Line: 1, Column: 0},
		reader: bufio.NewReader(reader),
	}
}

// Lex scans the input for the next token. It returns the position of the token,
// the token's type, and the literal value.
func (l *Lexer) Lex() (Position, Token, string) {
	// keep looping until we return a token
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, ""
			}

			// at this point there isn't much we can do, and the compiler
			// should just return the raw error to the user
			panic(err)
		}

		// update the column to the position of the newly read in rune
		l.pos.Column++

		switch r {
		case '\n':
			startPos := l.pos
			l.resetPosition()
			return startPos, NEWLINE, "\\n"
		case ';':
			return l.pos, SEMI, ";"
		case '(':
			return l.pos, BRACKETOPEN, "("
		case ')':
			return l.pos, BRACKETCLOSE, ")"
		case '=':
			if l.advance() == '=' {
				return l.pos, EQUALS, "=="
			}
			l.backup()
			return l.pos, EQUALS, "="
		case '>':
			if l.advance() == '=' {
				return l.pos, MOREOREQUAL, ">="
			}
			l.backup()
			return l.pos, MORE, ">"
		case '<':
			if l.advance() == '=' {
				return l.pos, LESSOREQUAL, "<="
			}
			l.backup()
			return l.pos, LESS, "<"
		case '"':
			startPos := l.pos
			quote := l.lexQuote()
			return startPos, QUOTE, quote
		case '.':
			return l.pos, DOT, "."
		case '&':
			if l.advance() == '&' {
				return l.pos, AND, "&&"
			}
			l.backup()
			l.backup()
			fallthrough
		case '|':
			if l.advance() == '|' {
				return l.pos, OR, "||"
			}
			l.backup()
			l.backup()
			fallthrough
		default:
			if unicode.IsSpace(r) {
				continue // nothing to do here, just move on
			} else if unicode.IsDigit(r) {
				// backup and let lexInt rescan the beginning of the int
				startPos := l.pos
				l.backup()
				lit := l.lexInt()
				return startPos, INT, lit
			} else if unicode.IsLetter(r) {
				// backup and let lexIdent rescan the beginning of the ident
				startPos := l.pos
				l.backup()
				lit := l.lexIdent()
				return startPos, tokenType(lit), lit
			} else {
				return l.pos, ILLEGAL, string(r)
			}
		}
	}
}

func (l *Lexer) resetPosition() {
	l.pos.Line++
	l.pos.Column = 0
}

func (l *Lexer) advance() rune {
	r, _, err := l.reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			// at the end of the int
			return r
		}
	}

	l.pos.Column++
	return r
}

func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}

	l.pos.Column--
}

// lexInt scans the input until the end of an integer and then returns the
// literal.
func (l *Lexer) lexInt() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// at the end of the int
				return lit
			}
		}

		l.pos.Column++
		if unicode.IsDigit(r) {
			lit = lit + string(r)
		} else {
			// scanned something not in the integer
			l.backup()
			return lit
		}
	}
}

// lexIdent scans the input until the end of an identifier and then returns the
// literal.
func (l *Lexer) lexIdent() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// at the end of the identifier
				return lit
			}
		}

		l.pos.Column++
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			lit = lit + string(r)
		} else {
			// scanned something not in the identifier
			l.backup()
			return lit
		}
	}
}

func (l *Lexer) lexQuote() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}

		l.pos.Column++
		if r != '"' {
			lit = lit + string(r)
		} else {
			// scanned quote
			return lit
		}
	}
}

func tokenType(lit string) Token {
	switch lit {
	case "if":
		return IF
	case "return":
		return RETURN
	case "(":
		return BRACKETOPEN
	case ")":
		return BRACKETCLOSE
	case "==":
		return EQUALS
	case "=":
		return ASSIGN
	case "\"":
		return QUOTE
	default:
		return IDENT
	}
}
