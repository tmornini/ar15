// Package wrappers wraps the stdlib
package wrappers

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"sync"
)

// Logger is the interface that wraps logging
type Logger interface {
	Timer(name string)
	Info(message string)
	Warning(message string)
	Error(err error)
	Halt(err error)

	Sync()
}

// NewNormalAbnormalLogger returns a new NormalAbnormalLogger
func NewNormalAbnormalLogger(normalWriter, abnormalWriter *os.File) Logger {
	return &normalAbnormalLogger{
		normalWriter:   normalWriter,
		abnormalWriter: abnormalWriter,
		mutex:          &sync.Mutex{},
		timers:         make(map[string]Timer),
	}
}

// LogHalt is the exit code for a fatal error
const LogHalt = 1

// ErrorWritingToLog is the exit code for an error writing to the log
const ErrorWritingToLog = 2

// IncompleteWriteToLog is the exit code for an incomplete write to the log
const IncompleteWriteToLog = 3

type normalAbnormalLogger struct {
	normalWriter   *os.File
	abnormalWriter *os.File
	mutex          *sync.Mutex
	timers         map[string]Timer
}

func (lw normalAbnormalLogger) Error(err error) {
	lw.write(
		NewTimestamp(),
		lw.abnormalWriter,
		fmt.Sprintf(
			"ERROR: %s",
			err.Error(),
		),
	)
}

func (lw normalAbnormalLogger) Info(message string) {
	lw.write(
		NewTimestamp(),
		lw.normalWriter,
		"INFO: "+message,
	)
}

func (lw normalAbnormalLogger) Timer(name string) {
	now := NewTimestamp()

	timer, available := lw.timers[name]

	if !available {
		timer = NewTimer(name)
		lw.timers[name] = timer
	}

	timer.RecordIteration(now)

	lw.write(
		now,
		lw.normalWriter,
		timer.String(),
	)
}

func (lw normalAbnormalLogger) Warning(message string) {
	lw.write(NewTimestamp(), lw.normalWriter, "WARN: "+message)
}

func (lw normalAbnormalLogger) Halt(err error) {
	lw.write(
		NewTimestamp(),
		lw.abnormalWriter,
		fmt.Sprintf(
			"HALT: %s",
			err.Error(),
		),
	)

	os.Exit(LogHalt)
}

func (lw normalAbnormalLogger) Sync() {
	err := lw.normalWriter.Sync()
	if err != nil {
		switch err.(type) {
		case *fs.PathError:
			// ignore
		default:
			panic(err)
		}
	}

	err = lw.abnormalWriter.Sync()
	if err != nil {
		switch err.(type) {
		case *fs.PathError:
			// ignore
		default:
			panic(err)
		}
	}
}

func (lw normalAbnormalLogger) write(now Timestamp, writer io.Writer, message string) {
	logMessage := "[" + now.String() + "] " + message + "\n"

	lw.mutex.Lock()
	defer lw.mutex.Unlock()

	bytesWritten, err := writer.Write([]byte(logMessage))

	if err != nil {
		os.Stderr.WriteString(
			"error writing log!\n" +
				logMessage +
				err.Error() + "\n",
		)

		os.Exit(ErrorWritingToLog)
	}

	if bytesWritten != len(logMessage) {
		os.Stderr.WriteString(
			"short write to log!\n" +
				logMessage +
				"\n" +
				fmt.Sprintf(
					"message length: %d, write length: %d",
					len(message),
					bytesWritten,
				),
		)

		os.Exit(IncompleteWriteToLog)
	}
}
