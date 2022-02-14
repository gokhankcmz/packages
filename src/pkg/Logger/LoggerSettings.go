package Logger

import (
	"Packages/src/pkg/Logger/Event"
	"Packages/src/pkg/Logger/LogTargets"
	"Packages/src/pkg/Logger/Rule"
	"os"
)

type Logger struct {
	Rules           []Rule.Rule
	Targets         []LogTargets.ILogTarget
	Enrichments     []func(event Event.Event, data map[string]string)
	ApplicationName string
	HostName        string
}

func NewLogger(ApplicationName string) *Logger {
	hostName, _ := os.Hostname()
	return &Logger{
		ApplicationName: ApplicationName,
		HostName:        hostName,
	}
}

func (ls *Logger) SetApplicationName(applicationName string) *Logger {
	ls.ApplicationName = applicationName
	return ls
}

func (ls *Logger) LogTo(EndpointsToIgnore ...LogTargets.ILogTarget) *Logger {
	for _, v := range EndpointsToIgnore {
		ls.Targets = append(ls.Targets, v)
	}
	return ls
}
func (ls *Logger) CreateLog() *LogEntry {
	entry := newEntry(ls)
	entry.printBasics()
	return entry
}
func (ls *Logger) SetRules(Rules ...Rule.Rule) *Logger {
	for _, v := range Rules {
		ls.Rules = append(ls.Rules, v)
	}
	return ls
}

func (ls *Logger) SetRuleFamily(Rules ...[]Rule.Rule) *Logger {
	for _, v := range Rules {
		for _, v2 := range v {
			ls.Rules = append(ls.Rules, v2)
		}
	}
	return ls
}

func (ls *Logger) Enrich(Enrichments ...func(event Event.Event, data map[string]string)) *Logger {
	for _, v := range Enrichments {
		ls.Enrichments = append(ls.Enrichments, v)
	}
	return ls
}
