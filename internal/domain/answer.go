package domain

import "time"

type Answer struct {
	ID         int
	QuestionID int
	UserID     string
	Text       string
	CreatedAt  time.Time
}
