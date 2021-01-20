package storage

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

// GetAnnagrams gets all annagrams of word in storage
func GetAnnagrams(word string) []string {
	lock.RLock()
	defer lock.RUnlock()
	hash := getWordHash(word)
	val, ok := vocab[hash]
	if !ok {
		return nil
	}
	return val
}

// LoadAnnagrams loads all words to storage
func LoadAnnagrams(words []string) {
	lock.Lock()
	defer lock.Unlock()
	for _, w := range words {
		if _, ok := checker[w]; ok {
			continue
		}
		checker[w] = struct{}{}
		lowercaseWord := strings.ToLower(w)
		hash := getWordHash(lowercaseWord)
		vocab[hash] = append(vocab[hash], w)
	}
}

// returns hash of given word
// calculates hash by coverting word to slice of runes, and sorting it.
func getWordHash(word string) string {
	arr := []rune(word)
	sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })
	hash := string(arr)
	return hash
}
