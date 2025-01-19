package test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/soulcodex/karma-api/cmd/di"
	"github.com/soulcodex/karma-api/test/arrangers"
)

type KarmaAPISuite struct {
	suite.Suite

	Ctx     context.Context
	KarmaDI *di.KarmaDi

	SuiteArranger *arrangers.KarmaAPISuiteArranger

	routerHandler http.Handler
}

func (suite *KarmaAPISuite) SetupSuite() {
	suite.Ctx = context.Background()
	suite.KarmaDI = di.InitKarmaDIWithEnvFiles(suite.Ctx, "../.env.example", "../.env.test")
	suite.SuiteArranger = arrangers.NewKarmaAPISuiteArranger(suite.KarmaDI.Common)
	suite.routerHandler = nil
}

func (suite *KarmaAPISuite) SetupTest() {
	suite.routerHandler = nil

	suite.KarmaDI.Common.RedisClient.FlushAll(suite.Ctx)
	suite.SuiteArranger.MustArrange(suite.Ctx)
}

func (suite *KarmaAPISuite) TearDownSuite() {
	suite.routerHandler = nil
}

func (suite *KarmaAPISuite) TearDownTest() {
	suite.routerHandler = nil
}

func (suite *KarmaAPISuite) ExecuteJsonRequest(
	verb string,
	path string,
	body []byte,
	headers map[string]string,
) *httptest.ResponseRecorder {
	req, err := http.NewRequest(verb, path, bytes.NewBuffer(body))
	assert.NoError(suite.T(), err)

	if len(headers) != 0 {
		for headerName, value := range headers {
			req.Header.Set(headerName, value)
		}
	}

	req.Header.Set("Content-Type", "application/json")
	return suite.ExecuteRequest(req)
}

func (suite *KarmaAPISuite) ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	resRec := httptest.NewRecorder()

	if suite.routerHandler == nil {
		suite.routerHandler = suite.KarmaDI.Common.Router.GetMuxRouter()
	}

	suite.routerHandler.ServeHTTP(resRec, req)

	return resRec
}

func (suite *KarmaAPISuite) CheckResponse(
	expectedStatusCode int,
	expectedResponse string,
	response *httptest.ResponseRecorder,
	formats ...interface{},
) {
	ja := jsonassert.New(suite.T())
	suite.CheckResponseCode(expectedStatusCode, response.Code)

	receivedResponse := response.Body.String()
	if receivedResponse == "" {
		assert.Equal(suite.T(), expectedResponse, receivedResponse)
		return
	}
	if formats != nil {
		ja.Assertf(receivedResponse, expectedResponse, formats)
	} else {
		ja.Assertf(receivedResponse, expectedResponse)
	}
}

func (suite *KarmaAPISuite) CheckResponseCode(expected, actual int) {
	if expected != actual {
		suite.T().Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
