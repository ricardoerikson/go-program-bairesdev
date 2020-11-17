package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"questionsandanswers.com/pkg/question/service"
)

type Endpoints struct {
	AddEndpoint            endpoint.Endpoint
	GetAllEndpoint         endpoint.Endpoint
	GetAllByUserIDEndpoint endpoint.Endpoint
	UpdateEndpoint         endpoint.Endpoint
	UpdateAnswerEndpoint   endpoint.Endpoint
	DeleteEndpoint         endpoint.Endpoint
}

func makeAddEndpoint(service service.QuestionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddQuestionRequest)
		question, err := service.Add(ctx, req.Question)
		if err != nil {
			return ErrorResponse{Err: err.Error()}, nil
		}
		return question, nil
	}
}

func makeGetAllEndpoint(service service.QuestionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		all, _ := service.GetAll(ctx)
		return all, nil
	}
}

func makeGetAllQuestionsByUserIDEndpoint(service service.QuestionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllQuestionsByUserIDRequest)
		all, err := service.GetAllByUserID(ctx, req.ID)
		if err != nil {
			return ErrorResponse{Err: err.Error()}, nil
		}
		return all, nil
	}
}

func makeUpdateQuestionEndpoint(service service.QuestionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateQuestionRequest)
		question, err := service.Update(ctx, req.QuestionID, req.Statement)
		if err != nil {
			return nil, err
		}
		return question, nil
	}
}

func makeUpdateAnswerEndpoint(service service.QuestionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateAnswerRequest)
		question, err := service.UpdateAnswer(ctx, req.ID, req.Answer)
		if err != nil {
			return nil, err
		}
		return question, nil
	}
}

func makeDeleteQuestionEndpoint(service service.QuestionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteQuestionRequest)
		err := service.Delete(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}

func MakeEndpoints(s service.QuestionService) Endpoints {
	return Endpoints{
		AddEndpoint:            makeAddEndpoint(s),
		GetAllEndpoint:         makeGetAllEndpoint(s),
		GetAllByUserIDEndpoint: makeGetAllQuestionsByUserIDEndpoint(s),
		UpdateEndpoint:         makeUpdateQuestionEndpoint(s),
		UpdateAnswerEndpoint:   makeUpdateAnswerEndpoint(s),
		DeleteEndpoint:         makeDeleteQuestionEndpoint(s),
	}
}
