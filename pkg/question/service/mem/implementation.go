package mem

import (
	"context"

	"gopkg.in/errgo.v2/fmt/errors"
	"questionsandanswers.com/pkg/question/entity"
)

type QuestionsDatabase []entity.Question

type QuestionServiceMemImpl struct {
	database QuestionsDatabase
}

func NewService() *QuestionServiceMemImpl {
	return &QuestionServiceMemImpl{database: make(QuestionsDatabase, 0)}
}

func (r *QuestionServiceMemImpl) Add(c context.Context, q entity.Question) (*entity.Question, error) {
	q.Answer = ""
	r.database = append(r.database, q)
	return &q, nil
}

func (r *QuestionServiceMemImpl) UpdateAnswer(c context.Context, ID int, answer string) (*entity.Question, error) {
	for i, question := range r.database {
		if question.ID == ID {
			r.database[i].Answer = answer
			return &r.database[i], nil
		}
	}
	return nil, errors.Newf(`question ID: "%d" not found`, ID)
}

func (r *QuestionServiceMemImpl) GetAll(c context.Context) ([]entity.Question, error) {
	return r.database, nil
}

func (r *QuestionServiceMemImpl) GetByID(c context.Context, ID int) (*entity.Question, error) {
	for _, question := range r.database {
		if question.ID == ID {
			return &question, nil
		}
	}
	return nil, errors.New("Question not found")
}

func (r *QuestionServiceMemImpl) GetAllByUserID(c context.Context, userID int) ([]entity.Question, error) {
	questions := make([]entity.Question, 0)
	for _, question := range r.database {
		if question.UserID == userID {
			questions = append(questions, question)
		}
	}
	return questions, nil

}

func (r *QuestionServiceMemImpl) Update(c context.Context, ID int, statement string) (*entity.Question, error) {
	for i, question := range r.database {
		if question.ID == ID {
			r.database[i].Statement = statement
			return &r.database[i], nil
		}
	}
	return nil, errors.Newf("Question with ID: %s does not exist", ID)
}

func (r *QuestionServiceMemImpl) Delete(c context.Context, ID int) error {
	for i, question := range r.database {
		if question.ID == ID {
			r.database = append(r.database[:i], r.database[i+1:]...)
			return nil
		}
	}
	return errors.New("Question not found")
}
