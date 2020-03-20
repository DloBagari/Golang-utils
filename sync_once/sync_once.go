package sync_once

import (
	"errors"
	"github.com/sirupsen/logrus"
	"sync"
)

// create an singleton instance using sync.Once

var (
	log     *logrus.Logger
	initLog sync.Once
)

func InitInstance() error {
	err := errors.New("already initialized")
	initLog.Do(func() {
		err = nil
		log = logrus.New()
		log.SetFormatter(&logrus.TextFormatter{})
	})
	return err
}

func Info(message string) {
	log.Info(message)
}
