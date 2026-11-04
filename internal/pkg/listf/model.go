package listf

import "time"

type Filter struct {
	From  time.Time
	Until time.Time
	Limit int
}

type FilterJson struct {
	From  time.Time `json:"from" validate:"required"`
	Until time.Time `json:"until" validate:"required"`
	Limit int       `json:"limit" validate:"required"`
}
