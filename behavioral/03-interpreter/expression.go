package main

// sample regular expression:
// expression = literal | sequence | alternation | repetition | '(' expression ')'
// literal = 'a' | 'b' | 'c' | ... {'a'|'b'|'c'|...}*
// sequence = expression '&' expression
// alternation = expression '|' expression
// repetition = expression '*'

// expression: (('dog'|'cat')* & 'weather')) matches 'dog dog cat weather'

type AbstractExpression interface {
	Interpret(ctx Context)
}

type Context struct {
}

type LiteralExpression struct{}

func (l *LiteralExpression) Interpret(ctx Context) {
}

type SequenceExpression struct{}

func (s *SequenceExpression) Interpret(ctx Context) {}

type AlternationExpression struct{}

func (a *AlternationExpression) Interpret(ctx Context) {}

type RepetitionExpression struct{}

func (r *RepetitionExpression) Interpret(ctx Context) {}
