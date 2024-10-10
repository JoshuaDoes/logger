package logger

/*
	verbosity := 2 (0 = default, 1 = debug, 2 = trace)
	log := logger.NewLogger(verbosity)
	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	log.Fatal("Bye.")
	// Calls panic() after logging
	log.Panic("I'm bailing.")
*/

import (
	//Advanced logger
	"github.com/sirupsen/logrus"

	//Better formatting for logrus
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type Logger struct {
	*logrus.Logger
	prefix    string
	Verbosity int
}

//NewLogger returns a logger with the specified verbosity level.
// prefix: string: the prefix for the logger
// verbosity: int: declares the verbosity level
//  - 0: default logging (info, warning, error)
//  - 1: includes 0, plus debug logging
//  - 2: includes 1, plus trace logging
func NewLogger(prefix string, verbosity int) *Logger {
	formatter := new(prefixed.TextFormatter)
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "Mon, Jan _2, 2006 - 03:04:05.000 PM MST"
	formatter.SetColorScheme(&prefixed.ColorScheme{
		InfoLevelStyle:  "green",
		WarnLevelStyle:  "orange",
		ErrorLevelStyle: "red",
		FatalLevelStyle: "purple",
		PanicLevelStyle: "white",
		DebugLevelStyle: "cyan",
		PrefixStyle:     "white+b",
		TimestampStyle:  "black+h",
	})

	log := logrus.New()
	log.Formatter = formatter

	switch verbosity {
	case 2:
		log.Level = logrus.TraceLevel
	case 1:
		log.Level = logrus.DebugLevel
	case 0:
		log.Level = logrus.InfoLevel
	}

	return &Logger{
		Logger: log,
		prefix: prefix,
		Verbosity: verbosity,
	}
}

func (logger *Logger) Trace(args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Trace(args...)
}
func (logger *Logger) Debug(args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Debug(args...)
}
func (logger *Logger) Info(args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Info(args...)
}
func (logger *Logger) Warn(args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Warn(args...)
}
func (logger *Logger) Error(args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Error(args...)
}
func (logger *Logger) Fatal(args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Fatal(args...)
}
func (logger *Logger) Panic(args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Panic(args...)
}

func (logger *Logger) Traceln(args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Traceln(args...)
}
func (logger *Logger) Debugln(args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Debugln(args...)
}
func (logger *Logger) Infoln(args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Infoln(args...)
}
func (logger *Logger) Warnln(args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Warnln(args...)
}
func (logger *Logger) Errorln(args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Errorln(args...)
}
func (logger *Logger) Fatalln(args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Fatalln(args...)
}
func (logger *Logger) Panicln(args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Panicln(args...)
}

func (logger *Logger) Tracef(fmt string, args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Tracef(fmt, args...)
}
func (logger *Logger) Debugf(fmt string, args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Debugf(fmt, args...)
}
func (logger *Logger) Infof(fmt string, args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Infof(fmt, args...)
}
func (logger *Logger) Warnf(fmt string, args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Warnf(fmt, args...)
}
func (logger *Logger) Errorf(fmt string, args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Errorf(fmt, args...)
}
func (logger *Logger) Fatalf(fmt string, args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Fatalf(fmt, args...)
}
func (logger *Logger) Panicf(fmt string, args ...interface{}) {
	logger.WithField("prefix", logger.prefix).Panicf(fmt, args...)
}

