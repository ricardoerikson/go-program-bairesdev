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

func decodeAddQuestionRequest(c context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.AddQuestionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetAllQuestionsRequest(c context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeGetAllQuestionsByUserIDRequest(c context.Context, r *http.Request) (interface{}, error) {
	userID, err := strconv.Atoi(mux.Vars(r)["userID"])
	if err != nil {
		return nil, err
	}

	return endpoint.GetAllQuestionsByUserIDRequest{ID: userID}, nil
}

func decodeUpdateQuestionRequest(c context.Context, r *http.Request) (interface{}, error) {
	questionID, err := strconv.Atoi(mux.Vars(r)["questionID"])
	if err != nil {
		return nil, err
	}
	var updateRequest endpoint.UpdateQuestionRequest
	json.NewDecoder(r.Body).Decode(&updateRequest)
	updateRequest.ID = questionID
	return updateRequest, nil
}

func decodeDeleteQuestionRequest(c context.Context, r *http.Request) (interface{}, error) {
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
			decodeAddQuestionRequest,
			encodeResponse))

	r.Methods("GET").Path("/questions").Handler(
		kithttp.NewServer(
			e.GetAllEndpoint,
			decodeGetAllQuestionsRequest,
			encodeResponse))

	r.Methods("GET").Path("/users/{userID}/questions").Handler(
		kithttp.NewServer(
			e.GetAllByUserIDEndpoint,
			decodeGetAllQuestionsByUserIDRequest,
			encodeResponse))

	r.Methods("PUT").Path("/questions/{questionID}").Handler(
		kithttp.NewServer(
			e.UpdateEndpoint,
			decodeUpdateQuestionRequest,
			encodeResponse))

	r.Methods("DELETE").Path("/questions/{questionID}").Handler(
		kithttp.NewServer(
			e.DeleteEndpoint,
			decodeDeleteQuestionRequest,
			encodeResponse))

	return r
}
