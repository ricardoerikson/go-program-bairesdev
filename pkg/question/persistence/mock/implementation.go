package mock

import (
	"context"

	smock "github.com/stretchr/testify/mock"
	"questionsandanswers.com/pkg/question/entity"
	"questionsandanswers.com/pkg/question/persistence"
)

type QuestionRepositoryMockImpl struct {
	smock.Mock
}

func (m *QuestionRepositoryMockImpl) NewRepository(args ...interface{}) persistence.QuestionRepository {
	return &QuestionRepositoryMockImpl{}
}

func (m *QuestionRepositoryMockImpl) Add(c context.Context, q entity.Question) (*entity.Question, error) {
	args := m.Called(c, q)
	return args.Get(0).(*entity.Question), args.Error(1)
}

func (m *QuestionRepositoryMockImpl) GetAll(c context.Context) ([]entity.Question, error) {
	args := m.Called(nil)
	return args.Get(0).([]entity.Question), args.Error(1)
}

func (m *QuestionRepositoryMockImpl) GetByID(c context.Context, ID int) (*entity.Question, error) {
	args := m.Called(nil, ID)
	return args.Get(0).(*entity.Question), args.Error(1)
}

func (m *QuestionRepositoryMockImpl) GetAllByUserID(c context.Context, userID int) ([]entity.Question, error) {
	args := m.Called(nil, userID)
	return args.Get(0).([]entity.Question), args.Error(1)
}

func (m *QuestionRepositoryMockImpl) UpdateAnswer(c context.Context, ID int, answer string) (*entity.Question, error) {
	args := m.Called(nil, ID, answer)
	return args.Get(0).(*entity.Question), args.Error(1)
}

func (m *QuestionRepositoryMockImpl) Update(c context.Context, ID int, statement string) (*entity.Question, error) {
	args := m.Called(nil, ID, statement)
	return args.Get(0).(*entity.Question), args.Error(1)
}

func (m *QuestionRepositoryMockImpl) Delete(c context.Context, ID int) error {
	args := m.Called(nil, ID)
	return args.Error(0)
}
