package mapUtil

// 聚合
func Collect[K comparable, T any](m map[K][]T, choose func(x1, x2 T) T) map[K]T {
	ret := make(map[K]T)
	if len(m) == 0 {
		return ret
	}
	for k, collection := range m {
		var initial T
		if len(collection) == 0 {
			ret[k] = initial
			continue
		}
		if len(collection) == 1 {
			ret[k] = collection[0]
			continue
		}
		initial = collection[0]
		for i := range collection {
			if i == 0 {
				continue
			}
			initial = choose(initial, collection[i])
		}
		ret[k] = initial
	}
	return ret
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))

	var i int
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))

	var i int
	for _, v := range m {
		values[i] = v
		i++
	}

	return values
}

func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	size := 0
	for i := range maps {
		size += len(maps[i])
	}

	result := make(map[K]V, size)

	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}

func ForEach[K comparable, V any](m map[K]V, iteratee func(key K, value V)) {
	for k, v := range m {
		iteratee(k, v)
	}
}

func Filter[K comparable, V any](m map[K]V, predicate func(key K, value V) bool) map[K]V {
	result := make(map[K]V)

	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}
	return result
}

func FilterByKeys[K comparable, V any](m map[K]V, keys []K) map[K]V {
	result := make(map[K]V)

	for k, v := range m {
		if Contain(keys, k) {
			result[k] = v
		}
	}
	return result
}

func Contain[T comparable](slice []T, target T) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}

	return false
}

func Minus[K comparable, V any](mapA, mapB map[K]V) map[K]V {
	result := make(map[K]V)

	for k, v := range mapA {
		if _, ok := mapB[k]; !ok {
			result[k] = v
		}
	}
	return result
}

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

func Entries[K comparable, V any](m map[K]V) []Entry[K, V] {
	entries := make([]Entry[K, V], 0, len(m))

	for k, v := range m {
		entries = append(entries, Entry[K, V]{
			Key:   k,
			Value: v,
		})
	}

	return entries
}

func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V {
	result := make(map[K]V, len(entries))

	for _, v := range entries {
		result[v.Key] = v.Value
	}

	return result
}

func MapKeys[K comparable, V any, T comparable](m map[K]V, iteratee func(key K, value V) T) map[T]V {
	result := make(map[T]V, len(m))

	for k, v := range m {
		result[iteratee(k, v)] = v
	}

	return result
}

func MapValues[K comparable, V any, T any](m map[K]V, iteratee func(key K, value V) T) map[K]T {
	result := make(map[K]T, len(m))

	for k, v := range m {
		result[k] = iteratee(k, v)
	}

	return result
}

func HasKey[K comparable, V any](m map[K]V, key K) bool {
	_, haskey := m[key]
	return haskey
}

func HasValue[K comparable, V comparable](m map[K]V, value V) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

func GetValue[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	if v, haskey := m[key]; haskey {
		return v
	}
	return defaultValue
}
