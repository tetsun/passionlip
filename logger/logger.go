package logger

import "github.com/sirupsen/logrus"

/*
DavLog for dav access
*/
func DavLog(code int, method string, remote string, msg string) {
	logrus.WithFields(logrus.Fields{
		"code":   code,
		"method": method,
		"remote": remote,
	}).Info(msg)
}

/*
Fatal instead of log.Fatal
*/
func Fatal(msg string) {
	logrus.Fatal(msg)
}
