package leetcode

func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for index, val := range nums {
		if p, ok := hashMap[target-val]; ok {
			return []int{p, index}
		}
		hashMap[val] = index
	}
	return nil
}
