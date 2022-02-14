package Logger

import (
	"Packages/src/pkg/Logger/Event"
	"Packages/src/pkg/Logger/LogLevel"
	"Packages/src/pkg/Logger/Rule"
	"errors"
	"sort"
)

func (le *LogEntry) LogWithRules() {
	PassingRules, err := getPassingRules(&le.LoggerSettings.Rules, &le.Event)
	if err != nil {
		le.LogInternalError(err)
	} else if len(PassingRules) > 0 {
		sort.Slice(PassingRules, func(i, j int) bool {
			return PassingRules[j].PriorityLevel < PassingRules[i].PriorityLevel
		})
		le.printEnrichments()
		switch PassingRules[0].Level {
		case LogLevel.Info:
			le.LogInfo()
		case LogLevel.Error:
			le.LogError()
		case LogLevel.Debug:
			le.LogDebug()
		case LogLevel.None:
			le.LogNone()
		case LogLevel.Critical:
			le.LogCritical()
		case LogLevel.Warning:
			le.LogWarning()
		case LogLevel.Trace:
			le.LogTrace()
		}
	}
}

func getPassingRules(rules *[]Rule.Rule, Event *Event.Event) ([]Rule.Rule, error) {
	var PassingRules []Rule.Rule
	if Event.Request == nil {
		return nil, errors.New("Request information not found. ")
	}
	if Event.ResponseCode == 0 {
		return nil, errors.New("Response code is empty. ")
	}

	for _, rule := range *rules {
		if rule.Exec(Event) {
			PassingRules = append(PassingRules, rule)
		}
	}
	return PassingRules, nil
}
