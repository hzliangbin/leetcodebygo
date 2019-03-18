/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 *
 * https://leetcode.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (42.27%)
 * Total Accepted:    527.9K
 * Total Submissions: 1.2M
 * Testcase Example:  '121'
 *
 * Determine whether an integer is a palindrome. An integer is a palindrome
 * when it reads the same backward as forward.
 * 
 * Example 1:
 * 
 * 
 * Input: 121
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it
 * becomes 121-. Therefore it is not a palindrome.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a
 * palindrome.
 * 
 * 
 * Follow up:
 * 
 * Coud you solve it without converting the integer to a string?
 * 1，沿用翻转整数的思路，翻转后相等即回文。
 * 2, input没有溢出，若回文，则翻转过程必不溢出，若溢出，则不回文。
 */
func isPalindrome(x int) bool {
	if x < 0 {
		return false;
	}
	a := x;
	var res int;
	for x != 0 {
		if ( res > math.MaxInt32 / 10 || res < math.MinInt32 / 10) {
			return false;
		}
		res = res * 10 + x % 10;
		x = x / 10;
	}  
	return a == res;
}

