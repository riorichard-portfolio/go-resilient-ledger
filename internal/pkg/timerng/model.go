package timerng

import "time"

type Filter struct {
	From  time.Time
	Until time.Time
}

type FilterJson struct {
	From  time.Time `json:"from" validate:"required"`
	Until time.Time `json:"until" validate:"required"`
}
