// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"
	"github.com/seniorGolang/tg/v2/example/interfaces"
	"github.com/seniorGolang/tg/v2/example/interfaces/types"
)

type UserGetUser func(ctx context.Context, cookie string, userAgent string) (user *types.User, err error)
type UserCustomResponse func(ctx context.Context, arg0 int, arg1 string, opts ...interface{}) (err error)
type UserCustomHandler func(ctx context.Context, arg0 int, arg1 string, opts ...interface{}) (err error)

type MiddlewareUser func(next interfaces.User) interfaces.User

type MiddlewareUserGetUser func(next UserGetUser) UserGetUser
type MiddlewareUserCustomResponse func(next UserCustomResponse) UserCustomResponse
type MiddlewareUserCustomHandler func(next UserCustomHandler) UserCustomHandler
