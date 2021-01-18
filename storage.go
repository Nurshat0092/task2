package main

import (
	"sort"
	"sync"
)

var (
	vocab = make(map[string][]string)
	lock  = sync.RWMutex{}
)

func getAnnagrams(word string) []string {
	lock.RLock()
	defer lock.RUnlock()
	hash := getHash(word)

	return vocab[hash]
}

func loadAnnagrams(words []string) {
	lock.Lock()
	defer lock.Unlock()
	for _, w := range words {
		hash := getHash(w)
		vocab[hash] = append(vocab[hash], w)
	}
}

func getHash(word string) string {
	arr := []rune(word)
	sort.SliceStable(arr, func(i, j int) bool { return i < j })
	hash := string(arr)
	return hash
}
