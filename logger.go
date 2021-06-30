package logger

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type logger struct {
	dest io.Writer
}

func GetALogger() logger {
	return logger{
		dest: ioutil.Discard,
	}
}

func (l logger) Log(format string) {
	fmt.Fprintln(l.dest, format)
}

func (l logger) LogF(format string, args ...interface{}) {
	fmt.Fprintf(l.dest, format, args...)
}

func (l *logger) Begin() {
	l.dest = os.Stdout
}

func (l *logger) End() {
	l.dest = ioutil.Discard
}
