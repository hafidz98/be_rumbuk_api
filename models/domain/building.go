package domain

import "time"

type Building struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
