package httpserver

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/google/jsonapi"
	"github.com/xeipuuv/gojsonschema"

	xjsonapi "github.com/soulcodex/karma-api/pkg/json-api"
	xjsonapiresponse "github.com/soulcodex/karma-api/pkg/json-api/response"
)

type RequestValidatorMiddleware struct {
	responseMiddleware *xjsonapi.JsonApiResponseMiddleware
	schemaFilePath     string
}

func NewRequestValidatorMiddleware(
	responseMiddleware *xjsonapi.JsonApiResponseMiddleware,
	schemaFilePath string,
) *RequestValidatorMiddleware {
	return &RequestValidatorMiddleware{
		responseMiddleware: responseMiddleware,
		schemaFilePath:     schemaFilePath,
	}
}

func (rvm *RequestValidatorMiddleware) Middleware() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			absPath, _ := filepath.Abs(rvm.schemaFilePath)
			schemaLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%s", absPath))
			bodyBytes, err := io.ReadAll(CloneRequest(r).Body)
			if err != nil {
				validationErrors := xjsonapiresponse.NewBadRequestForInvalidPayload()
				rvm.responseMiddleware.WriteErrorResponse(r.Context(), w, validationErrors, http.StatusBadRequest, err)
				return
			}

			documentLoader := gojsonschema.NewBytesLoader(bodyBytes)

			result, err := gojsonschema.Validate(schemaLoader, documentLoader)
			if err != nil {
				validationErrors := xjsonapiresponse.NewBadRequestForInvalidPayload()
				rvm.responseMiddleware.WriteErrorResponse(r.Context(), w, validationErrors, http.StatusBadRequest, err)
				return
			}

			if !result.Valid() {
				var validationErrors []*jsonapi.ErrorObject
				errors := result.Errors()
				for i := range errors {
					desc := errors[i]
					var details map[string]interface{} = desc.Details()
					validationErrors = xjsonapiresponse.NewInvalidPayloadCustom(desc.Type(), desc.Description(), desc.String(), details)
				}
				rvm.responseMiddleware.WriteErrorResponse(r.Context(), w, validationErrors, http.StatusBadRequest, nil)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
