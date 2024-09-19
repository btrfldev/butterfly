package store

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/iamsoloma/butterfly/logger"
)

// KV storage based on Badger.
// It may be rewritten to its own Btree engine.
type DiskStore struct {
	path string
	kvs  *badger.DB
}

func NewDiskStore(path string) (diskStore *DiskStore, err error) {
	errLocation := "btrfl.store.DiskStore.NewDiskStore"

	kvs, err := badger.Open(badger.DefaultOptions(path).WithLoggingLevel(badger.ERROR))
	if err != nil {
		return nil, logger.NewErr(errLocation, "Can`t open badger")
	}

	//defer kvs.Close()
	return &DiskStore{
		path: path,
		kvs:  kvs,
	}, nil
}

func (d *DiskStore) CloseDiskStore() {
	d.kvs.Close()
}

func (d *DiskStore) Put(key, value string) error {
	errLocation := "btrfl.store.DiskStore.Put"

	err := d.kvs.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(value))
		return err
	})

	if err != nil {
		return logger.NewErr(errLocation, "Can`t make PUT transaction")
	}

	return nil
}

func (d *DiskStore) Get(key string) (value string, err error) {
	errLocation := "btrfl.store.DiskStore.Get"

	var valCopy []byte
	err = d.kvs.View(func(txn *badger.Txn) error {
		item, e := txn.Get([]byte(key))
		if e != nil {
			return logger.NewErr(errLocation, "Can`t make GET transaction")
		}

		valCopy, err = item.ValueCopy(nil)
		if err != nil {
			return logger.NewErr(errLocation, "Can`t copy value from GET transaction")
		}

		return nil
	})

	value = string(valCopy)
	
	if err != nil {
		return value, logger.NewErr(errLocation, "Can`t make GET transaction")
	} else {
		return value, nil
	}
}

//func (d *DiskStore) Delete(key string) (value string, err)