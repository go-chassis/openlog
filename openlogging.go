package openlogging

var logger Logger

func SetLogger(l Logger) {
	logger = l
}
func GetLogger() Logger {
	return logger
}

//TODO define caller level
func Info(message string, t ...Tags) {
	logger.Info(message, t...)
}
