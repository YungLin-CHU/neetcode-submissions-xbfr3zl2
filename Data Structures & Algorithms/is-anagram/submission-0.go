func isAnagram(s string, t string) bool {
    chars := make(map[rune]int)

    for _, val := range []rune(s) {
        chars[val] += 1
    }

    for _, val2 := range []rune(t) {
        chars[val2] -= 1
    }

    for _, v := range chars {
        if v != 0 {
            return false
        }
    }

    return true
}
