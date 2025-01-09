package httpserver

import (
	"github.com/google/jsonapi"
	"net/http"

	xjsonapiresponse "github.com/soulcodex/karma-api/pkg/json-api/response"
	"github.com/soulcodex/karma-api/pkg/logger"
)

type PanicRecoverMiddleware struct {
	logger logger.Logger
}

func NewPanicRecoverMiddleware(l logger.Logger) *PanicRecoverMiddleware {
	return &PanicRecoverMiddleware{logger: l}
}

func (prm *PanicRecoverMiddleware) Middleware() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					prm.logger.Error(r.Context(), "unhandled error")
					response := xjsonapiresponse.NewInternalServerErrorWithDetails("unhandled error")
					w.Header().Set("Content-Type", jsonapi.MediaType)
					w.WriteHeader(http.StatusInternalServerError)

					if marshalErr := jsonapi.MarshalErrors(w, response); marshalErr != nil {
						prm.logger.Error(r.Context(), "Could not write the response of the panic error")
					}
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
