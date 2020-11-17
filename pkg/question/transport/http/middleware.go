package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"questionsandanswers.com/pkg/question/entity"
	"questionsandanswers.com/pkg/question/persistence"
	"questionsandanswers.com/pkg/question/service"
)

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

type LoggingMiddleware struct {
	Logger log.Logger
	Next   service.QuestionService
}

func (mw LoggingMiddleware) NewService(r persistence.QuestionRepository) service.QuestionService {
	panic("not implemented") // TODO: Implement
}

func (mw LoggingMiddleware) Add(c context.Context, q entity.Question) (response *entity.Question, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "Add",
			"input", fmt.Sprintf("<q: %v>", q),
			"err", err,
			"took", time.Since(begin))
	}(time.Now())
	response, err = mw.Next.Add(c, q)
	return
}

func (mw LoggingMiddleware) GetAll(c context.Context) (response []entity.Question, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "GetAll",
			"err", err,
			"took", time.Since(begin))
	}(time.Now())
	response, err = mw.Next.GetAll(c)
	return
}

func (mw LoggingMiddleware) GetByID(c context.Context, ID int) (response *entity.Question, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "GetByID",
			"input", fmt.Sprintf("<ID: %d>", ID),
			"err", err,
			"took", time.Since(begin))
	}(time.Now())
	response, err = mw.Next.GetByID(c, ID)
	return
}

func (mw LoggingMiddleware) GetAllByUserID(c context.Context, userID int) (response []entity.Question, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "GetAllByUserID",
			"input", fmt.Sprintf("<userID: %d>", userID),
			"err", err,
			"took", time.Since(begin))
	}(time.Now())
	response, err = mw.Next.GetAllByUserID(c, userID)
	return

}

func (mw LoggingMiddleware) Update(c context.Context, ID int, statement string) (response *entity.Question, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "Update",
			"input", fmt.Sprintf(`<ID: %d, statement: "%s">`, ID, statement),
			"output", response,
			"err", err,
			"took", time.Since(begin))
	}(time.Now())
	response, err = mw.Next.Update(c, ID, statement)
	return

}

func (mw LoggingMiddleware) UpdateAnswer(c context.Context, ID int, answer string) (response *entity.Question, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "UpdateAnswer",
			"input", fmt.Sprintf(`<ID: %d, answer: "%s">`, ID, answer),
			"err", err,
			"took", time.Since(begin))
	}(time.Now())
	response, err = mw.Next.UpdateAnswer(c, ID, answer)
	return
}

func (mw LoggingMiddleware) Delete(c context.Context, ID int) (err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "Delete",
			"input", fmt.Sprintf(`<ID: %d>`, ID),
			"err", err,
			"took", time.Since(begin))
	}(time.Now())
	err = mw.Next.Delete(c, ID)
	return
}
