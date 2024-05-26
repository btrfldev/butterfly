package butterfly

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
)

//			  1   2  3  4  5   6
//1 string = key;key;-;key;-;key;
//               3          5
//2 string = numoftrash:numoftrash
//3 string = user:password;-:-;user:password;
//[6; infinitive] = value
// 					value
//					  -
//					value
//					  -
//					value

func GetKeySpace(fullpath string) (KeySpace []string, err error) {
	file, err := os.OpenFile(fullpath, os.O_RDONLY, 0600)
	if err != nil {
		return nil, errors.New("Can`t open the file while getting Key Space: " + err.Error())
	}

	defer file.Close()
	kspstr, _, err := GetLineByNum(file, 0)
	if err != nil {
		return nil, errors.New("Can`t get the line whith a Key Space: " + err.Error())
	}

	KeySpace = strings.Split(string(kspstr), ";")
	return KeySpace, nil
}

func LineCounter(r io.Reader /*fullpath string*/) (int, error) {
	buf := make([]byte, 1*1024) //1 Kbyte
	count := 0
	lineSep := []byte{'\n'}

	/*file, err := os.OpenFile(fullpath, os.O_RDONLY, 0600)
	if err != nil {
		return count, errors.New("Error occurred when determining the last line of the file: " + err.Error())
	}

	defer file.Close()*/

	for {
		c, err := r.Read(buf) //file.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func GetLineByNum(r io.Reader, lineNum int) (line []byte, lastLine int, err error) {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		if lastLine == lineNum {
			return sc.Bytes(), lastLine, sc.Err()
		}
		lastLine++
	}
	if lastLine < lineNum {
		return line, lastLine, io.EOF
	} else {
		return line, lastLine, nil
	}
}
