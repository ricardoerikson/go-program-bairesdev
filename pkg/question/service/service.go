package service

import (
	"context"

	"questionsandanswers.com/pkg/question/entity"
	"questionsandanswers.com/pkg/question/persistence"
)

type QuestionService interface {
	Add(c context.Context, q entity.Question) (*entity.Question, error)
	GetAll(c context.Context) ([]entity.Question, error)
	GetByID(c context.Context, ID int) (*entity.Question, error)
	GetByUserID(c context.Context, userID int) ([]entity.Question, error)
	Update(c context.Context, ID int, statement string) (*entity.Question, error)
	UpdateAnswer(c context.Context, ID int, answer string) (*entity.Question, error)
	Delete(c context.Context, ID int) error
}

type QuestionServiceImpl struct {
	Repository persistence.QuestionRepository
}

func (s QuestionServiceImpl) Add(c context.Context, q entity.Question) (*entity.Question, error) {
	return s.Repository.Add(c, q)
}

func (s QuestionServiceImpl) UpdateAnswer(c context.Context, ID int, answer string) (*entity.Question, error) {
	return s.Repository.UpdateAnswer(c, ID, answer)
}

func (s QuestionServiceImpl) GetAll(c context.Context) ([]entity.Question, error) {
	return s.Repository.GetAll(c)
}

func (s QuestionServiceImpl) GetByID(c context.Context, ID int) (*entity.Question, error) {
	return s.Repository.GetByID(c, ID)
}

func (s QuestionServiceImpl) GetByUserID(c context.Context, userID int) ([]entity.Question, error) {
	return s.Repository.GetAllByUserID(c, userID)
}

func (s QuestionServiceImpl) Update(c context.Context, ID int, statement string) (*entity.Question, error) {
	return s.Repository.Update(c, ID, statement)
}

func (s QuestionServiceImpl) Delete(c context.Context, ID int) error {
	return s.Repository.Delete(c, ID)
}
