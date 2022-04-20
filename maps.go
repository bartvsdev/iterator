package iterator

// These functions will be replaced by maps.go in Go 1.19

// Keys returns the keys of the map m.
// The keys will be in an indeterminate order.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// Values returns the values of the map m.
// The values will be in an indeterminate order.
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

type Entry[K comparable, V any] struct {
	key   K
	value V
}

// Entries returns the entries of the map m as a slice of structs.
// The entries will be in an indeterminate order.
func Entries[E Entry[K, V], M ~map[K]V, K comparable, V any](m M) []E {
	r := make([]E, 0, len(m))
	for k, v := range m {
		r = append(r, E{k, v})
	}
	return r
}
