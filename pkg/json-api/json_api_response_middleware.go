package xjsonapi

import (
	"context"
	"net/http"

	"github.com/google/jsonapi"

	"github.com/soulcodex/karma-api/pkg/logger"
)

type JsonApiResponseMiddleware struct {
	logger logger.Logger
}

func NewJsonApiResponseMiddleware(logger logger.Logger) *JsonApiResponseMiddleware {
	return &JsonApiResponseMiddleware{logger: logger}
}

func (jrm *JsonApiResponseMiddleware) WriteErrorResponse(
	ctx context.Context,
	writer http.ResponseWriter,
	errors []*jsonapi.ErrorObject,
	statusCode int,
	previous error,
) {
	writer.Header().Set("Content-Type", jsonapi.MediaType)
	writer.WriteHeader(statusCode)

	jrm.logError(ctx, previous, statusCode)

	if err := jsonapi.MarshalErrors(writer, errors); err != nil {
		jrm.logger.Error(ctx, "unexpected error marshalling json api response error")
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (jrm *JsonApiResponseMiddleware) WriteResponse(
	ctx context.Context,
	writer http.ResponseWriter,
	payload interface{},
	statusCode int,
) {
	writer.Header().Set("Content-Type", jsonapi.MediaType)
	writer.WriteHeader(statusCode)

	if payload == nil {
		return
	}

	if err := jsonapi.MarshalPayload(writer, payload); err != nil {
		jrm.logger.Error(ctx, "error marshalling json api response", logger.ErrValue("error", err))
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (jrm *JsonApiResponseMiddleware) logError(ctx context.Context, err error, statusCode int) {
	if err == nil {
		return
	}

	if statusCode >= http.StatusInternalServerError {
		jrm.logger.Error(ctx, err.Error(), logger.ErrValue("error", err))
		return
	}

	jrm.logger.Warn(ctx, err.Error(), logger.ErrValue("error", err))
}
