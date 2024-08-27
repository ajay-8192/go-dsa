package leetcode

func TwoSum(nums []int, target int)	[]int {
	hashMap := make(map[int]int)
    for i, val := range nums {
        cont := target - val
        if idx, ok := hashMap[cont]; ok {
            return []int{idx, i}
        }
        hashMap[val] = i
    }

    return []int{}
}
