package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	tfymock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"questionsandanswers.com/pkg/question/endpoint"
	"questionsandanswers.com/pkg/question/entity"
	"questionsandanswers.com/pkg/question/service/mock"
	qhttp "questionsandanswers.com/pkg/question/transport/http"
)

type QuestionTestSuite struct {
	suite.Suite
	mockService *mock.QuestionServiceMockImpl
	handler     http.Handler
}

func (suite *QuestionTestSuite) SetupSuite() {
	suite.mockService = mock.NewService()
	endpoints := endpoint.NewEndpoints(suite.mockService)
	suite.handler = qhttp.NewHTTPTransport(endpoints)
}

func (suite *QuestionTestSuite) SetupTest() {

}

func (suite *QuestionTestSuite) TestAddQuestion() {
	actual := entity.Question{ID: 1, UserID: 1, Statement: "What is 2 + 2?"}
	body, _ := json.Marshal(actual)
	req := httptest.NewRequest("POST", "/questions", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	suite.mockService.On("Add", tfymock.Anything, actual).Return(&actual, nil)

	suite.handler.ServeHTTP(rr, req)

	// Response
	var expected entity.Question
	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		suite.T().Error("error reading payload")
	}
	json.Unmarshal(body, &expected)

	// Assertions
	assert := assert.New(suite.T())
	suite.mockService.AssertExpectations(suite.T())
	assert.Equal(rr.Code, http.StatusOK)
	assert.Equal(actual, expected)
}

func TestStartSuite(t *testing.T) {
	suite.Run(t, new(QuestionTestSuite))
}
