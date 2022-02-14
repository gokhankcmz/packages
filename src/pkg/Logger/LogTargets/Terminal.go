package LogTargets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type terminal struct {
	LogLevelKeyword string
}

func Terminal(LogLevelKeyword string) *terminal {
	return &terminal{LogLevelKeyword: LogLevelKeyword}
}

func (t *terminal) LogInfo(Data map[string]string) {
	Data[t.LogLevelKeyword] = "info"
	bytearray, _ := marshal(Data)
	fmt.Println(string(bytearray))
}

func (t *terminal) LogError(Data map[string]string) {
	Data[t.LogLevelKeyword] = "error"
	bytearray, _ := marshal(Data)
	fmt.Println(string(bytearray))
}

func (t *terminal) LogCritical(Data map[string]string) {
	Data[t.LogLevelKeyword] = "critical"
	bytearray, _ := marshal(Data)
	fmt.Println(string(bytearray))
}

func (t *terminal) LogDebug(Data map[string]string) {
	Data[t.LogLevelKeyword] = "Debug"
	bytearray, _ := marshal(Data)
	fmt.Println(string(bytearray))
}

func (t *terminal) LogNone(Data map[string]string) {
}

func (t *terminal) LogTrace(Data map[string]string) {
	Data[t.LogLevelKeyword] = "Trace"
	bytearray, _ := marshal(Data)
	fmt.Println(string(bytearray))
}

func (t *terminal) LogWarning(Data map[string]string) {
	Data[t.LogLevelKeyword] = "Warning"
	bytearray, _ := marshal(Data)
	fmt.Println(string(bytearray))
}

func (t *terminal) LogFatal(Data map[string]string, exitCode int) {
	Data[t.LogLevelKeyword] = "Fatal"
	bytearray, _ := marshal(Data)
	fmt.Println(string(bytearray))
	os.Exit(exitCode)
}

func marshal(i interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(i)
	return bytes.TrimRight(buffer.Bytes(), "\n"), err
}
