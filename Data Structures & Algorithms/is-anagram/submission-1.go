func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    chars := make(map[rune]int)

    for _, val := range []rune(s) {
        chars[val] += 1
    }

    for _, val := range []rune(t) {
        chars[val] -= 1
    }

    for _, val := range chars {
        if val != 0 {
            return false
        }
    }

    return true
}
