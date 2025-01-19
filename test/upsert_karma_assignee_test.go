package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"

	kadomain "github.com/soulcodex/karma-api/internal/karma-assignee/domain"
	"github.com/soulcodex/karma-api/pkg/utils"
)

type UpsertKarmaAssigneeSuite struct {
	KarmaAPISuite
}

func (suite *UpsertKarmaAssigneeSuite) SetupSuite() {
	suite.KarmaAPISuite.SetupSuite()
}

func (suite *UpsertKarmaAssigneeSuite) SetupTest() {
	suite.KarmaAPISuite.SetupTest()
}

func TestUpsertKarmaAssigneeSuite(t *testing.T) {
	suite.Run(t, new(UpsertKarmaAssigneeSuite))
}

func (suite *UpsertKarmaAssigneeSuite) TestUpsertKarmaAssigneeInvalidRequestSent() {
	res := suite.ExecuteJsonRequest(
		http.MethodPut,
		"/karma-assignees",
		[]byte("{}"),
		map[string]string{},
	)

	suite.CheckResponseCode(http.StatusBadRequest, res.Code)
}

func (suite *UpsertKarmaAssigneeSuite) TestUpsertKarmaAssigneeRequestCreation() {
	res := suite.ExecuteJsonRequest(
		http.MethodPut,
		"/karma-assignees",
		UpsertKarmaAssigneeBody(),
		map[string]string{},
	)

	suite.CheckResponseCode(http.StatusNoContent, res.Code)

	username, assigner := kadomain.Username("soulcodex"), kadomain.Username("john.doe")
	ka, findErr := suite.KarmaDI.KarmaAssignee.Repository.FindByUsernameAndAssigner(suite.Ctx, username, assigner)
	suite.Require().NoError(findErr)

	suite.Assert().Equal(username.String(), ka.Username().String())
	suite.Assert().Equal(assigner.String(), ka.Assigner().String())
	suite.Assert().Equal(uint64(1), ka.CounterAsNumber())
}

func (suite *UpsertKarmaAssigneeSuite) TestUpsertKarmaAssigneeRequestKarmaIncrease() {
	ulidProvider, timeProvider := utils.NewFixedUlidProvider(), utils.NewFixedTimeProvider()
	username, assigner := kadomain.Username("soulcodex"), kadomain.Username("john.doe")
	existentKarmaAssignee := kadomain.NewKarmaAssigneeFromPrimitives(
		ulidProvider.New().String(),
		username.String(),
		assigner.String(),
		uint64(1),
		timeProvider.Now(),
	)
	saveErr := suite.KarmaDI.KarmaAssignee.Repository.Save(suite.Ctx, existentKarmaAssignee)
	suite.Require().NoError(saveErr)

	res := suite.ExecuteJsonRequest(
		http.MethodPut,
		"/karma-assignees",
		UpsertKarmaAssigneeBody(),
		map[string]string{},
	)

	suite.CheckResponseCode(http.StatusNoContent, res.Code)

	ka, findErr := suite.KarmaDI.KarmaAssignee.Repository.FindByUsernameAndAssigner(suite.Ctx, username, assigner)
	suite.Require().NoError(findErr)

	suite.Assert().Equal(username.String(), ka.Username().String())
	suite.Assert().Equal(assigner.String(), ka.Assigner().String())
	suite.Assert().Equal(uint64(2), ka.CounterAsNumber())
}

func (suite *UpsertKarmaAssigneeSuite) TestUpsertKarmaAssigneeRequestKarmaLimitReached() {
	ulidProvider, timeProvider := utils.NewFixedUlidProvider(), utils.NewFixedTimeProvider()
	username, assigner := kadomain.Username("soulcodex"), kadomain.Username("john.doe")
	existentKarmaAssignee := kadomain.NewKarmaAssigneeFromPrimitives(
		ulidProvider.New().String(),
		username.String(),
		assigner.String(),
		uint64(5),
		timeProvider.Now(),
	)
	saveErr := suite.KarmaDI.KarmaAssignee.Repository.Save(suite.Ctx, existentKarmaAssignee)
	suite.Require().NoError(saveErr)

	res := suite.ExecuteJsonRequest(
		http.MethodPut,
		"/karma-assignees",
		UpsertKarmaAssigneeBody(),
		map[string]string{},
	)

	suite.CheckResponseCode(http.StatusConflict, res.Code)
}

func UpsertKarmaAssigneeBody() []byte {
	return []byte(`
		{
			"data": {
				"type": "karma_assignee",
				"attributes": {
					"username": "soulcodex",
					"assigner": "john.doe"
				}
			}
		}`,
	)
}
