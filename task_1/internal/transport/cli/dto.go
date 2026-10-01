package cli

import (
	"time"
)

type TaskDTO struct {
	ID        int64
	Title     string
	Status    string
	Deadline  time.Time
	CreatedAt time.Time
}
