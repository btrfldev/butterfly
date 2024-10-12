package store

import (
	"github.com/btrfldev/butterfly/logger"
	"github.com/dgraph-io/badger/v4"
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

func (d *DiskStore) Put(kv map[string]string) error {
	errLocation := "btrfl.store.DiskStore.Put"

	err := d.kvs.Update(func(txn *badger.Txn) error {
		for k, v := range kv {
			err := txn.Set([]byte(k), []byte(v))
			if err != nil {
				return logger.NewErr(errLocation, "Can`t set "+k)
			}
		}
		return nil
	})

	if err != nil {
		return logger.NewErr(errLocation, "Can`t make PUT transaction")
	}

	return nil
}

func (d *DiskStore) Get(keys []string) (kv map[string]string, err error) {
	errLocation := "btrfl.store.DiskStore.Get"

	kv = make(map[string]string)
	err = d.kvs.View(func(txn *badger.Txn) error {
		for _, k := range keys {
			item, err := txn.Get([]byte(k))
			if err == badger.ErrKeyNotFound {
				return logger.NewErr(errLocation, "the key ("+k+") does not exists")
			} else if err != nil {
				return err
			}

			var valCopy []byte
			valCopy, err = item.ValueCopy(nil)

			if err != nil {
				return logger.NewErr(errLocation, "Can`t copy value from GET query for "+k)
			}
			kv[k] = string(valCopy)
		}

		return nil
	})

	if err != nil {
		return kv, logger.NewErr(errLocation, "Can`t make GET transaction")
	} else {
		return kv, nil
	}
}

func (d *DiskStore) List(search func(k string, c string) bool, comp string) (keys []string, err error) {
	errLocation := "btrfl.store.DiskStore.Get"

	err = d.kvs.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			key := string(item.KeyCopy(nil))
			if search(key, comp) {
				keys = append(keys, key)
			}
		}
		return nil
	})
	if err != nil {
		return keys, logger.NewErr(errLocation, "Can`t make LIST iteration")
	} else {
		return keys, nil
	}
}

func (d *DiskStore) Delete(keys []string) (kv map[string]string, err error) {
	errLocation := "btrfl.store.DiskStore.Delete"

	kv = make(map[string]string)
	err = d.kvs.Update(func(txn *badger.Txn) error {
		for _, k := range keys {
			item, err := txn.Get([]byte(k))
			if err == badger.ErrKeyNotFound {
				return logger.NewErr(errLocation, "the key ("+k+") does not exists")
			} else if err != nil {
				return logger.NewErr(errLocation, "Can`t get "+k)
			}

			var valCopy []byte
			valCopy, err = item.ValueCopy(nil)

			if err != nil {
				return logger.NewErr(errLocation, "Can`t copy value from GET query for "+k)
			}
			kv[k] = string(valCopy)

			err = txn.Delete([]byte(k))
			if err != nil {
				return logger.NewErr(errLocation, "Can`t delete "+k)
			}
		}

		return nil
	})

	if err != nil {
		return kv, logger.NewErr(errLocation, "Can`t make DELETE transaction")
	} else {
		return kv, nil
	}

}
