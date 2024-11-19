package utils

import (
	"golang.org/x/exp/maps"
)

type Map[K comparable, V any] map[K]V

func ExistList[M ~map[K]V, K comparable, V any](m M, t []K) []K {
	anies := make([]K, 0)
	for _, a := range t {
		if _, ok := m[a]; ok {
			anies = append(anies, a)
		}
	}
	return anies
}

func NotExistList[M ~map[K]V, K comparable, V any](m M, t []K) []K {
	anies := make([]K, 0)
	for _, a := range t {
		if _, ok := m[a]; !ok {
			anies = append(anies, a)
		}
	}
	return anies
}

func (m Map[K, V]) Keys() []K {
	return maps.Keys(m)
}

func (m Map[K, V]) Values() []V {
	return maps.Values(m)
}
