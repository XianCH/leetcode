package leetcode

import (
	"sort"
	"strings"
	"sync"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, str := range strs {
		key := sortGroup(str)
		groups[key] = append(groups[key], str)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

func sortGroup(s string) string {
	// sort the string
	spilt := strings.Split(s, "")
	sort.Strings(spilt)
	return strings.Join(spilt, "")
}

func groupAnagrams2(strs []string) [][]string {
	groups := make(map[string][]string)

	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	for _, str := range strs {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			key := sortGroup(s)
			mutex.Lock()
			groups[key] = append(groups[key], s)
			mutex.Unlock()
		}(str)
	}
	wg.Wait()

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}
