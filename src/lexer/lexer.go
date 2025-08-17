package lexer

import "regexp"

type ragexHandler func(lex *lexer, regex *regexp.Regexp)

type regexPatterns struct {
	regexp  *regexp.Regexp
	handler ragexHandler
}

type lexer struct {
	patterns []regexPatterns
	Tokens   []Token
	source   string
	pos      int
}

func (lex *lexer) advanceN(n int) {
	lex.pos += n
}

func (lex *lexer) push(token Token) {
	lex.Tokens = append(lex.Tokens, token)
}

func Tokenize(source string) []Token {
	lex := creteLexer(source)

	return lex.Tokens
}

func defaultHandler(kind TokenKind, value string) ragexHandler {
	return func(lex *lexer, regex *regexp.Regexp) {
		lex.advanceN(len(value))
		lex.push(NewToken(kind, value))
	}
}

func creteLexer(source string) *lexer {
	return &lexer{
		pos:    0,
		source: source,
		Tokens: make([]Token, 0),
		patterns: []regexPatterns{
			{regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
		},
	}
}
