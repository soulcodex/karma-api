package xjsonapiresponse

import (
	"net/http"
	"strconv"

	"github.com/google/jsonapi"

	"github.com/soulcodex/karma-api/pkg/utils"
)

const (
	internalServerErrorTitle          = "Internal Server Error"
	internalServerErrorCode           = "internal_server_error"
	internalServerDefaultErrorMessage = "Internal Server Error"
)

func NewInternalServerError() []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   internalServerErrorCode,
		Title:  internalServerErrorTitle,
		Detail: internalServerDefaultErrorMessage,
		Status: strconv.Itoa(http.StatusInternalServerError),
	}}
}

func NewInternalServerErrorWithDetails(detail string, items ...MetadataItem) []*jsonapi.ErrorObject {
	metadata := NewMetadata(items...).MetadataMap()

	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   internalServerErrorCode,
		Title:  internalServerErrorTitle,
		Detail: detail,
		Status: strconv.Itoa(http.StatusInternalServerError),
		Meta:   &metadata,
	}}
}
