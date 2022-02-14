package Event

import (
	"net/http"
	"time"
)

type Event struct {
	CreationTime     time.Time
	ResponseTime     time.Time
	ResponseCode     int
	ResponseBodySize int64
	Request          *http.Request
}

func (e *Event) GetResponseDuration() time.Duration {
	return e.ResponseTime.Sub(e.CreationTime)
}
