package With

import (
	"Packages/src/pkg/Logger/Event"
	"fmt"
	"strconv"
)

func ResponseInfoBW(active bool) func(event Event.Event, data map[string]string) {
	if active {
		return ResponseInfo
	}
	return nil
}
func ResponseInfo(event Event.Event, data map[string]string) {
	data["ResponseSize"] = strconv.FormatInt(event.ResponseBodySize, 10)
	data["ResponseTime"] = event.ResponseTime.String()
	data["ResponseCode"] = strconv.Itoa(event.ResponseCode)
	data["ResponseDuration"] = fmt.Sprint(event.ResponseTime.Sub(event.CreationTime))
}

func RequestInfoBW(active bool) func(event Event.Event, data map[string]string) {
	if active {
		return RequestInfo
	}
	return nil
}
func RequestInfo(event Event.Event, data map[string]string) {
	scheme := "http"
	if event.Request.TLS != nil {
		scheme = "https"
	}
	data["RequestTime"] = event.CreationTime.String()
	data["QueryString"] = event.Request.URL.RawQuery
	data["Scheme"] = scheme
	data["Path"] = event.Request.URL.Path
	data["Method"] = event.Request.Method

}
