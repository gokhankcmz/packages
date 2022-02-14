package LogTargets

type ILogTarget interface {
	LogInfo(map[string]string)
	LogError(map[string]string)
	LogCritical(map[string]string)
	LogDebug(map[string]string)
	LogNone(map[string]string)
	LogTrace(map[string]string)
	LogWarning(map[string]string)
	LogFatal(map[string]string, int)
}
