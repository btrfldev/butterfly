package btrfl

import (
	"errors"
	"os"
	"strings"
)

func GetKeySpace(Rfile *os.File) (KeySpace []string, err error) {
	sterr := "btrfl.GetKeySpace"
	kspbytes, _, err := FI{}.GetLineByNum(Rfile, 0)
	if err != nil {
		return nil, errors.New(sterr + " " + err.Error())
	}
	kspstr := string(kspbytes)

	KeySpace = strings.Split(kspstr, ";")
	return KeySpace, nil
}

func WriteKeySpace(Wfile *os.File, KeySpace []string) (err error) {
	sterr := "btrfl.WriteKeySpace"
	kspstr := strings.Join(KeySpace, ";")

	if err = FI.WriteFile(FI{}, Wfile, kspstr); err != nil {
		return errors.New(sterr + ": " + err.Error())
	} else {
		return nil
	}
}

func AppendValues(AWfile *os.File, Rfile *os.File, values []string) (lastAppended int, err error) {
	sterr := "btrfl.AppendValues"
	curLine := 0

	if curLine, err = FI.LineCounter(FI{}, Rfile); err != nil {
		return 0, errors.New(sterr + ": " + err.Error())
	}
	lastAppended += curLine

	for _, str := range values {
		if err := FI.AppendToFile(FI{}, AWfile, str); err != nil {
			return lastAppended, errors.New(sterr + ": " + err.Error())
		}
		lastAppended++
	}
	return lastAppended, nil
}
