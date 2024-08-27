package main

import (
	"dsa/leetcode"
	"fmt"
)

func main() {
	// 1. Two Sum
	fmt.Println("================================================")
	fmt.Println("1. Two Sum")

	arr := []int{1, 2, 3, 4}
	target := 5
	result := leetcode.TwoSum(arr, target)
	fmt.Println("Index: ", result)
	fmt.Println("================================================")

	// 2. Add Two Numbers
	fmt.Println("================================================")
	fmt.Println("2. Add Two Numbers")

	l1 := &leetcode.ListNode{Val: 2, Next: &leetcode.ListNode{Val: 4, Next: &leetcode.ListNode{Val: 3,Next: nil}}}
	l2 := &leetcode.ListNode{Val: 5, Next: &leetcode.ListNode{Val: 6, Next: &leetcode.ListNode{Val: 4, Next: nil}}}

	sum := leetcode.AddTwoNumbers(l1, l2)
	for sum.Next != nil {
		fmt.Print(sum.Val, " ")
		sum = sum.Next
	}
	fmt.Print(sum.Val, "\n")
	fmt.Println("================================================")

	// 3. Longest Substring without Repeating Characters
	fmt.Println("================================================")
	fmt.Println("3. Longest Substring without Repeating Characters")
	st := "bbbbb"
	length := leetcode.LongestSubstringWithoutRepeating(st)
	fmt.Println("Length of longest substring:", length)
	fmt.Println("================================================")

	// 4. Median of Two Sorted Arrays
	fmt.Println("================================================")
	fmt.Println("4. Median of Two Sorted Arrays")
	nums1 := []int{1,3}
	nums2 := []int{2,4}
	nu := leetcode.MedianOfTwoSortedArrays(nums1, nums2)
	fmt.Println("4. Median:", nu)
	fmt.Println("================================================")

	// 5. Longest Palindrome Substring
	fmt.Println("================================================")
	fmt.Println("5. Longest Palindrome Substring")
	// str := "babad"
	// str := "cbbd"
	// str := "bb"
	str5 := "aaaa"
	substr5 := leetcode.LongestPalindromeSubstring(str5)
	fmt.Println("Palindrome:", substr5)
	fmt.Println("================================================")

	// 6. ZigZag Conversion
	fmt.Println("================================================")
	fmt.Println("6. ZigZag Conversion")
	str6 := "PAYPALISHIRING"
	rows6 := 4
	newStr6 := leetcode.ZigZagConversion(str6, rows6)
	fmt.Println("Palindrome:", newStr6)
	fmt.Println("================================================")
}
