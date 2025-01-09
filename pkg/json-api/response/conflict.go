package xjsonapiresponse

import (
	"net/http"
	"strconv"

	"github.com/google/jsonapi"

	"github.com/soulcodex/karma-api/pkg/utils"
)

const (
	conflictDefaultTitle = "Conflict"
	conflictDefaultCode  = "conflict"
)

func NewConflict(detail string) []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   conflictDefaultCode,
		Title:  conflictDefaultTitle,
		Detail: detail,
		Status: strconv.Itoa(http.StatusConflict),
	}}
}

func NewConflictWithDetails(detail string, items ...MetadataItem) []*jsonapi.ErrorObject {
	metadata := NewMetadata(items...).MetadataMap()

	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   conflictDefaultCode,
		Title:  conflictDefaultTitle,
		Detail: detail,
		Status: strconv.Itoa(http.StatusConflict),
		Meta:   &metadata,
	}}
}
