package LogTargets

import log "github.com/sirupsen/logrus"

type Logrus struct{}

func (l Logrus) LogInfo(Data map[string]string) {
	entry := getLogrusEntry(Data)
	entry.Info()
}

func (l Logrus) LogError(Data map[string]string) {
	entry := getLogrusEntry(Data)
	entry.Error()
}

func (l Logrus) LogPanic(Data map[string]string) {
	entry := getLogrusEntry(Data)
	entry.Panic()
}

func getLogrusEntry(Data map[string]string) *log.Entry {
	var entry *log.Entry
	for key, value := range Data {
		if entry == nil {
			entry = log.WithField(key, value)
		}
		entry = entry.WithField(key, value)
	}
	return entry
}
