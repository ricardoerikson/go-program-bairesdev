package entity

import "fmt"

type Question struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Statement string `json:"statement"`
	Answer    string `json:"answer,omitempty"`
}

func (s Question) String() string {
	return fmt.Sprintf(`[ID: %d, UserID: %d, Statement: "%s", Answer: "%s"]`, s.ID, s.UserID, s.Statement, s.Answer)
}
