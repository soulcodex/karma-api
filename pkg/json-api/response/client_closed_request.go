package xjsonapiresponse

import (
	"strconv"

	"github.com/google/jsonapi"

	"github.com/soulcodex/karma-api/pkg/utils"
)

const (
	clientClosedRequestHTTPCode = 499
	clientClosedRequestCode     = "client_closed_request"
	clientClosedRequestTitle    = "Client Closed Request"
)

func NewClientClosedRequest(detail string) []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   clientClosedRequestCode,
		Title:  clientClosedRequestTitle,
		Detail: detail,
		Status: strconv.Itoa(clientClosedRequestHTTPCode),
	}}
}

func NewClientClosedRequestWithDetails(detail string, items ...MetadataItem) []*jsonapi.ErrorObject {
	metadata := NewMetadata(items...).MetadataMap()

	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   clientClosedRequestCode,
		Title:  clientClosedRequestTitle,
		Detail: detail,
		Status: strconv.Itoa(clientClosedRequestHTTPCode),
		Meta:   &metadata,
	}}
}
