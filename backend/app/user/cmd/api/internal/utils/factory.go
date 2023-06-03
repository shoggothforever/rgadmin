package utils

import "Table/app/user/cmd/api/internal/types"

func NewResponse(code uint32, msg string) types.Response {
	return types.Response{code, msg}
}
