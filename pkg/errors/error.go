package errors

import "glox/pkg/tokens"

type ErrorHandler func(int, string)

type TokenErrorHandler func(tokens.Token, error)

type TokenError struct {
	Token tokens.Token
	Err   error
}

func (e *TokenError) Error() string {
	return ""
}
