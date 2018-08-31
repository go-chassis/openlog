package openlogging

type Tags map[string]interface{}

// Logger is a interface for log tool
type Logger interface {
	Debug(message string, t ...Tags)
	Info(message string, t ...Tags)
	Warn(message string, t ...Tags)
	Error(message string, t ...Tags)
	Fatal(message string, t ...Tags)

	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
}
