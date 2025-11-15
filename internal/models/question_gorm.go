package models

import (
	"time"
)

type QuestionGORM struct {
	ID        int       `gorm:"primaryKey"`
	Text      string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	Answers []AnswerGORM `gorm:"foreignKey:QuestionID;constraint:OnDelete:CASCADE"`
}

func (QuestionGORM) TableName() string {
	return "questions"
}
