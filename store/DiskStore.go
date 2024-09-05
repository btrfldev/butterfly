package store

import (
	"fmt"
	"os"
	"strings"

	"github.com/iamsoloma/butterfly/store/kvf"
)

// FileSystem store based on .kv
// use the | symbol to differentiate
// filepath | key
type DiskStore struct {
	path string
}

func NewDiskStore(path string) *DiskStore {
	return &DiskStore{
		path: path,
	}
}

// TODO Check
func (d *DiskStore) Put(key string, value string) error {
	//get filepath and key
	paths := strings.Split(key, "|")
	if len(paths) != 2 {
		return fmt.Errorf("the file and path boundaries are not known (%v), use only one '|' symbol for this", key)
	}
	pathtofile := d.path + "/" + paths[0]
	_, err := os.Stat(pathtofile)
	if err != nil {
		if os.IsNotExist(err) {
			//create folders if not exist
			if err = os.MkdirAll(pathtofile, 0777); err != nil {
				return fmt.Errorf("can`t create path to the file (%v): %s", key, err.Error())
			}
			//create file if not exist
			if _, err = os.Create(pathtofile); err != nil {
				return fmt.Errorf("can`t create the file for (%v): %s", key, err.Error())
			}
		} else {
			return fmt.Errorf("can`t open the file for (%v)", key)
		}
	}

	//key is existing?
	found := false
	RFile, err := os.OpenFile(pathtofile, os.O_RDONLY, 0777)
	if err != nil {
		return fmt.Errorf("can`t open the file for (%v) in (%s)", key, paths[0])
	}
	defer RFile.Close()

	ks, err := kvf.GetKeySpace(RFile)
	if err != nil {
		return fmt.Errorf("can`t get key space")
	}
	for _, k := range ks {
		if k == paths[1] {
			found = true
		}
	}

	//putting
	if !found {
		//append
		ks = append(ks, paths[1])
		WFile, err := os.OpenFile(pathtofile, os.O_WRONLY, 0777)
		if err != nil {
			return fmt.Errorf("can`t open the file for (%v) in (%s)", key, paths[0])
		}
		defer WFile.Close()
		if err = kvf.WriteKeySpace(WFile, ks); err != nil {
			return fmt.Errorf("can`t update the keyspace in (%s) for (%v)", paths[0], key)
		}
		AWFile, err := os.OpenFile(pathtofile, os.O_APPEND|os.O_WRONLY, 0777)
		if err != nil {
			return fmt.Errorf("can`t open the file for (%v) in (%s)", key, paths[0])
		}
		defer AWFile.Close()
		RFile, err := os.OpenFile(pathtofile, os.O_RDONLY, 0777)
		if err != nil {
			return fmt.Errorf("can`t open the file for (%v) in (%s)", key, paths[0])
		}
		defer RFile.Close()
		kvf.AppendValues(AWFile, RFile, []string{string(value)})
	} else {
		//rewrite

		//read
		RFile, err := os.OpenFile(pathtofile, os.O_RDONLY, 0777)
		if err != nil {
			return fmt.Errorf("can`t open the file for (%v) in (%s)", key, paths[0])
		}
		defer RFile.Close()
		R2File, err := os.OpenFile(pathtofile, os.O_RDONLY, 0777)
		if err != nil {
			return fmt.Errorf("can`t open the file for (%v) in (%s)", key, paths[0])
		}
		defer R2File.Close()
		R3File, err := os.OpenFile(pathtofile, os.O_RDONLY, 0777)
		if err != nil {
			return fmt.Errorf("can`t open the file for (%v) in (%s)", key, paths[0])
		}
		defer R3File.Close()
		ks, err := kvf.GetKeySpace(RFile)
		if err != nil {
			return fmt.Errorf("can`t read the current keyspace in (%s) for (%v)", paths[0], key)
		}
		kv, err := kvf.ReadValues(R2File, R3File, ks)
		if err != nil {
			return fmt.Errorf("can`t read the current values in (%s) for rewriting file with (%v)", paths[0], key)
		}

		keys, values := []string{}, []string{}

		for k, v := range kv {
			keys = append(keys, k)
			values = append(values, v)
		}

		//write
		WFile, err := os.OpenFile(pathtofile, os.O_WRONLY, 0777)
		if err != nil {
			return fmt.Errorf("can`t open the file for (%v) in (%s)", key, paths[0])
		}
		defer WFile.Close()
		if err = kvf.WriteKeySpace(WFile, keys); err != nil {
			return fmt.Errorf("can`t update the keyspace in (%s) for (%v)", paths[0], key)
		}

		AWFile, err := os.OpenFile(pathtofile, os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			return fmt.Errorf("can`t open the file for (%v) in (%s)", key, paths[0])
		}
		defer AWFile.Close()
		R4File, err := os.OpenFile(pathtofile, os.O_RDONLY, 0666)
		if err != nil {
			return fmt.Errorf("can`t open the file for (%v) in (%s)", key, paths[0])
		}
		defer R4File.Close()

		_, err = kvf.AppendValues(AWFile, R4File, values)
		if err != nil {
			return fmt.Errorf("can`t update the values in (%s) for (%v)", paths[0], key)
		}

	}

	return nil
}

// TODO Realise
/*func (d *DiskStore[K, V]) List(prefix string) (keys []K, err error) {
	//get path and filename
	paths := strings.Split(prefix, "|")
	pathtofile := d.path + "/" + paths[0]
	if len(paths) != 2 {
		_, err := os.Stat(pathtofile)
		if err != nil {
			if os.IsExist(err) {
				if os.
			}
		}
	}

	//check existing of folders and file
	//GetKeySpace

	return nil, nil
}*/

// TODO Check
func (d *DiskStore) Get(key string) (value string, err error) {
	//get path and filename

	//get filepath and key
	paths := strings.Split(key, "|")
	if len(paths) != 2 {
		return value, fmt.Errorf("the file and path boundaries are not known (%v), use only one '|' symbol for this", key)
	}
	pathtofile := d.path + "/" + paths[0]
	_, err = os.Stat(pathtofile)
	if err != nil {
		if os.IsNotExist(err) {
			return value, fmt.Errorf("can`t open the file for (%v)", key)
		}
	}

	//key is existing?
	found := false
	RFile, err := os.OpenFile(pathtofile, os.O_RDONLY, 0777)
	if err != nil {
		return value, fmt.Errorf("can`t open the file for (%v) in (%s)", key, paths[0])
	}
	defer RFile.Close()

	ks, err := kvf.GetKeySpace(RFile)
	if err != nil {
		return value, fmt.Errorf("can`t get key space")
	}
	for _, k := range ks {
		if k == paths[1] {
			found = true
		}
	}

	//read
	if !found {
		return value, fmt.Errorf("can`t find the (%v) in (%s)", key, paths[0])
	} else {

		//read
		RFile, err := os.OpenFile(pathtofile, os.O_RDONLY, 0777)
		if err != nil {
			return value, fmt.Errorf("can`t open the file for (%v) in (%s)", key, paths[0])
		}
		defer RFile.Close()
		R2File, err := os.OpenFile(pathtofile, os.O_RDONLY, 0777)
		if err != nil {
			return value, fmt.Errorf("can`t open the file for (%v) in (%s)", key, paths[0])
		}
		defer R2File.Close()
		kv, err := kvf.ReadValues(RFile, R2File, []string{string(key)})
		if err != nil {
			return value, fmt.Errorf("can`t read the current values in (%s) for rewriting file with (%v)", paths[0], key)
		}

		_/*keys*/, values := []string{}, []string{}

		for /*k*/_, v := range kv {
			//keys = append(keys, k)
			values = append(values, v)
		}
		value = values[0]

		return value, nil

	}
}

// TODO Realise
func (c *DiskStore) Update(key string, value string) error {
	//get path and filename
	//check existing of folders and file
	//rewrite file or append record to the way and task of the rewriting

	return nil
}

// TODO Realise
func (d *DiskStore) Delete(key string) (value string, err error) {
	//get path and filename
	//check existing of folders and file
	//rewrite the file
	//delete file if need

	return value, nil
}
