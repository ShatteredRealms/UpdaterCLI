package log

import (
    "fmt"
    "time"
)

// LoggerService Handles logging information to the screen and passing to prometheus metrics when needed
type LoggerService interface {
    // Log Write message and log level to screen if maximum log level isn't exceeded
    Log(level logLevel, message string)

    // Logf Format message and Log
    Logf(level logLevel, message string, format ...interface{})

    // Info Writes Info message to screen with Log
    Info(message string)

    // Infof Format message and Info Log
    Infof(message string, format ...interface{})
}

type logger struct {
    timeFormat  string
    maxLogLevel int
}

type logLevel struct {
    Name  string
    Level int
}

var (
    Error   = logLevel{"Error", 0}
    Warning = logLevel{"Warning", 1}
    Info    = logLevel{"Info", 2}
    Debug   = logLevel{"Debug", 3}
    Verbose = logLevel{"Verbose", 4}
)

// NewLogger Creates a new logger
func NewLogger(max logLevel, format string) LoggerService {
    if format == "" {
        format = "2006-01-02 15:04:05"
    }

    return logger{
        timeFormat:  format,
        maxLogLevel: max.Level,
    }
}

// Log Writes the message to standard out with the current time, log level and message if the logger max log level is
// greater than or equal to the given log level
func (l logger) Log(level logLevel, message string) {
    if level.Level <= l.maxLogLevel {
        fmt.Printf("%s | %s: %s\n", l.formattedTime(), level.Name, message)
    }
}

func (l logger) Logf(level logLevel, message string, format ...interface{}) {
    l.Log(level, fmt.Sprintf(message, format...))
}

func (l logger) Info(message string) {
    l.Log(Info, message)
}

func (l logger) Infof(message string, format ...interface{}) {
    l.Log(Info, fmt.Sprintf(message, format...))
}

// Formats the current time to the logger time format
func (l logger) formattedTime() string {
    return time.Now().Format(l.timeFormat)
}