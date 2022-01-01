package model

import "time"

type Period struct {
	RecordMeta

	StartAt time.Time `json:"start_at"`
	EndAt time.Time `json:"end_at"`
	Interval int64 `json:"interval"`
}
