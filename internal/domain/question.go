package domain

import (
	"time"
)

type Question struct {
	ID        int
	Text      string
	CreatedAt time.Time
}
