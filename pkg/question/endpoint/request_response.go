package endpoint

import (
	"questionsandanswers.com/pkg/question/entity"
)

type ErrorResponse struct {
	Err string `json:"error,omitempty"`
}

type AddQuestionRequest struct {
	entity.Question
}

type GetAllQuestionsByUserIDRequest struct {
	ID int `json:"user_id"`
}

type GetOneQuestionsByIDRequest struct {
	ID int `json:"question_id"`
}

type UpdateQuestionRequest struct {
	QuestionID int    `json:"id"`
	Statement  string `json:"statement,omitempty"`
}

type UpdateAnswerRequest struct {
	ID     int
	Answer string
}

type DeleteQuestionRequest struct {
	ID int `json:"id"`
}
