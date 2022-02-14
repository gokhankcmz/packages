package Logger

import (
	"Packages/src/pkg/KVP"
	"Packages/src/pkg/Logger/Event"
	"net/http"
	"time"
)

type LogEntry struct {
	Data           map[string]string
	LoggerSettings *Logger
	Event          Event.Event
}

func newEntry(ls *Logger) *LogEntry {
	return &LogEntry{
		Data:           map[string]string{},
		LoggerSettings: ls,
		Event: Event.Event{
			CreationTime: time.Now(),
		},
	}
}

func (le *LogEntry) InformRequest(c *http.Request) *LogEntry {
	le.Event.Request = c
	return le
}

func (le *LogEntry) InformResponse(ResponseCode int, RespBodySizeByte int64) *LogEntry {
	le.Event.ResponseCode = ResponseCode
	le.Event.ResponseBodySize = RespBodySizeByte
	le.Event.ResponseTime = time.Now()
	return le
}
func (le *LogEntry) AddField(fieldName, fieldValue string) *LogEntry {
	le.Data[fieldName] = fieldValue
	return le
}

func (le *LogEntry) AddStruct(s interface{}) *LogEntry {
	kvPairs := KVP.GetKVPs(s, "", "", map[string]string{})
	for k, v := range kvPairs {
		le.AddField(k, v)
	}
	return le
}

func (le *LogEntry) printBasics() *LogEntry {
	le.Data["ApplicationName"] = le.LoggerSettings.ApplicationName
	le.Data["HostName"] = le.LoggerSettings.HostName
	return le
}
func (le *LogEntry) printEnrichments() *LogEntry {
	for _, v := range le.LoggerSettings.Enrichments {
		v(le.Event, le.Data)
	}
	return le
}

func (le *LogEntry) LogInfo() {
	for _, v := range le.LoggerSettings.Targets {
		v.LogInfo(le.Data)
	}
}

func (le *LogEntry) LogError() {
	for _, v := range le.LoggerSettings.Targets {
		v.LogError(le.Data)
	}
}

func (le *LogEntry) LogCritical() {
	for _, v := range le.LoggerSettings.Targets {
		v.LogCritical(le.Data)
	}
}

func (le *LogEntry) LogDebug() {
	for _, v := range le.LoggerSettings.Targets {
		v.LogDebug(le.Data)
	}
}

func (le *LogEntry) LogNone() {
	for _, v := range le.LoggerSettings.Targets {
		v.LogNone(le.Data)
	}
}

func (le *LogEntry) LogTrace() {
	for _, v := range le.LoggerSettings.Targets {
		v.LogTrace(le.Data)
	}
}

func (le *LogEntry) LogWarning() {
	for _, v := range le.LoggerSettings.Targets {
		v.LogWarning(le.Data)
	}
}

func (le *LogEntry) LogFatal(ExitCode int) {
	for _, v := range le.LoggerSettings.Targets {
		v.LogFatal(le.Data, ExitCode)
	}
}

func (le *LogEntry) LogInternalError(err error) {
	for _, v := range le.LoggerSettings.Targets {
		v.LogCritical(le.Data)
	}
}
