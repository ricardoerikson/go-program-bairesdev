package pg

import (
	"context"

	"questionsandanswers.com/pkg/question/entity"
)

type QuestionRepositoryPGImpl struct {
}

func (r *QuestionRepositoryPGImpl) Init() {
	panic("not implemented") // TODO: Implement
}

func (r *QuestionRepositoryPGImpl) Add(c context.Context, q entity.Question) (*entity.Question, error) {
	panic("not implemented") // TODO: Implement
}

func (r *QuestionRepositoryPGImpl) GetAll(c context.Context) ([]entity.Question, error) {
	panic("not implemented") // TODO: Implement
}

func (r *QuestionRepositoryPGImpl) GetByID(c context.Context, ID int) (*entity.Question, error) {
	panic("not implemented") // TODO: Implement
}

func (r *QuestionRepositoryPGImpl) GetAllByUserID(c context.Context, userID int) ([]entity.Question, error) {
	panic("not implemented") // TODO: Implement
}

func (r *QuestionRepositoryPGImpl) Update(c context.Context, ID int, statement string) (*entity.Question, error) {
	panic("not implemented") // TODO: Implement
}

func (r *QuestionRepositoryPGImpl) UpdateAnswer(c context.Context, ID int, answer string) (*entity.Question, error) {
	panic("not implemented") // TODO: Implement
}

func (r *QuestionRepositoryPGImpl) Delete(c context.Context, ID int) error {
	panic("not implemented") // TODO: Implement
}
