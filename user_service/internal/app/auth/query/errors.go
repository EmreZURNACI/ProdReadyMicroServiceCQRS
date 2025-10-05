package query

import "errors"

var (
	ErrUuidNotCreated        = errors.New("uuid not created")
	ErrTokenInvalid          = errors.New("authentication failed: invalid token")
	ErrTokenNotProvided      = errors.New("authentication failed: token is missing or could not be retrieved")
	ErrTokenSignatureInvalid = errors.New("authentication failed: token signature does not match")
	ErrTokenClaimsInvalid    = errors.New("authentication failed: token claims are malformed or incomplete")
	ErrTokenExpInvalid       = errors.New("authentication failed: 'exp' field missing or invalid in token")
	ErrTokenExpired          = errors.New("authentication failed: token has expired")
	ErrTokenNickNotFound     = errors.New("authentication failed: user nickname not found in token claims")
	ErrTokenRoleNotFound     = errors.New("authentication failed: user role not found in token claims")
	ErrRefreshTokenNotSaved  = errors.New("authentication failed: refresh token not created in db")
)
