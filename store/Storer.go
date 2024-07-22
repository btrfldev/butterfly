package store

//  |   K - key  |  С - comparable value (for search)   |   V - value   |   S - search   |
type Storer[K, C comparable, V any, S func(K, C) bool] interface {
	Put(K, V) error
	Get(K) (V, error)
	Update(K, V) error
	Delete(K) (V, error)
	List(S, C) ([]K, error)
}
