package models

import "time"

type Category struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdateAt  time.Time
}
