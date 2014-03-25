package util

import (
    "fmt"
    "time"
)

type Timer struct {
    start time.Time
}

func NewTimer() *Timer {
    return &Timer { time.Now() }
}

func (me *Timer) Elapsed() time.Duration {
    return time.Now().Sub(me.start)
}

func TimeIt(msg string, action func()) {
    fmt.Printf("%s ... ", msg)
    t := NewTimer()
    action()
    fmt.Printf("%s\n", t.Elapsed())
}