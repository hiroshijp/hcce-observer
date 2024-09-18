package domain

import "time"

type History struct {
	ID          int       `json:"id"`
	Visitor     Visitor   `json:"visitor"`
	VisitedFrom string    `json:"visited_from"`
	VisitedAt   time.Time `json:"visited_at"`
}
