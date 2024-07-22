package kvf

import (
	"errors"
	"io"
	"os"
	"strings"
)

func GetKeySpace(/*Rfile *os.File*/ r io.Reader) (KeySpace []string, err error) {
	sterr := "btrfl.store.kvf.kv.GetKeySpace"
	lkspbytes, _, err := FI{}.GetLinesByNums(r, []int{0})
	if err != nil {
		return nil, errors.New(sterr + " " + err.Error())
	}
	kspstr := string(lkspbytes[0])

	KeySpace = strings.Split(kspstr, ";")
	return KeySpace, nil
}

func WriteKeySpace(Wfile *os.File, KeySpace []string) (err error) {
	sterr := "btrfl.store.kvf.kv.WriteKeySpace"
	kspstr := strings.Join(KeySpace, ";")

	if err = FI.WriteFirstLine(FI{}, Wfile, kspstr); err != nil {
		return errors.New(sterr + ": " + err.Error())
	} else {
		return nil
	}
}

func AppendValues(AWfile *os.File, Rfile *os.File, values []string) (lastAppended int, err error) {
	sterr := "btrfl.store.kvf.kv.AppendValues"
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


func ReadValues(Rfile *os.File, R2file *os.File, KeySpace []string) (kv map[string]string, err error) {
	sterr := "btrfl.store.kvf.kv.ReadValues"
	var ksp []string           //available keyspace
	ki := make(map[string]int) //key:id
	tkv := make(map[string]string)
	
	var r io.Reader = Rfile
	if ksp, err = GetKeySpace(r); err != nil {
		return nil, errors.New(sterr + ": " + err.Error())
	}

	for i, availk := range ksp {
		for _, needk := range KeySpace {
			if availk == needk {
				ki[needk] = i+1
				//println(availk, needk, i)
			}
		}
	}

	nums := []int{}
	for _, i := range ki {
		nums = append(nums, i)
	}
	r = R2file
	bvals, _, err := FI.GetLinesByNums(FI{}, r, nums)
	if err != nil {
		return nil, errors.New(sterr + ": " + err.Error())
	}
	for k, i := range ki {
		tkv[k] = string(bvals[i])
	}

	/*for k, i := range ki {
		var r io.Reader = Rfile
		bval, ll, err := FI.GetLinesByNums(FI{}, r, i+1)
		println(i, ll, string(bval[:]))
		if err != nil {
			return nil, errors.New(sterr + ": " + err.Error())
		}
		tkv[k] = string(bval[:])
	}*/
	return tkv, nil
}
