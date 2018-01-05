package main

import (
	"time"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"
)




type loggingMiddleware struct {
	logger log.Logger
	next   <%= serviceCamelCase %>Interface
}

// NewLoggingMiddleware - Creates new loggingMiddleware
func NewLoggingMiddleware(s <%= serviceCamelCase %>Interface, logger log.Logger) <%= serviceCamelCase %>Interface {
	return loggingMiddleware{logger, s}
}


<% for(endpoint of endpoints) { %>

func (m loggingMiddleware) <%= endpoint.methodName %>(ctx context.Context, req <%= endpoint.methodName %>Request) (<%= endpoint.methodName %>Response, error) {
	defer func(start time.Time) {
		m.logger.Log(
			"Method", "<%= endpoint.methodName %>",
			"RequestTime", time.Since(start),
		)
	}(time.Now())
	return m.next.<%= endpoint.methodName %>(ctx, req)
}

<% } %>
