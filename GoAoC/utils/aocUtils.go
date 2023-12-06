package utils

import "strings"

func IsAllSpace(s string) bool {
	for _, v := range s {
		if v != ' ' {
			return false
		}

	}
	return true
}

func FromWords(s []string) string {
	var result string
	for _, v := range s {
		result += v + " "
	}
	return result[:len(result)-1]
}

func Words(s string) []string {
	return strings.Split(s, " ")
}

func NotIsAllSpace(s string) bool {
	return !IsAllSpace(s)
}

func MapFromLists[K comparable, V any](a []K, b []V) map[K]V {
	if len(b) < len(a) {
		panic("Not enough values to map keys to")
	}
	res := make(map[K]V)
	for i, v := range a {
		res[v] = b[i]
	}
	return res
}

func Any[K comparable](l []K, fn func(K) bool) bool {
	for _, j := range l {
		if fn(j) {
			return true
		}
	}
	return false
}

func All[K comparable](l []K, fn func(K) bool) bool {
	for _, j := range l {
		if !fn(j) {
			return false
		}
	}
	return true
}

func Map[K any](l []K, fn func(K) K) []K {
	var result []K
	for _, v := range l {
		result = append(result, fn(v))
	}
	return result
}

func Filter[K any](l []K, fn func(K) bool) []K {
	var result []K
	for _, v := range l {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
