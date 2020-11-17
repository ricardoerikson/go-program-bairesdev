package mem_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"questionsandanswers.com/pkg/question/entity"
	"questionsandanswers.com/pkg/question/persistence/mem"
	"questionsandanswers.com/pkg/question/service"
)

type RepositoryMemSuite struct {
	suite.Suite
	service service.QuestionService
}

func (s *RepositoryMemSuite) SetupTest() {
	repo := new(mem.QuestionRepositoryMemImpl).NewRepository()
	serv := new(service.QuestionServiceImpl)
	s.service = serv.NewService(repo)
}

// Add Question
func (s *RepositoryMemSuite) TestAddQuestion() {
	question := entity.Question{ID: 1, UserID: 1, Statement: "What is 2+2?"}
	added, err := s.service.Add(nil, question)
	if err != nil {
		s.Fail("Error while adding a question")
	}
	assert := assert.New(s.T())
	assert.Equal(&question, added)
	assert.Empty(added.Answer)
	assert.Equal("What is 2+2?", added.Statement)
}

func (s *RepositoryMemSuite) TestAddQuestionWithEmptyStatement() {
	question := entity.Question{ID: 1, UserID: 1}
	_, err := s.service.Add(nil, question)
	if assert.Error(s.T(), err) {
		assert.EqualError(s.T(), err, "Empty statement is not allowed")
	}
}

func TestFunc(t *testing.T) {
	suite.Run(t, new(RepositoryMemSuite))
}
