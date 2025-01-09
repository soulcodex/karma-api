package xjsonapiresponse

import (
	"net/http"
	"strconv"

	"github.com/google/jsonapi"

	"github.com/soulcodex/karma-api/pkg/utils"
)

const (
	invalidPayloadReceivedDefaultTitle   = "Invalid Payload"
	invalidPayloadReceivedDefaultCode    = "invalid_payload_received"
	invalidPayloadReceivedDefaultMessage = "Invalid payload received"

	invalidBadRequestTitle = "Bad Request"
	invalidBadRequestCode  = "bad_request"
)

func NewBadRequest(detail string) []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   invalidBadRequestCode,
		Title:  invalidBadRequestTitle,
		Detail: detail,
		Status: strconv.Itoa(http.StatusBadRequest),
	}}
}

func NewBadRequestForInvalidPayloadWithDetails(items ...MetadataItem) []*jsonapi.ErrorObject {
	metadata := NewMetadata(items...).MetadataMap()

	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   invalidPayloadReceivedDefaultCode,
		Title:  invalidPayloadReceivedDefaultTitle,
		Detail: invalidPayloadReceivedDefaultMessage,
		Status: strconv.Itoa(http.StatusBadRequest),
		Meta:   &metadata,
	}}
}

func NewBadRequestForInvalidPayload() []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   invalidPayloadReceivedDefaultCode,
		Title:  invalidPayloadReceivedDefaultTitle,
		Detail: invalidPayloadReceivedDefaultMessage,
		Status: strconv.Itoa(http.StatusBadRequest),
	}}
}

func NewInvalidPayloadCustom(code string, desc string, detail string, meta map[string]interface{}) []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   code,
		Title:  desc,
		Detail: detail,
		Status: strconv.Itoa(http.StatusBadRequest),
		Meta:   &meta,
	}}
}
