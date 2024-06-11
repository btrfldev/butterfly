package btrfl

import (
	"errors"
	"os"
	"strings"
)

func GetKeySpace(Rfile *os.File) (KeySpace []string, err error) {
	sterr := "btrfl.meta.GetKeySpace"
	kspbytes, _, err := GetLineByNum(Rfile, 0)
	if err != nil {
		return nil, errors.New(sterr + " " + err.Error())
	}
	kspstr := string(kspbytes)

	KeySpace = strings.Split(kspstr, ";")
	return KeySpace, nil
}

func WriteKeySpace(Wfile *os.File, KeySpace []string) (err error) {
	sterr := "btrfl.meta.WriteKeySpace"
	kspstr := strings.Join(KeySpace, ";")

	if err = WriteFile(Wfile, kspstr); err != nil {
		return errors.New(sterr + ": " + err.Error())
	} else {
		return nil
	}
}

func AppendValues(Afile *os.File, Rfile *os.File, values []string) (lastAppended int, err error) {
	sterr := "btrfl.meta.WriteKeySpace"
	curLine := 0

	if curLine, err = LineCounter(Rfile); err != nil {
		return 0, errors.New(sterr + ": " + err.Error())
	}
	lastAppended += curLine

	for _, str := range values {
		if err := AppendToFile(Afile, str); err != nil {
			return lastAppended, errors.New(sterr + ": " + err.Error())
		}
		lastAppended++
	}
	return lastAppended, nil
}
