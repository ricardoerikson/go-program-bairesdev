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
	DeleteEndpoint         endpoint.Endpoint
}

func makeAddEndpoint(service service.QuestionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddQuestionRequest)
		question, err := service.Add(ctx, req.Question)
		if err != nil {
			return ErrorResponse{Err: err.Error()}, err
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
		all, err := service.GetByUserID(ctx, req.ID)
		if err != nil {
			return ErrorResponse{Err: err.Error()}, err
		}
		return all, nil
	}
}

func makeUpdateQuestionEndpoint(service service.QuestionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateQuestionRequest)
		err := service.Update(ctx, req.ID, req.Question)
		if err != nil {
			return nil, err
		}
		return req.Question, nil
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
		DeleteEndpoint:         makeDeleteQuestionEndpoint(s),
	}
}
