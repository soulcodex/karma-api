package httpserver

import (
	"context"
	"net/http"

	"github.com/soulcodex/karma-api/pkg/utils"
)

type correlationId string

const HeaderRequestIdentifier = "X-Request-Id"
const contextKeyRequestIdentifier correlationId = "correlation_id"

type RequestIdentifierMiddleware struct {
	requestIdProvider utils.UuidProvider
}

func NewRequestIdentifierMiddleware(requestIdProvider utils.UuidProvider) *RequestIdentifierMiddleware {
	return &RequestIdentifierMiddleware{requestIdProvider: requestIdProvider}
}

func (rim *RequestIdentifierMiddleware) Middleware() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := rim.requestIdProvider.New().String()
			r.Header.Set(HeaderRequestIdentifier, requestId)

			newRequest := r.WithContext(context.WithValue(r.Context(), contextKeyRequestIdentifier, requestId))

			w.Header().Set(HeaderRequestIdentifier, requestId)

			next.ServeHTTP(w, newRequest)
		})
	}
}
