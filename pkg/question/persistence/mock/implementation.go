package mock

import (
	"context"

	smock "github.com/stretchr/testify/mock"
	"questionsandanswers.com/pkg/question/entity"
)

type QuestionRepositoryMockImpl struct {
	smock.Mock
}

func (m *QuestionRepositoryMockImpl) Init() {

}

func (m *QuestionRepositoryMockImpl) Add(_ context.Context, q entity.Question) (*entity.Question, error) {
	args := m.Called(nil, q)
	return args.Get(0).(*entity.Question), args.Error(1)
}

func (m *QuestionRepositoryMockImpl) GetAll(_ context.Context) ([]entity.Question, error) {
	args := m.Called(nil)
	return args.Get(0).([]entity.Question), args.Error(1)
}

func (m *QuestionRepositoryMockImpl) GetByID(_ context.Context, ID int) (*entity.Question, error) {
	args := m.Called(nil, ID)
	return args.Get(0).(*entity.Question), args.Error(1)
}

func (m *QuestionRepositoryMockImpl) GetAllByUserID(_ context.Context, userID int) ([]entity.Question, error) {
	args := m.Called(nil, userID)
	return args.Get(0).([]entity.Question), args.Error(1)
}

func (m *QuestionRepositoryMockImpl) Update(_ context.Context, ID int, q entity.Question) error {
	args := m.Called(nil, ID, q)
	return args.Error(0)
}

func (m *QuestionRepositoryMockImpl) Delete(_ context.Context, ID int) error {
	args := m.Called(nil, ID)
	return args.Error(0)
}
