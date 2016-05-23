package logger

import (
	"net/http"
	"os"

	"github.com/bbc/mozart-api-common/Godeps/_workspace/src/github.com/Sirupsen/logrus"
)

var log = logrus.New()

// Init is used to bootstrap the requirements for this package
func Init() {
	log.Formatter = new(logrus.JSONFormatter)

	if env := os.Getenv("APP_ENV"); env == "test" {
		f, e := os.Create("../../tests.log")
		if e != nil {
			log.Fatal("Failed to create log file for whilst tests are running")
		}
		log.Out = f
	}

	if env := os.Getenv("APP_ENV"); env != "test" {
		hook := NewStatsDHook()
		log.Hooks.Add(hook)
	}
}

// Log is a HTTP logger abstraction
func Log(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(logrus.Fields{
			"url":    r.RequestURI,
			"method": r.Method,
		}).Info(name)
		inner.ServeHTTP(w, r)
	})
}

// Info is a logger abstraction
func Info(message map[string]interface{}) {
	log.WithFields(message).Info()
}

// Error is a logger abstraction
func Error(message map[string]interface{}) {
	log.WithFields(message).Error("error")
}

// Warn is a logger abstraction
func Warn(message map[string]interface{}) {
	log.WithFields(message).Warn()
}

// Debug is a logger abstraction
func Debug(message map[string]interface{}) {
	log.WithFields(message).Debug()
}

// Panic is a logger abstraction
func Panic(message map[string]interface{}) {
	log.WithFields(message).Panic("panic")
}

// Fatal is a logger abstraction
func Fatal(message map[string]interface{}) {
	log.WithFields(message).Fatal("fatal")
}
