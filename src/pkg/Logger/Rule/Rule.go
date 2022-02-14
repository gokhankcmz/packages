package Rule

import (
	"Packages/src/pkg/Logger/Event"
	"Packages/src/pkg/Logger/LogLevel"
	"Packages/src/pkg/Logger/Subject"
	"strconv"
	"strings"
)

type Rule struct {
	Subject       *Subject.Subject
	Condition     func(Event *Event.Event) bool
	Level         LogLevel.Loglevel
	PriorityLevel int
}

func (r *Rule) Exec(Event *Event.Event) bool {
	if r.Subject.SubjectType == Subject.ApiResponse && r.Subject.SubjectName == strconv.Itoa(Event.ResponseCode) {
		return r.Condition(Event)
	}
	if r.Subject.SubjectType == Subject.ApiPath && strings.Contains(Event.Request.URL.Path, r.Subject.SubjectName) {
		return r.Condition(Event)
	}
	return false
}

func For(Subject *Subject.Subject, condition func(Event *Event.Event) bool, level LogLevel.Loglevel, PriorityLevel int) Rule {
	return Rule{
		Subject:       Subject,
		Condition:     condition,
		Level:         level,
		PriorityLevel: PriorityLevel,
	}
}

func ForFamily(Subjects *[]Subject.Subject, condition func(Event *Event.Event) bool, level LogLevel.Loglevel, PriorityLevel int) []Rule {
	var rules []Rule
	for _, subject := range *Subjects {
		s := subject
		rule := Rule{
			Subject:       &s,
			Condition:     condition,
			Level:         level,
			PriorityLevel: PriorityLevel,
		}
		rules = append(rules, rule)
	}
	return rules
}
