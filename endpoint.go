package lark

import (
	"context"
)

type (
	ApiEndpoint   func(ctx context.Context, rawHttpReq *RawRequestReq, resp interface{}) (*Response, error)
	ApiMiddleware func(ApiEndpoint) ApiEndpoint
)

func chainApiMiddleware(mws ...ApiMiddleware) ApiMiddleware {
	return func(next ApiEndpoint) ApiEndpoint {
		for i := len(mws) - 1; i >= 0; i-- {
			next = mws[i](next)
		}
		return next
	}
}
