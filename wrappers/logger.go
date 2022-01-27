package wrappers

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type Logger interface {
	Timer(name string)
	Info(message string)
	Warning(message string)
	Error(err error)
	Halt(err error)

	Sync()
}

func NewNormalAbnormalLogger(normalWriter, abnormalWriter *os.File) Logger {
	return &normalAbnormalLogger{
		normalWriter:   normalWriter,
		abnormalWriter: abnormalWriter,
		mutex:          &sync.Mutex{},
		timers:         make(map[string]Timer),
	}
}

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

	os.Exit(LOG_HALT)
}

func (lw normalAbnormalLogger) Sync() {
	lw.normalWriter.Sync()
	lw.abnormalWriter.Sync()
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

		os.Exit(ERROR_WRITING_TO_LOG)
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

		os.Exit(INCOMPLETE_WRITE_TO_LOG)
	}
}

const LOG_HALT = 1
const ERROR_WRITING_TO_LOG = 2
const INCOMPLETE_WRITE_TO_LOG = 3
