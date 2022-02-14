package When

import (
	"Packages/src/pkg/Logger/Event"
)

func Pass(activate bool) func(Event *Event.Event) bool {
	if activate {
		return Always
	}
	return Never
}

func Always(Event *Event.Event) bool {
	return true
}

func ResponseBodyBiggerThan(MaxRespBody int64) func(Event *Event.Event) bool {
	return func(Event *Event.Event) bool {
		return Event.ResponseBodySize >= MaxRespBody
	}
}

func ResponseTimeBiggerThan(MaxResponseTime int64) func(Event *Event.Event) bool {
	return func(Event *Event.Event) bool {
		return Event.GetResponseDuration().Milliseconds() > MaxResponseTime
	}
}

func Never(Event *Event.Event) bool {
	return false
}
