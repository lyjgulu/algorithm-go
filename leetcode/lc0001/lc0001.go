package leetcode

func twoSum(num []int, target int) []int {
	midMap := map[int]int{}
	for index, value := range num {
		if v, ok := midMap[target-value]; ok {
			return []int{v, index}
		}
		midMap[value] = index
	}
	return nil
}
