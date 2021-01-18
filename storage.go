package main

import (
	"sort"
	"strings"
	"sync"
)

var (
	vocab   = make(map[string][]string)
	checker = make(map[string]struct{})
	lock    = sync.RWMutex{}
)

func getAnnagrams(word string) []string {
	lock.RLock()
	defer lock.RUnlock()
	hash := getHash(word)
	val, ok := vocab[hash]
	if !ok {
		return nil
	}
	return val
}

func loadAnnagrams(words []string) {
	lock.Lock()
	defer lock.Unlock()
	for _, w := range words {
		if _, ok := checker[w]; ok {
			continue
		}
		checker[w] = struct{}{}
		lowercaseWord := strings.ToLower(w)
		hash := getHash(lowercaseWord)
		// fmt.Printf("Word: %s, Hash: %s\n", w, hash)
		vocab[hash] = append(vocab[hash], w)
	}
}

func getHash(word string) string {
	arr := []rune(word)
	// fmt.Printf("До хэша: %v\n", arr)
	sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })
	// fmt.Printf("После хэша: %v\n", arr)
	hash := string(arr)
	return hash
}
