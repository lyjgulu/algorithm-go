package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hashTable := map[rune]int{}
	for _, ch := range s {
		hashTable[ch]++
	}
	for _, ch := range t {
		hashTable[ch]--
		if hashTable[ch] < 0 {
			return false
		}
	}
}
