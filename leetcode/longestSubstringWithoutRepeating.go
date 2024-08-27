package leetcode

func LongestSubstringWithoutRepeating(s string) int {
	hashMap := make(map[string]int)
	var i int = 0
	var j int = 0

	maxLength := 0

	stringLength := len(s)

	for j < stringLength {
		ch := string(s[j])
		if val, ok := hashMap[ch]; ok {
			i = max(val, i)
		}

		maxLength = max(maxLength, j - i + 1)
		hashMap[ch] = j + 1
		j += 1
	}

	return maxLength
}
