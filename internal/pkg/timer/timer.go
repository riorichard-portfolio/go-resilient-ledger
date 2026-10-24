package timer

import (
	"time"
)

type Timer struct {
	timer      *time.Timer
	c          <-chan time.Time
	timeoutDur time.Duration
}

func NewTimer(timeoutMs int) (*Timer, error) {
	if timeoutMs <= 0 {
		return nil, InvalidTimeoutErr
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
		c:          t.C,
		timeoutDur: dur,
	}, nil
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
	return t.c
}
