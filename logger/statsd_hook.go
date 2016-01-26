package logger

import (
	"fmt"
	"os"

	"github.com/bbc/mozart-api-common/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"github.com/bbc/mozart-api-common/Godeps/_workspace/src/gopkg.in/alexcesaro/statsd.v1"
)

type StatsDHook struct {
	Client *statsd.Client
}

func NewStatsDHook() *StatsDHook {
	return &StatsDHook{nil}
}

func (hook *StatsDHook) Fire(entry *logrus.Entry) error {
	event := entry.Data["event"]
	if event != nil {
		client, err := hook.getClient()
		if err != nil {
			return err
		}
		client.Increment(event.(string))
	}
	return nil
}

func (hook *StatsDHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

func (hook *StatsDHook) getClient() (*statsd.Client, error) {
	if hook.Client == nil {
		client, err := statsd.New(
			fmt.Sprintf("%s:%s", os.Getenv("STATSD_HOST"), os.Getenv("STATSD_PORT")),
			statsd.WithPrefix("mozart-config-api."),
		)
		if err != nil {
			return client, err
		}
		hook.Client = client
	}
	return hook.Client, nil
}
