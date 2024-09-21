package logger

import (
	"errors"
	"fmt"
)

type Logger struct {
	LogLevel               int
	LogToTerminal          bool
	CustomPublishErrMethod func(err error, errLocationInCode, message string, ErrLevel int) error
}

const (
	LogLevelInfo  = 0
	LogLevelWarn  = 1
	LogLevelError = 2
	LogLevelCrash = 3
)

func (l *Logger) Info(err error, errLocationInCode, message string) {
	if l.LogLevel <= LogLevelInfo {
		str := fmt.Sprintf("(" + errLocationInCode + ")" + "\n  " + message + ": " + err.Error())
		if !l.LogToTerminal {
			pubErr := l.CustomPublishErrMethod(err, errLocationInCode, message, LogLevelInfo)
			if pubErr != nil {
				fmt.Println("Can`t Publish Error!" + "[INFO]" + str)
			}
		} else {
			fmt.Println("[INFO]" + str)
		}
	}

}

func (l *Logger) Warn(err error, errLocationInCode, message string) {
	if l.LogLevel <= LogLevelWarn {
		str := fmt.Sprintf("(" + errLocationInCode + ")" + "\n  " + message + ": " + err.Error())
		if !l.LogToTerminal {
			pubErr := l.CustomPublishErrMethod(err, errLocationInCode, message, LogLevelWarn)
			if pubErr != nil {
				fmt.Println("Can`t Publish Error!" + "[WARN]" + str)
			}
		} else {
			fmt.Println("[WARN]" + str)
		}
	}

}

func (l *Logger) Error(err error, errLocationInCode, message string) {
	if l.LogLevel <= LogLevelError {
		str := fmt.Sprintf("(" + errLocationInCode + ")" + "\n  " + message + ": " + err.Error())
		if !l.LogToTerminal {
			pubErr := l.CustomPublishErrMethod(err, errLocationInCode, message, LogLevelError)
			if pubErr != nil {
				fmt.Println("Can`t Publish Error!" + "[Error]" + str)
			}
		} else {
			fmt.Println("[Error]" + str)
		}
	}

}

func (l *Logger) Crash(err error, errLocationInCode, message string) {
	if l.LogLevel <= LogLevelCrash {
		str := fmt.Sprintf("(" + errLocationInCode + ")" + "\n  " + message + ": " + err.Error())
		if !l.LogToTerminal {
			pubErr := l.CustomPublishErrMethod(err, errLocationInCode, message, LogLevelCrash)
			if pubErr != nil {
				panic("Can`t Publish Error!" + "[Panic]" + str)
			}
		} else {
			panic("[Error]" + str)
		}
	}

}

func CheckErr(err error, errLocationInCode, message string) error {
	if err != nil {
		str := fmt.Sprintf(errLocationInCode + "\n  " + message + ": " + err.Error())
		return errors.New(str)
	} else {
		return nil
	}
}

func NewErr(errLocationInCode, message string) error {
	str := fmt.Sprintf(errLocationInCode + "\n  " + message)
	return errors.New(str)
}
