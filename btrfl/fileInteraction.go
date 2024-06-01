package btrfl

import (
	"bufio"
	"bytes"
	"errors"
	//"fmt"
	"io"
	"os"
)


func LineCounter(r io.Reader /*fullpath string*/) (int, error) {
	sterr := "btrfl.fileinteraction.LineCounter"
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
			return count, errors.New(sterr + ": " + err.Error())
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
		return nil, lastLine, io.EOF
	} else {
		return line, lastLine, nil
	}
}

func WriteLineByNum(RWfile *os.File, line string, lineNum int) (err error) {
	sterr := "btrfl.fileinteraction.WriteLineByNum"

	lastLine := 0
	sc := bufio.NewScanner(RWfile)
	for sc.Scan() {
		if lastLine == lineNum {
			/*_, err := RWfile.WriteAt([]byte(line), 5)//fmt.Fprint(RWfile, line)
			if err != nil {
				return errors.New(sterr + ": " + err.Error())
			}*/
			curline := sc.Text()
			if _, err := io.WriteString(RWfile, curline); err!=nil{
				return errors.New(sterr + ": " + err.Error())
			}
		}

		lastLine++
	}
	if err := sc.Err(); err!=nil{
		return errors.New(sterr + ": " + err.Error())
	}
	if lastLine<lineNum {
		return io.EOF
	} else {
		return nil
	}
	/*reader := bufio.NewReader(RWfile)
	var buffer bytes.Buffer
	var line string

	for {
		b, _, err := reader.ReadLine()
		if err==io.EOF{
			break
		}
		line
	}*/
}
