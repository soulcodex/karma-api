package karma_assignee_http

import (
	"github.com/google/jsonapi"
	"net/http"

	application "github.com/soulcodex/karma-api/internal/karma-assignee/application"
	kadomain "github.com/soulcodex/karma-api/internal/karma-assignee/domain"
	"github.com/soulcodex/karma-api/pkg/bus/command"
	xjsonapi "github.com/soulcodex/karma-api/pkg/json-api"
	xjsonapiresponse "github.com/soulcodex/karma-api/pkg/json-api/response"
)

const UpsertKarmaAssigneeJsonSchemaPath = "upsert-karma-assignee.schema.json"

func HandleUpsertKarmaAssignee(
	bus command.Bus,
	jrm *xjsonapi.JsonApiResponseMiddleware,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bus = nil
		var req UpsertKarmaAssigneeRequest
		if err := jsonapi.UnmarshalPayload(r.Body, &req); err != nil {
			res, statusCode := xjsonapiresponse.NewBadRequest("Invalid received request"), http.StatusBadRequest
			jrm.WriteErrorResponse(r.Context(), w, res, statusCode, err)
			return
		}

		cmd := &application.UpsertKarmaAssigneeCommand{
			Username: req.Username,
			Assigner: req.Assigner,
		}

		err := bus.Dispatch(r.Context(), cmd)

		switch err.(type) {
		case nil:
			jrm.WriteResponse(r.Context(), w, nil, http.StatusNoContent)
			break
		case *kadomain.KarmaAssigneeLimitReached:
			res, statusCode := xjsonapiresponse.NewConflict(err.Error()), http.StatusConflict
			jrm.WriteErrorResponse(r.Context(), w, res, statusCode, err)
			break
		default:
			res, statusCode := xjsonapiresponse.NewInternalServerErrorWithDetails(err.Error()), http.StatusInternalServerError
			jrm.WriteErrorResponse(r.Context(), w, res, statusCode, err)
		}
	}
}

type UpsertKarmaAssigneeRequest struct {
	Username string `jsonapi:"attr,username"`
	Assigner string `jsonapi:"attr,assigner"`
}
