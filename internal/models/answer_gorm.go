package models

import (
	"time"
)

type AnswerGORM struct {
	ID         int       `gorm:"primaryKey"`
	QuestionID int       `gorm:"not null;index"`
	UserID     string    `gorm:"type:uuid;not null"`
	Text       string    `gorm:"type:text;not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

func (AnswerGORM) TableName() string {
	return "answers"
}
