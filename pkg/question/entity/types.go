package entity

type Question struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Question string `json:"question"`
	// Answer   answer.Answer `json:"answer"`
}
