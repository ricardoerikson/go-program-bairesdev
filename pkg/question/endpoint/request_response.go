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

type GetAllRequest struct {
}

type GetAllQuestionsByUserIDRequest struct {
	ID int `json:"user_id"`
}

type UpdateQuestionRequest struct {
	entity.Question
}

type UpdateAnswerRequest struct {
	ID     int
	Answer string
}

type UpdateAnswerResponse struct {
	entity.Question
}

type DeleteQuestionRequest struct {
	ID int `json:"id"`
}
