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
	"questionsandanswers.com/pkg/question/persistence/mock"
	"questionsandanswers.com/pkg/question/service"
	qhttp "questionsandanswers.com/pkg/question/transport/http"
)

type QuestionTestSuite struct {
	suite.Suite
	service service.QuestionService
	mock    *mock.QuestionRepositoryMockImpl
	handler http.Handler
}

func (suite *QuestionTestSuite) SetupSuite() {
	mockRepo := mock.QuestionRepositoryMockImpl{}
	suite.mock = &mockRepo
	suite.service = new(service.QuestionServiceImpl).NewService(&mockRepo)
	endpoints := endpoint.NewEndpoints(suite.service)
	suite.handler = qhttp.NewHTTPTransport(endpoints)
}

func (suite *QuestionTestSuite) SetupTest() {

}

func (suite *QuestionTestSuite) TestAddQuestion() {
	actual := entity.Question{ID: 1, UserID: 1, Statement: "What is 2 + 2?"}
	body, _ := json.Marshal(actual)
	req := httptest.NewRequest("POST", "/questions", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	suite.mock.On("Add", tfymock.Anything, actual).Return(&actual, nil)

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
	suite.mock.AssertExpectations(suite.T())
	assert.Equal(rr.Code, http.StatusOK)
	assert.Equal(actual, expected)
}

func TestStartSuite(t *testing.T) {
	suite.Run(t, new(QuestionTestSuite))
}
