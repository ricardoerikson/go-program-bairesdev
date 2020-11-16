package entity

type Question struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Statement string `json:"statement"`
	Answer    string `json:"answer,omitempty"`
}
