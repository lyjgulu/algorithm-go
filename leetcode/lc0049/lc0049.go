package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedString := string(s)
		mp[sortedString] = append(mp[sortedString], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, strings := range mp {
		ans = append(ans, strings)
	}
	return ans
}
