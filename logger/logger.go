package logger

import(
	"fmt"
	"errors"
)

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