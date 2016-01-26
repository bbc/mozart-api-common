package logger

import (
	"net/http"

	"github.com/bbc/mozart-api-common/Godeps/_workspace/src/github.com/Sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Formatter = new(logrus.JSONFormatter)

	hook := NewStatsDHook()
	log.Hooks.Add(hook)
}

func Log(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(logrus.Fields{
			"url":    r.RequestURI,
			"method": r.Method,
		}).Info(name)
		inner.ServeHTTP(w, r)
	})
}

func Info(message map[string]interface{}) {
	log.WithFields(message).Info()
}

func Error(message map[string]interface{}) {
	log.WithFields(message).Error("error")
}

func Warn(message map[string]interface{}) {
	log.WithFields(message).Warn()
}

func Debug(message map[string]interface{}) {
	log.WithFields(message).Debug()
}

func Panic(message map[string]interface{}) {
	log.WithFields(message).Panic("panic")
}

func Fatal(message map[string]interface{}) {
	log.WithFields(message).Fatal("fatal")
}
