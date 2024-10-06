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

// Panic
func Panic(message ...any) {
	log.Panic(message...)
}
func (myLog *Logger) Panic(message ...any) {
	myLog.log(PANIC, Red+"PAN"+Reset, message...)
	os.Exit(101)
}

// Fatal
func Fatal(message ...any) {
	log.Fatal(message...)
}
func (myLog *Logger) Fatal(message ...any) {
	myLog.log(FATAL, Red+"FAT"+Reset, message...)
	os.Exit(1)
}

// Error
func Error(message ...any) {
	log.Error(message...)
}
func (myLog *Logger) Error(message ...any) {
	myLog.log(ERROR, Red+"ERR"+Reset, message...)
}

// Warn
func Warn(message ...any) {
	log.Warn(message...)
}
func (myLog Logger) Warn(message ...any) {
	myLog.log(WARN, Yellow+"WRN"+Reset, message...)
}

//Info
func Info(message ...any) {
	log.Info(message...)
}
func Println(message ...any) {
	log.Info(message...)
}
func (myLog Logger) Info(message ...any) {
	myLog.log(INFO, Green+"INF"+Reset, message...)
}

// Debug
func Debug(message ...any) {
	log.Debug(message...)
}
func (myLog Logger) Debug(message ...any) {
	myLog.log(DEBUG, Blue+"DEB"+Reset, message...)
}

// convert...
func (myLog Logger) log(level int, prefix string, message ...any) {
	if level>=myLog.currLevel {
		defaultlog.Print(prefix +": ", (fmt.Sprintln(message...)))
	}
}
/* Set the debug level
   SetDebugLevel(1) | verbose
   1 = verbose
   2 = debug
   3 = info
   4 = warning
   5 = error
*/
func SetDebugLevel(level int) {
	log.SetDebugLevel(level)
}
func (myLog *Logger) SetDebugLevel(level int) {
	myLog.currLevel = level
}

var log Logger;
