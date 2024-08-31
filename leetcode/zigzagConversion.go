package leetcode

func ZigZagConversion(s string, numRows int) string {
	length_of_string := len(s)

	if numRows == 1 {
		return s
	}

	number_of_loops := (length_of_string / numRows) + 1

	substring := ""

	for i := 0; i < numRows; i++ {

		for j := 0; j < number_of_loops; j++ {
			if i == 0 {
				index := j * 2 * (numRows - 1)
				if index < length_of_string {
					substring += string(s[index])
				}
			} else if i == numRows-1 {
				index := j * 2 * (numRows - 1)
				subIndex := index + numRows - 1
				if subIndex < length_of_string {
					substring += string(s[subIndex])
				}
			} else {
				subIndex1 := j*2*(numRows-1) + i
				subIndex2 := j*2*(numRows-1) + (2 * (numRows - 1)) - i
				if subIndex1 < length_of_string {
					substring += string(s[subIndex1])
				}
				if subIndex2 < length_of_string {
					substring += string(s[subIndex2])
				}
			}
		}
	}

	return substring
}