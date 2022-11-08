package util

import "sort"

type FrequencyCounter[T any] struct {
	Counts     map[string]int64
	References map[string]T
}

type FrequencyPair[T any] struct {
	Count int64
	Value T
}

func (f *FrequencyCounter[T]) Add(key string, value T, weight int64) {
	if f.Counts == nil {
		f.Counts = make(map[string]int64)
	}
	if f.References == nil {
		f.References = make(map[string]T)
	}

	if _, ok := f.Counts[key]; ok {
		f.Counts[key] += weight
	} else {
		f.Counts[key] = weight
		f.References[key] = value
	}
}

func (f *FrequencyCounter[T]) SortK(K int) []FrequencyPair[T] {
	var pairs []FrequencyPair[T]

	for key, count := range f.Counts {
		pairs = append(pairs, FrequencyPair[T]{count, f.References[key]})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Count > pairs[j].Count
	})

	return pairs[:K]
}
