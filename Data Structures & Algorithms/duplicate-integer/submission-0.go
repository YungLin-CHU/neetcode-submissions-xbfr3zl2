func hasDuplicate(nums []int) bool {
    m := make(map[int]int)
    for _, val := range nums {
        if _, ok := m[val]; ok {
            return true
        }
        m[val] = 1
    }
    return false
}
