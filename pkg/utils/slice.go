package utils

import (
	"strings"

	"github.com/spf13/cast"
)

type Slice[T any] []T

type CompareSlice[T comparable] []T

var (
	_ = func(i int) bool {
		if i == 0 {
			return false
		}
		return true
	}
)

func (s Slice[T]) ToString() string {
	str := make([]string, 0)
	for _, t := range s {
		str = append(str, cast.ToString(t))
	}
	return strings.Join(str, ",")
}

func (s Slice[T]) ToStringSlice() Slice[string] {
	strArr := make([]string, 0)
	for _, t := range s {
		strArr = append(strArr, cast.ToString(t))
	}
	return strArr
}

func (s Slice[T]) Filter(f func(T) bool) Slice[T] {
	tArr := make([]T, 0)
	for _, t := range s {
		if f(t) {
			tArr = append(tArr, t)
		}
	}
	return tArr
}

func (s CompareSlice[T]) ToMapVStruct() Map[T, struct{}] {
	m := make(map[T]struct{})
	for _, t := range s {
		m[t] = struct{}{}
	}
	return m
}

func (s Slice[T]) ElseEmpty() Slice[T] {
	if s == nil {
		return Slice[T]{}
	}
	return s
}

type TwoSlice[T any] [][]T

func CanSafeVisitTwoArray[T any](array TwoSlice[T], index int) bool {
	return len(array) > index
}

func CanSafeVisitArray[T any](array Slice[T], index int) bool {
	return len(array) > index
}
