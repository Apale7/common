package constdef

import "errors"

const (
	ErrInvalidParamsMsg = "invalid params"
	ErrInvalidTokenMsg  = "invalid token"
	ErrUnpackFailedMsg  = "unpack failed"
)

var (
	// ErrInvalidParams error for invalid params
	ErrInvalidParams = errors.New(ErrInvalidParamsMsg)
	ErrInvalidToken  = errors.New(ErrInvalidTokenMsg)
	ErrUnpackFailed  = errors.New(ErrUnpackFailedMsg)
)
