package ammana

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

//logger
type logger struct {
	traceHandle   io.Writer
	infoHandle    io.Writer
	warningHandle io.Writer
	errorHandle   io.Writer
}

//newLogger, logger constructor
func newLogger() *logger {
	return &logger{
		traceHandle:   ioutil.Discard,
		infoHandle:    os.Stdout,
		warningHandle: os.Stdout,
		errorHandle:   os.Stderr,
	}
}

//trace
func (l *logger) trace() *log.Logger {
	return log.New(l.traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

//info
func (l *logger) info() *log.Logger {
	return log.New(l.infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

//warning
func (l *logger) warning() *log.Logger {
	return log.New(l.warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

//error
func (l *logger) error() *log.Logger {
	return log.New(l.errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
