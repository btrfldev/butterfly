package btrfl

import (
	"errors"
	"os"
	"strings"
)

func GetKeySpace( /*fullpath string*/ Rfile *os.File) (KeySpace []string, err error) {
	sterr := "btrfl.meta.GetKeySpace"
	/*file, err := os.OpenFile(fullpath, os.O_RDONLY, 0600)
	if err != nil {
		return nil, errors.New("Can`t open the file while getting Key Space: " + err.Error())
	}

	defer file.Close()*/
	kspbytes, _, err := GetLineByNum(Rfile, 0)
	kspstr := string(kspbytes)
	if err != nil {
		return nil, errors.New(sterr + " " + err.Error())
	}

	KeySpace = strings.Split(kspstr, ";")
	return KeySpace, nil
}

func UpdateKeySpace(WRfile *os.File, /*Rfile *os.File,*/ KeySpace []string) (err error) {
	sterr := "btrfl.meta.UpdateKeySpace"
	kspstr := strings.Join(KeySpace, ";")

	err = WriteLineByNum(WRfile, kspstr, 0)
	if err != nil {
		return errors.New(sterr + ": " + err.Error())
	} else {
		return nil
	}
}
