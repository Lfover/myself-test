package dayu_trace

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport"

	"github.com/go-kratos/kratos/v2/log"
)

const xesTraceId = "Traceid"

// TraceID returns openresty dayu_trace_id.
func TraceID() log.Valuer {
	return func(ctx context.Context) interface{} {
		if ctx == nil {
			return ""
		}
		if header, ok := transport.FromServerContext(ctx); ok {
			return header.RequestHeader().Get(xesTraceId)
		}
		return ""
	}
}
