package LogLevel

type Loglevel int

const (
	Info Loglevel = iota
	Error
	Critical
	Debug
	None
	Trace
	Warning
	Fatal
)

//String Returns the string equivalent of the enum.
func (hcs Loglevel) String() string {
	return [...]string{"info", "error", "critical", "debug", "", "trace", "warning", "Fatal"}[hcs]
}
