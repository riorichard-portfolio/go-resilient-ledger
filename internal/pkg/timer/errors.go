package timer

import "github.com/cockroachdb/errors"

var (
	InvalidTimeoutErr = errors.New("timer timout must greater than 0")
)