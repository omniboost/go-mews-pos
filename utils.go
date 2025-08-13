package go_mews_pos

import "time"

type Time struct {
	time.Time
}

func (t Time) String() string {
	return t.Format("2006-01-02T15:04:05+00:00")
}
