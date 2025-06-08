package gtools

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogInit(t *testing.T) {
	LogInit("./log", "syn.log")
	logrus.Errorf("testddd:%diaddd", 1000)
	logrus.Infof("info:%diaddd", 1000)
	logrus.WithFields(logrus.Fields{
		"user_id": "12345",
		"action":  "login",
	}).Info("User logged in")
}
