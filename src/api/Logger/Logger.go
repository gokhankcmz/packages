package Logging

import (
	"Packages/src/Configs"
	"Packages/src/pkg/Logger"
	"Packages/src/pkg/Logger/LogLevel"
	"Packages/src/pkg/Logger/LogTargets"
	"Packages/src/pkg/Logger/Rule"
	"Packages/src/pkg/Logger/Subject"
	"Packages/src/pkg/Logger/When"
	"Packages/src/pkg/Logger/With"
)

func GetLogger(config *Configs.AppConfig) *Logger.Logger {
	logger := Logger.NewLogger(config.ApplicationName).
		Enrich(
			With.RequestInfo,
			With.ResponseInfo,
		).
		SetRuleFamily(
			Rule.ForFamily(Subject.SuccessfulResponses(), When.ResponseTimeBiggerThan(1000), LogLevel.Info, 0),
			Rule.ForFamily(Subject.ClientErrors(), When.Always, LogLevel.Info, 0),
			Rule.ForFamily(Subject.ClientErrors(), When.ResponseTimeBiggerThan(1000), LogLevel.Error, 1),
			Rule.ForFamily(Subject.ServerErrors(), When.Always, LogLevel.Error, 0),
		).
		LogTo(LogTargets.Terminal("level"))

	return logger
}

/*
func DynamicLogger(configs *Configs.AppConfig) *Logger.Logger {
	s := configs.LoggerSettings
	logger := Logger.NewLogger(configs.ApplicationName).
		Enrich(
			With.RequestInfoBW(s.PrintRequestInfo),
			With.ResponseInfoBW(s.PrintResponseInfo),
		).
		SetRuleFamily(
			Rule.ForFamily(Subject.SuccessfulResponses(), When.Pass(s.LogSuccessful.Active), loglevel(s.LogSuccessful.Loglevel), 0),
			Rule.ForFamily(Subject.SuccessfulResponses(), When.ResponseBodyBiggerThan(s.MaxRespDuration), loglevel(s.LogSuccessful.Loglevel), 1),
			Rule.ForFamily(Subject.ClientErrors(), When.Pass(s.LogClientErrors.Active), loglevel(s.LogClientErrors.Loglevel), 0),
			Rule.ForFamily(Subject.ClientErrors(), When.ResponseTimeBiggerThan(s.MaxRespDuration), LogLevel.Error, 1),
			Rule.ForFamily(Subject.ServerErrors(), When.Pass(s.LogServerErrors.Active), loglevel(s.LogServerErrors.Loglevel), 0),
		)
	if s.LogToTerminal {
		logger.LogTo(LogTargets.Terminal(s.LogLevelKeyword))
	}

	return logger
}

func loglevel(level string) LogLevel.Loglevel {
	switch strings.ToLower(level) {
	case "info":
		return LogLevel.Info
	case "error":
		return LogLevel.Error
	default:
		return LogLevel.None
	}
}
*/
