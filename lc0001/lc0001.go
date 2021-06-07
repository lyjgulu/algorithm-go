package lc0001

func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for index, val := range nums {
		if p, ok := hashMap[target-val]; ok {
			return []int{p, index}
		}
	}
	return nil
}
