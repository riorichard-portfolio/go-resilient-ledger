package timer

import (
	"time"
)

type Timer struct {
	timer      *time.Timer
	timeoutDur time.Duration
}

// timeoutMs default 10ms if timeoutMs <= 0 ms
func NewTimer(timeoutMs int) (*Timer) {
	if timeoutMs <= 0 {
		timeoutMs = 10
	}
	dur := time.Duration(timeoutMs) * time.Millisecond
	t := time.NewTimer(dur)
	if !t.Stop() {
		select {
		case <-t.C:
		default:
		}
	}
	return &Timer{
		timer:      t,
		timeoutDur: dur,
	}
}

func (t *Timer) drainTimerC() {
	if !t.timer.Stop() {
		select {
		case <-t.timer.C:
		default:
		}
	}
}

func (t *Timer) Start() {
	t.drainTimerC()
	t.timer.Reset(t.timeoutDur)
}

func (t *Timer) Stop() {
	t.drainTimerC()
}

func (t *Timer) Timeout() <-chan time.Time {
	return t.timer.C
}
