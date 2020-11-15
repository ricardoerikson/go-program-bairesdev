package persistence

import (
	"context"

	"questionsandanswers.com/pkg/question/entity"
)

type QuestionRepository interface {
	Add(c context.Context, q entity.Question) (*entity.Question, error)
	GetAll(c context.Context) ([]entity.Question, error)
	GetByID(c context.Context, ID int) (*entity.Question, error)
	GetAllByUserID(c context.Context, userID int) ([]entity.Question, error)
	Update(c context.Context, ID int, q entity.Question) error
	Delete(c context.Context, ID int) error
}
