package mem_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"questionsandanswers.com/pkg/question/entity"
	"questionsandanswers.com/pkg/question/persistence/mem"
	"questionsandanswers.com/pkg/question/service"
)

type QuestionServiceWithMemDBSuite struct {
	suite.Suite
	service service.QuestionService
}

func (s *QuestionServiceWithMemDBSuite) SetupTest() {
	repo := new(mem.QuestionRepositoryMemImpl).NewRepository()
	serv := new(service.QuestionServiceImpl)
	s.service = serv.NewService(repo)
}

// Add Question
func (s *QuestionServiceWithMemDBSuite) TestAddQuestion() {
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

func (s *QuestionServiceWithMemDBSuite) TestAddQuestionWithEmptyStatement() {
	question := entity.Question{ID: 1, UserID: 1}
	_, err := s.service.Add(nil, question)
	if assert.Error(s.T(), err) {
		assert.EqualError(s.T(), err, "Empty statement is not allowed")
	}
}

func (s *QuestionServiceWithMemDBSuite) TestGetAQuestion() {
	entity := entity.Question{ID: 1, UserID: 1, Statement: "What is 2+2?"}
	s.service.Add(nil, entity)
	q, err := s.service.GetByID(nil, 1)

	assert := assert.New(s.T())
	assert.Nil(err)
	assert.Equal(&entity, q)
}

func (s *QuestionServiceWithMemDBSuite) TestGetAllQuestions() {
	s.service.Add(nil, entity.Question{ID: 1, UserID: 1, Statement: "What is 3+5?"})
	s.service.Add(nil, entity.Question{ID: 2, UserID: 1, Statement: "What is 5+5?"})
	s.service.Add(nil, entity.Question{ID: 2, UserID: 1, Statement: "What is 4*5?"})
	questions, err := s.service.GetAll(nil)
	assert := assert.New(s.T())
	assert.Nil(err)
	assert.Len(questions, 3)
}

func (s *QuestionServiceWithMemDBSuite) TestGetAllWithEmptyDatabase() {
	questions, err := s.service.GetAll(nil)
	assert := assert.New(s.T())
	assert.Nil(err)
	assert.Len(questions, 0)
}

func (s *QuestionServiceWithMemDBSuite) TestGetAllByUserID() {
	s.service.Add(nil, entity.Question{ID: 1, UserID: 1, Statement: "What is 3+5?"})
	s.service.Add(nil, entity.Question{ID: 2, UserID: 2, Statement: "What is 5+5?"})
	s.service.Add(nil, entity.Question{ID: 2, UserID: 1, Statement: "What is 4*5?"})
	questions, err := s.service.GetAllByUserID(nil, 1)
	assert := assert.New(s.T())
	assert.Nil(err)
	assert.Len(questions, 2)
}

func (s *QuestionServiceWithMemDBSuite) TestUpdateQuestion() {
	s.service.Add(nil, entity.Question{ID: 1, UserID: 1, Statement: "What is 3+5?"})
	assert := assert.New(s.T())

	q, err := s.service.GetByID(nil, 1)
	assert.Nil(err)
	assert.Equal("What is 3+5?", q.Statement)
	s.service.Update(nil, 1, "What is 4 * 2?")
	q, err = s.service.GetByID(nil, 1)
	assert.Nil(err)
	assert.Equal("What is 4 * 2?", q.Statement)
}

func (s *QuestionServiceWithMemDBSuite) TestUpdateQuestionWithEmptyStatement() {
	s.service.Add(nil, entity.Question{ID: 1, UserID: 1, Statement: "What is 3+5?"})
	assert := assert.New(s.T())
	q, err := s.service.GetByID(nil, 1)
	assert.Nil(err)
	assert.Equal("What is 3+5?", q.Statement)

	q, err = s.service.Update(nil, 1, "")
	if assert.Error(err) {
		assert.EqualError(err, "Empty statement is not allowed")
	}
	assert.Nil(q)
}

func (s *QuestionServiceWithMemDBSuite) TestUpdateAnswer() {
	s.service.Add(nil, entity.Question{ID: 1, UserID: 1, Statement: "What is 3+5?"})
	_, err := s.service.UpdateAnswer(nil, 1, "The answer is 8")

	assert := assert.New(s.T())
	assert.Nil(err)
	q, err := s.service.GetByID(nil, 1)
	assert.Nil(err)
	assert.Equal("The answer is 8", q.Answer)
}

func (s *QuestionServiceWithMemDBSuite) TestUpdateAnswerWithAnEmptyAnswer() {
	s.service.Add(nil, entity.Question{ID: 1, UserID: 1, Statement: "What is 3+5?"})
	_, err := s.service.UpdateAnswer(nil, 1, "")
	assert := assert.New(s.T())
	if assert.Error(err) {
		assert.EqualError(err, "Empty answer is not allowed")
	}
}

func (s *QuestionServiceWithMemDBSuite) TestDeleteAnswer() {
	assert := assert.New(s.T())
	questions, err := s.service.GetAll(nil)
	assert.Nil(err)
	assert.Empty(questions)
	s.service.Add(nil, entity.Question{ID: 1, UserID: 1, Statement: "What is 3+5?"})
	questions, err = s.service.GetAll(nil)
	assert.Nil(err)
	assert.Len(questions, 1)
	err = s.service.Delete(nil, 1)
	assert.Nil(err)
	questions, err = s.service.GetAll(nil)
	assert.Nil(err)
	assert.Empty(questions)
}

func (s *QuestionServiceWithMemDBSuite) TestDeleteQuestionWithNonExistingID() {
	assert := assert.New(s.T())
	questions, err := s.service.GetAll(nil)
	assert.Nil(err)
	assert.Empty(questions)
	s.service.Add(nil, entity.Question{ID: 1, UserID: 1, Statement: "What is 3+5?"})
	questions, err = s.service.GetAll(nil)
	assert.Nil(err)
	assert.Len(questions, 1)
	err = s.service.Delete(nil, 2)
	if assert.Error(err) {
		assert.EqualError(err, "Question ID: 2 not found")
	}
	questions, err = s.service.GetAll(nil)
	assert.Nil(err)
	assert.Len(questions, 1)
}

func TestStartSuite(t *testing.T) {
	suite.Run(t, new(QuestionServiceWithMemDBSuite))
}
