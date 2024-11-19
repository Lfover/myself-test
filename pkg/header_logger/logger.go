package header_logger

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
)

func Header() log.Valuer {
	return func(ctx context.Context) interface{} {
		if ctx == nil {
			return ""
		}
		if header, ok := transport.FromServerContext(ctx); ok {
			headers := make(map[string]interface{})
			for _, k := range header.RequestHeader().Keys() {
				if k == "Authorization" {
					continue
				}
				headers[k] = header.RequestHeader().Get(k)
			}
			return headers
		}
		return ""
	}
}
