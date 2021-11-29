package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		x, ok := m[target-n]
		if ok {
			return []int{i, x}
		}
		m[n] = i
	}
	return []int{}
}

func main() {
	// 1.两数之和
	twoSum([]int{1,2,3,4}, 5)
}
