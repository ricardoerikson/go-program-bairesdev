package mock

import (
	"context"

	smock "github.com/stretchr/testify/mock"
	"questionsandanswers.com/pkg/question/entity"
)

type QuestionServiceMockImpl struct {
	smock.Mock
}

func NewService() *QuestionServiceMockImpl {
	return &QuestionServiceMockImpl{}
}

func (m *QuestionServiceMockImpl) Add(c context.Context, q entity.Question) (*entity.Question, error) {
	args := m.Called(c, q)
	return args.Get(0).(*entity.Question), args.Error(1)
}

func (m *QuestionServiceMockImpl) GetAll(c context.Context) ([]entity.Question, error) {
	args := m.Called(nil)
	return args.Get(0).([]entity.Question), args.Error(1)
}

func (m *QuestionServiceMockImpl) GetByID(c context.Context, ID int) (*entity.Question, error) {
	args := m.Called(nil, ID)
	return args.Get(0).(*entity.Question), args.Error(1)
}

func (m *QuestionServiceMockImpl) GetAllByUserID(c context.Context, userID int) ([]entity.Question, error) {
	args := m.Called(nil, userID)
	return args.Get(0).([]entity.Question), args.Error(1)
}

func (m *QuestionServiceMockImpl) UpdateAnswer(c context.Context, ID int, answer string) (*entity.Question, error) {
	args := m.Called(nil, ID, answer)
	return args.Get(0).(*entity.Question), args.Error(1)
}

func (m *QuestionServiceMockImpl) Update(c context.Context, ID int, statement string) (*entity.Question, error) {
	args := m.Called(nil, ID, statement)
	return args.Get(0).(*entity.Question), args.Error(1)
}

func (m *QuestionServiceMockImpl) Delete(c context.Context, ID int) error {
	args := m.Called(nil, ID)
	return args.Error(0)
}
