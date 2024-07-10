package store

// FileSystem store based on .btrfl
type GraphiteStore[K comparable, V any] struct {
}

func NewGraphiteStore[K comparable, V any]() *GraphiteStore[K, V] {
	return &GraphiteStore[K, V]{}
}

// TODO Realise
func (g *GraphiteStore[K, V]) Put(key K, value V) error {
	//get path and filename
	//create folders if not exist
	//create file if not exist
	//write kv

	return nil
}

// TODO Reamade this, need prefix.
// TODO Realise
func (g *GraphiteStore[K, V]) List(key K, value V) (keys []K, err error) {
	//get path and filename
	//check existing of folders and file
	//GetKeySpace

	return nil, nil
}

// TODO Realise
func (g *GraphiteStore[K, V]) Get(key K) (value V, err error) {
	//get path and filename
	//check existing of folders and file
	//Read kv

	return value, nil
}

// TODO Realise
func (c *GraphiteStore[K, V]) Update(key K, value V) error {
	//get path and filename
	//check existing of folders and file
	//rewrite file or append record to the way and task of the rewriting

	return nil
}

// TODO Realise
func (c *GraphiteStore[K, V]) Delete(key K) (value V, err error) {
	//get path and filename
	//check existing of folders and file
	//rewrite the file
	//delete file if need

	return value, nil
}