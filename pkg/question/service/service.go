package service

import (
	"context"
	"errors"

	"questionsandanswers.com/pkg/question/entity"
	"questionsandanswers.com/pkg/question/persistence"
)

type QuestionService interface {
	NewService(r persistence.QuestionRepository) QuestionService
	Add(c context.Context, q entity.Question) (*entity.Question, error)
	GetAll(c context.Context) ([]entity.Question, error)
	GetByID(c context.Context, ID int) (*entity.Question, error)
	GetAllByUserID(c context.Context, userID int) ([]entity.Question, error)
	Update(c context.Context, ID int, statement string) (*entity.Question, error)
	UpdateAnswer(c context.Context, ID int, answer string) (*entity.Question, error)
	Delete(c context.Context, ID int) error
}

type QuestionServiceImpl struct {
	Repository persistence.QuestionRepository
}

func (impl *QuestionServiceImpl) NewService(r persistence.QuestionRepository) QuestionService {
	return &QuestionServiceImpl{Repository: r}
}

func (impl *QuestionServiceImpl) Add(c context.Context, q entity.Question) (*entity.Question, error) {
	if len(q.Statement) == 0 {
		return nil, errors.New("Empty statement is not allowed")
	}
	return impl.Repository.Add(c, q)
}

func (impl *QuestionServiceImpl) GetAll(c context.Context) ([]entity.Question, error) {
	return impl.Repository.GetAll(c)
}

func (impl *QuestionServiceImpl) GetByID(c context.Context, ID int) (*entity.Question, error) {
	return impl.Repository.GetByID(c, ID)
}

func (impl *QuestionServiceImpl) GetAllByUserID(c context.Context, userID int) ([]entity.Question, error) {
	return impl.Repository.GetAllByUserID(c, userID)
}

func (impl *QuestionServiceImpl) Update(c context.Context, ID int, statement string) (*entity.Question, error) {
	if len(statement) == 0 {
		return nil, errors.New("Empty statement is not allowed")
	}
	return impl.Repository.Update(c, ID, statement)
}

func (impl *QuestionServiceImpl) UpdateAnswer(c context.Context, ID int, answer string) (*entity.Question, error) {
	if len(answer) == 0 {
		return nil, errors.New("Empty answer is not allowed")
	}
	return impl.Repository.UpdateAnswer(c, ID, answer)
}

func (impl *QuestionServiceImpl) Delete(c context.Context, ID int) error {
	return impl.Repository.Delete(c, ID)
}
