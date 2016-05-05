package logger

import (
	"net/http"
	"os"

	"github.com/bbc/mozart-api-common/Godeps/_workspace/src/github.com/Sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Formatter = new(logrus.JSONFormatter)
	log.Out = os.Stderr // this should be the default but if this fails I'll try os.Stderr

	hook := NewStatsDHook()
	log.Hooks.Add(hook)
}

// Log is a HTTP log wrapper abstraction
func Log(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(logrus.Fields{
			"url":    r.RequestURI,
			"method": r.Method,
		}).Info(name)
		inner.ServeHTTP(w, r)
	})
}

// Info is a wrapper abstraction
func Info(message map[string]interface{}) {
	log.WithFields(message).Info()
}

// Error is a wrapper abstraction
func Error(message map[string]interface{}) {
	log.WithFields(message).Error("error")
}

// Warn is a wrapper abstraction
func Warn(message map[string]interface{}) {
	log.WithFields(message).Warn()
}

// Debug is a wrapper abstraction
func Debug(message map[string]interface{}) {
	log.WithFields(message).Debug()
}

// Panic is a wrapper abstraction
func Panic(message map[string]interface{}) {
	log.WithFields(message).Panic("panic")
}

// Fatal is a wrapper abstraction
func Fatal(message map[string]interface{}) {
	log.WithFields(message).Fatal("fatal")
}
