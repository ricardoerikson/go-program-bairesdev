package service

import (
	"context"

	"questionsandanswers.com/pkg/question/entity"
)

type QuestionService interface {
	Add(c context.Context, q entity.Question) (*entity.Question, error)
	GetAll(c context.Context) ([]entity.Question, error)
	GetByID(c context.Context, ID int) (*entity.Question, error)
	GetAllByUserID(c context.Context, userID int) ([]entity.Question, error)
	Update(c context.Context, ID int, statement string) (*entity.Question, error)
	UpdateAnswer(c context.Context, ID int, answer string) (*entity.Question, error)
	Delete(c context.Context, ID int) error
}
