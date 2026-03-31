func twoSum(nums []int, target int) []int {
    num_map := make(map[int]int)
    for idx, val := range nums {
        if i, ok := num_map[target - val]; ok {
            return []int{i, idx} 
        }
        num_map[val] = idx
    }

    return []int{}
}
