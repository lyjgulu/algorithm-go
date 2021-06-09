package leetcode

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for index, val := range nums {
		if p, ok := hashTable[target-val]; ok {
			return []int{p, index}
		}
	}
	return nil
}
