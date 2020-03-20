package logrus_example

import "github.com/sirupsen/logrus"

type Hook struct {
	id string
}

// what levels this hook will fire on
func (h *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

//this method will trigger whenever we log
func (h *Hook) Fire(entry *logrus.Entry) error {
	entry.Data["id"] = h.id
	return nil
}

// set some configs for logrus package
func Logrus() {
	// change the format from json to text
	logrus.SetFormatter(&logrus.TextFormatter{})
	//set level of logging to info
	logrus.SetLevel(logrus.InfoLevel)
	// add hook
	logrus.AddHook(&Hook{id: "12345"})

	//define a field
	fields := logrus.Fields{}
	fields["success"] = true
	fields["complex_struct"] = struct {
		Event string
		When  string
	}{"Something happened", "just now"}
	//add fields to logrus and get logrus entry
	entry := logrus.WithFields(fields)
	// test entry
	entry.Warn("warning message")
	entry.Error("error message")
}
