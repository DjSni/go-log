package log

import (
	"fmt"
	"os"
	defaultlog "log"
)

//// Config Log
var (
	PANIC int = 7	//panic
	FATAL int = 6	//fatal
	ERROR int = 5	//error
	WARN int = 4	//warning
	INFO int = 3	//info
	DEBUG int= 2	//debug
	VERBOSE int= 1	//verbose
)

// set colors
var (
	Reset = "\033[0m"		// Reset
	Red = "\033[31m"		// Error
	BRed = "\033[31;1m"		// Error
	Green = "\033[32m"		// Info
	BGreen = "\033[32;1m"	// Info
	Yellow = "\033[33m"		// Warn
	BYellow = "\033[33;1m"	// Warn
	Blue = "\033[34m"		// Debug
	BBlue = "\033[34;1m"	// Debug
	Magenta = "\033[35m"
	BMagenta = "\033[35;1m"	
	Cyan = "\033[36m"		
	BCyan = "\033[36;1m"
	Gray = "\033[37m"		
	BGray = "\033[37;1m"
	White = "\033[97m"
)

type Logger struct{
	currLevel int
}


// Panic logs message at level PANIC and then exits with code 101
func Panic(message ...any) {
	log.Panic(message...)
}
// Panic logs message at level PANIC and then exits with code 101
func (myLog *Logger) Panic(message ...any) {
	myLog.log(PANIC, Red+"PAN"+Reset, message...)
	os.Exit(101)
}

// Panicf logs formatted message at level PANIC and then exits with code 101
func Panicf(format string, args ...any) {
    log.Panicf(format, args...)
}
// Panicf logs formatted message at level PANIC and then exits with code 101
func (myLog *Logger) Panicf(format string, args ...any) {
    myLog.logf(PANIC, Red+"PAN"+Reset, format, args...)
    os.Exit(101)
}

// Fatal logs message at level FATAL and then exits with code 1
func Fatal(message ...any) {
	log.Fatal(message...)
}
// Fatal logs message at level FATAL and then exits with code 1
func (myLog *Logger) Fatal(message ...any) {
	myLog.log(FATAL, Red+"FAT"+Reset, message...)
	os.Exit(1)
}

// Fatalf logs formatted message at level FATAL and then exits with code 1
func Fatalf(format string, args ...any) {
    log.Fatalf(format, args...)
}
// Fatalf logs formatted message at level FATAL and then exits with code 1
func (myLog *Logger) Fatalf(format string, args ...any) {
    myLog.logf(FATAL, Red+"FAT"+Reset, format, args...)
    os.Exit(1)
}

// Error logs message at level ERROR
func Error(message ...any) {
	log.Error(message...)
}
// Error logs a message at the ERROR level using the Logger instance.
// The message can include any number of arguments, which will be formatted
// similarly to fmt.Println outputs. This method does not terminate the program.
func (myLog *Logger) Error(message ...any) {
	myLog.log(ERROR, Red+"ERR"+Reset, message...)
}

// Errorf logs formatted message at level ERROR
func Errorf(format string, args ...any) {
    log.Errorf(format, args...)
}
// Errorf logs a formatted message at the ERROR level
func (myLog *Logger) Errorf(format string, args ...any) {
    myLog.logf(ERROR, Red+"ERR"+Reset, format, args...)
}

// Warn logs message at level WARN. The message can include any number of
// arguments, which will be formatted similarly to fmt.Println outputs. This
// method does not terminate the program.
func Warn(message ...any) {
	log.Warn(message...)
}
// Warn logs message at level WARN. The message can include any number of
// arguments, which will be formatted similarly to fmt.Println outputs. This
// method does not terminate the program.
func (myLog Logger) Warn(message ...any) {
	myLog.log(WARN, Yellow+"WRN"+Reset, message...)
}

// Warnf logs formatted message at level WARN
func Warnf(format string, args ...any) {
    log.Warnf(format, args...)
}
// Warnf logs a formatted message at the WARN level
func (myLog Logger) Warnf(format string, args ...any) {
    myLog.logf(WARN, Yellow+"WRN"+Reset, format, args...)
}

// Info logs a message at the INFO level using the global Logger instance.
// The message can include any number of arguments, which will be formatted
// similarly to fmt.Println outputs. This method does not terminate the program.
func Info(message ...any) {
	log.Info(message...)
}
// Println is a convenience function for logging a message at the INFO level.
// It takes any number of arguments, formats them using fmt.Sprintln, and
// prints the result at the INFO level.
func Println(message ...any) {
	log.Info(message...)
}
// Info logs a message at the INFO level using the Logger instance.
// The message can include any number of arguments, which will be formatted
// similarly to fmt.Println outputs. This method does not terminate the program.
func (myLog Logger) Info(message ...any) {
	myLog.log(INFO, Green+"INF"+Reset, message...)
}

// Infof logs formatted message at level INFO
func Infof(format string, args ...any) {
    log.Infof(format, args...)
}
// Printf is a convenience function for logging a formatted message at the INFO level
func Printf(format string, args ...any) {
    log.Infof(format, args...)
}
// Infof logs a formatted message at the INFO level
func (myLog Logger) Infof(format string, args ...any) {
    myLog.logf(INFO, Green+"INF"+Reset, format, args...)
}

// Debug logs a message at the DEBUG level using the global Logger instance.
// The message can include any number of arguments, which will be formatted
// similarly to fmt.Println outputs. This method does not terminate the program.
func Debug(message ...any) {
	log.Debug(message...)
}
// Debug logs a message at the DEBUG level using the Logger instance. 
// The message can include any number of arguments, which will be formatted 
// similarly to fmt.Println outputs. This method does not terminate the program.
func (myLog Logger) Debug(message ...any) {
	myLog.log(DEBUG, Blue+"DEB"+Reset, message...)
}

// Debugf logs formatted message at level DEBUG
func Debugf(format string, args ...any) {
    log.Debugf(format, args...)
}
// Debugf logs a formatted message at the DEBUG level
func (myLog Logger) Debugf(format string, args ...any) {
    myLog.logf(DEBUG, Blue+"DEB"+Reset, format, args...)
}

// log logs a message at the given level using the Logger instance.
// The message can include any number of arguments, which will be formatted
// similarly to fmt.Println outputs. This method does not terminate the program.
// The prefix is used to identify the log message, and will be printed before
// the message.
func (myLog Logger) log(level int, prefix string, message ...any) {
	if level>=myLog.currLevel {
		defaultlog.Print(prefix +": ", (fmt.Sprintln(message...)))
	}
}

// logf logs a formatted message at the given level
func (myLog Logger) logf(level int, prefix string, format string, args ...any) {
    if level >= myLog.currLevel {
        defaultlog.Print(prefix + ": " + fmt.Sprintf(format, args...))
    }
}

// Set the debug level of the global logger. The level can be one of the following values:
//   1 = verbose, 2 = debug, 3 = info, 4 = warning, 5 = error
func SetDebugLevel(level int) {
	log.SetDebugLevel(level)
}

// SetDebugLevel sets the current logging level for the Logger instance.
// Messages with a level lower than the current level will be ignored.
// The level parameter can be one of the following:
//   1 = verbose, 2 = debug, 3 = info, 4 = warning, 5 = error
func (myLog *Logger) SetDebugLevel(level int) {
	myLog.currLevel = level
}

var log Logger;