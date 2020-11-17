package mem

import (
	"context"

	"gopkg.in/errgo.v2/fmt/errors"
	"questionsandanswers.com/pkg/question/entity"
	"questionsandanswers.com/pkg/question/persistence"
)

type QuestionsDatabase []entity.Question

type QuestionRepositoryMemImpl struct {
	database QuestionsDatabase
}

func (r *QuestionRepositoryMemImpl) NewRepository(args ...interface{}) persistence.QuestionRepository {
	return &QuestionRepositoryMemImpl{database: make(QuestionsDatabase, 0)}
}

func (r *QuestionRepositoryMemImpl) Add(c context.Context, q entity.Question) (*entity.Question, error) {
	r.database = append(r.database, q)
	return &q, nil
}

func (r *QuestionRepositoryMemImpl) UpdateAnswer(c context.Context, ID int, answer string) (*entity.Question, error) {
	for i, question := range r.database {
		if question.ID == ID {
			r.database[i].Answer = answer
			return &r.database[i], nil
		}
	}
	return nil, errors.Newf(`question ID: "%d" not found`, ID)
}

func (r *QuestionRepositoryMemImpl) GetAll(c context.Context) ([]entity.Question, error) {
	return r.database, nil
}

func (r *QuestionRepositoryMemImpl) GetByID(c context.Context, ID int) (*entity.Question, error) {
	for _, question := range r.database {
		if question.ID == ID {
			return &question, nil
		}
	}
	return nil, errors.New("Question not found")
}

func (r *QuestionRepositoryMemImpl) GetAllByUserID(c context.Context, userID int) ([]entity.Question, error) {
	questions := make([]entity.Question, 0)
	for _, question := range r.database {
		if question.UserID == userID {
			questions = append(questions, question)
		}
	}
	return questions, nil

}

func (r *QuestionRepositoryMemImpl) Update(c context.Context, ID int, statement string) (*entity.Question, error) {
	for i, question := range r.database {
		if question.ID == ID {
			r.database[i].Statement = statement
			return &r.database[i], nil
		}
	}
	return nil, errors.Newf("Question with ID: %d does not exist", ID)
}

func (r *QuestionRepositoryMemImpl) Delete(c context.Context, ID int) error {
	for i, question := range r.database {
		if question.ID == ID {
			r.database = append(r.database[:i], r.database[i+1:]...)
			return nil
		}
	}
	return errors.Newf("Question ID: %d not found", ID)
}
