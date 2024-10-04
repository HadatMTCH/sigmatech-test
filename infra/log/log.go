package log

import (
	"base-api/constants"
	"base-api/infra/log_rotator"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	rotator = &log_rotator.Logger{
		Filename: constants.LogFile,
	}
)

func InitializeLogger() {
	logrus.SetOutput(rotator)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logFormat{})
}

type logFormat struct {
	logrus.TextFormatter
}

// Format is a function that creating format for custom logging
func (f *logFormat) Format(entry *logrus.Entry) ([]byte, error) {
	var dateTime, message, logLevel, logMessage, callerFile string
	var userID, request, response, endpoint, errors, responseMessages interface{}
	var callerLine int
	var finalFormat strings.Builder

	userID = entry.Data["ID"]
	if userID == nil {
		userID = "Unknown"
	}
	request = entry.Data["Request"]
	response = entry.Data["Response"]
	endpoint = entry.Data["Endpoint"]
	errors = entry.Data["Error"]
	responseMessages = entry.Data["Messages"]
	if errors != "" {
		logLevel = strings.ToUpper(logrus.ErrorKey)
	} else {
		logLevel = "INFO"
	}
	dateTime = entry.Time.Format(constants.LogDateFormatWithTime)
	logMessage = entry.Message
	callerFile = entry.Caller.File
	callerLine = entry.Caller.Line
	// Defining the line for logging that contain the log file and what is the level and message.
	finalFormat.WriteString(fmt.Sprintf("\n\n[%s - %v] %s : %s", dateTime, userID, logLevel, logMessage))
	if endpoint != nil {
		// Defining the line for logging that contain the log file and what endpoint the user accessing is.
		finalFormat.WriteString(fmt.Sprintf("\n[%s - %v] ENDPOINT : %s", dateTime, userID, endpoint))
	}

	if request != nil {
		// Defining the line for logging that contain the log file and what request body the user send is.
		finalFormat.WriteString(fmt.Sprintf("\n[%s - %v] REQUEST : %v", dateTime, userID, request))
	}

	if response != nil {
		// Defining the line for logging that contain the log file and what response body the user get is.
		finalFormat.WriteString(fmt.Sprintf("\n[%s - %v] RESPONSE : %v", dateTime, userID, response))
	}

	if responseMessages != nil {
		finalFormat.WriteString(fmt.Sprintf("\n[%s - %v] RESPONSE MESSAGE  : %v", dateTime, userID, responseMessages))
	}

	if errors != nil {
		finalFormat.WriteString(fmt.Sprintf("\n[%s - %v] ERROR RESPONSE : %v", dateTime, userID, errors))
	}

	// Defining the line for logging that contains the caller or location of the logger is called
	finalFormat.WriteString(fmt.Sprintf("\n[%s - %v] FILE : %s:%d", dateTime, userID, callerFile, callerLine))

	message = finalFormat.String()
	return []byte(message), nil
}
