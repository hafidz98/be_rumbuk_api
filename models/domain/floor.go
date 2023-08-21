package domain

import "time"

type Floor struct {
	ID         int
	Name       string
	BuildingID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
