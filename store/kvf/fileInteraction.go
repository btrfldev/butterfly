package kvf

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
)

type FI struct{}

func (FI)LineCounter(r io.Reader) (int, error) {
	sterr := "btrfl.store.fileinteraction.LineCounter"
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

func (FI)GetLinesByNums(r io.Reader, linesNums []int) (lines map[int][]byte, lastLine int, err error) {
	sterr := "btrfl.store.fileinteraction.GetLineByNum"
	sc := bufio.NewScanner(r)
	result := make(map[int][]byte) //lines

	found := 0
	need := len(linesNums)

	biggestnum := 0
	for _, n := range linesNums {
		if biggestnum < n {
			biggestnum = n
		}
	}

	for sc.Scan() {
		for _, n := range linesNums{
			if lastLine == n{
				result[n] = sc.Bytes()
				found ++
			}
			if need == found {
				return result, lastLine, nil
			}
		}
		/*if lastLine == lineNum {
			return sc.Bytes(), lastLine, nil
		}*/
		if sc.Err()!=nil{
			return nil, lastLine, errors.New(sterr + ": " + sc.Err().Error())
		}
		lastLine++
	}
	if lastLine < biggestnum {
		return nil, lastLine, errors.New(sterr + ": " + io.EOF.Error())
	} else {
		return result, lastLine, nil
	}
}

func (FI)WriteFile(w *os.File, data string) (err error) {
	sterr := "btrfl.store.fileinteraction.WriteFullFile"
	if _, err := w.WriteString(data); err != nil {
		return errors.New(sterr + ": " + err.Error())
	} else {
		return nil
	}
}

func (FI)AppendToFile(aw *os.File, data string) (err error) {
	sterr := "btrfl.store.fileinteraction.AppendToFile"
	if _, err := aw.WriteString("\n" + data); err != nil {
		return errors.New(sterr + ": " + err.Error())
	} else {
		return nil
	}
}