package pg

import (
	"context"

	gopg "github.com/go-pg/pg/v10"
	"questionsandanswers.com/pkg/question/entity"
	"questionsandanswers.com/pkg/question/persistence"
)

type QuestionRepositoryPgImpl struct {
	DB *gopg.DB
}

func (r *QuestionRepositoryPgImpl) NewRepository(args ...interface{}) persistence.QuestionRepository {
	return &QuestionRepositoryPgImpl{DB: args[0].(*gopg.DB)}
}

func (r *QuestionRepositoryPgImpl) Add(c context.Context, q entity.Question) (*entity.Question, error) {
	_, err := r.DB.Model(&q).Insert()
	if err != nil {
		return nil, err
	}
	return &q, nil
}

func (r *QuestionRepositoryPgImpl) GetAll(c context.Context) ([]entity.Question, error) {
	var questions []entity.Question
	err := r.DB.Model(&questions).Select()
	if err != nil {
		return nil, err
	}
	if len(questions) == 0 {
		return make([]entity.Question, 0), nil
	}
	return questions, nil
}

func (r *QuestionRepositoryPgImpl) GetByID(c context.Context, ID int) (*entity.Question, error) {
	var question entity.Question
	err := r.DB.Model(&question).Where("id = ?", ID).Select()
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (r *QuestionRepositoryPgImpl) GetAllByUserID(c context.Context, userID int) ([]entity.Question, error) {
	var questions []entity.Question
	err := r.DB.Model(&questions).Where("user_id = ?", userID).Select()
	if err != nil {
		return nil, err
	}
	if len(questions) == 0 {
		return make([]entity.Question, 0), nil
	}
	return questions, nil
}

func (r *QuestionRepositoryPgImpl) Update(c context.Context, ID int, statement string) (*entity.Question, error) {
	question := entity.Question{Statement: statement}
	_, err := r.DB.Model(&question).Set("statement = ?statement").Where("id = ?", ID).Returning("*").Update(&question)
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (r *QuestionRepositoryPgImpl) UpdateAnswer(c context.Context, ID int, answer string) (*entity.Question, error) {
	question := entity.Question{Answer: answer}
	_, err := r.DB.Model(&question).Set("answer = ?answer").Where("id = ?", ID).Returning("*").Update(&question)
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (r *QuestionRepositoryPgImpl) Delete(c context.Context, ID int) error {
	question := entity.Question{ID: ID}
	_, err := r.DB.Model(&question).Where("id = ?", ID).Delete()
	if err != nil {
		return err
	}
	return nil
}
