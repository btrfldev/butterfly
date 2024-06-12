package btrfl

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
)

type FI struct{}

func (FI)LineCounter(r io.Reader) (int, error) {
	sterr := "btrfl.fileinteraction.LineCounter"
	buf := make([]byte, 1*1024) //1 Kbyte
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, errors.New(sterr + ": " + err.Error())
		}
	}
}

func (FI)GetLineByNum(r io.Reader, lineNum int) (line []byte, lastLine int, err error) {
	sterr := "btrfl.fileinteraction.GetLineByNum"
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		if lastLine == lineNum {
			return sc.Bytes(), lastLine, nil
		}
		if sc.Err()!=nil{
			return nil, lastLine, errors.New(sterr + ": " + sc.Err().Error())
		}
		lastLine++
	}
	if lastLine < lineNum {
		return nil, lastLine, io.EOF
	} else {
		return line, lastLine, nil
	}
}

func (FI)WriteFile(w *os.File, data string) (err error) {
	sterr := "btrfl.fileinteraction.WriteFullFile"
	if _, err := w.WriteString(data); err != nil {
		return errors.New(sterr + ": " + err.Error())
	} else {
		return nil
	}
}

func (FI)AppendToFile(aw *os.File, data string) (err error) {
	sterr := "btrfl.fileinteraction.AppendToFile"
	if _, err := aw.WriteString("\n" + data); err != nil {
		return errors.New(sterr + ": " + err.Error())
	} else {
		return nil
	}
}
