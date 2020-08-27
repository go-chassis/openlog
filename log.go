package openglog

import (
	"log"
)

const (
	info  = "INFO: "
	debug = "DEBUG: "
	error = "ERROR: "
	warn  = "WARN: "
	fatal = "FATAL: "
)

type golog struct {
}

func (l golog) Debug(message string, opts ...Option) {
	log.Println(debug + message)
}

func (l golog) Info(message string, opts ...Option) {
	log.Println(info + message)
}
func (l golog) Warn(message string, opts ...Option) {
	log.Println(warn + message)
}
func (l golog) Error(message string, opts ...Option) {
	log.Println(error + message)
}
func (l golog) Fatal(message string, opts ...Option) {
	log.Panic(fatal + message)
}
