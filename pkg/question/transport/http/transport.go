package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"questionsandanswers.com/pkg/question/endpoint"
)

func DecodeAddQuestionRequest(c context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.AddQuestionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeGetAllQuestionsRequest(c context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func DecodeGetAllQuestionsByUserIDRequest(c context.Context, r *http.Request) (interface{}, error) {
	userID, err := strconv.Atoi(mux.Vars(r)["userID"])
	if err != nil {
		return nil, err
	}

	return endpoint.GetAllQuestionsByUserIDRequest{ID: userID}, nil
}

func DecodeUpdateQuestionRequest(c context.Context, r *http.Request) (interface{}, error) {
	questionID, err := strconv.Atoi(mux.Vars(r)["questionID"])
	if err != nil {
		return nil, err
	}
	var updateRequest endpoint.UpdateQuestionRequest
	json.NewDecoder(r.Body).Decode(&updateRequest)
	updateRequest.ID = questionID
	return updateRequest, nil
}

func DecodeDeleteQuestionRequest(c context.Context, r *http.Request) (interface{}, error) {
	questionID, err := strconv.Atoi(mux.Vars(r)["questionID"])
	if err != nil {
		return nil, err
	}
	return endpoint.DeleteQuestionRequest{ID: questionID}, nil
}

func ConfigureRoutes(e endpoint.Endpoints) http.Handler {
	var r = mux.NewRouter()

	r.Methods("POST").Path("/questions").Handler(
		kithttp.NewServer(
			e.AddEndpoint,
			DecodeAddQuestionRequest,
			EncodeResponse))

	r.Methods("GET").Path("/questions").Handler(
		kithttp.NewServer(
			e.GetAllEndpoint,
			DecodeGetAllQuestionsRequest,
			EncodeResponse))

	r.Methods("GET").Path("/users/{userID}/questions").Handler(
		kithttp.NewServer(
			e.GetAllByUserIDEndpoint,
			DecodeGetAllQuestionsByUserIDRequest,
			EncodeResponse))

	r.Methods("PUT").Path("/questions/{questionID}").Handler(
		kithttp.NewServer(
			e.UpdateEndpoint,
			DecodeUpdateQuestionRequest,
			EncodeResponse))

	r.Methods("DELETE").Path("/questions/{questionID}").Handler(
		kithttp.NewServer(
			e.DeleteEndpoint,
			DecodeDeleteQuestionRequest,
			EncodeResponse))

	return r
}
