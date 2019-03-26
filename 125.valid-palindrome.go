/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 *
 * https://leetcode.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (30.46%)
 * Total Accepted:    331.8K
 * Total Submissions: 1.1M
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * Given a string, determine if it is a palindrome, considering only
 * alphanumeric characters and ignoring cases.
 * 
 * Note: For the purpose of this problem, we define empty string as valid
 * palindrome.
 * 
 * Example 1:
 * 
 * 
 * Input: "A man, a plan, a canal: Panama"
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "race a car"
 * Output: false
 * 
 * 
 */
func isPalindrome(s string) bool {
	t := strings.ToLower(s);
	x, y := 0, len(t) - 1;
	for x < y {
		for (x < y) && !((96 < t[x] && t[x] < 123) || (47 < t[x] && t[x] < 58)){
			x++;
		}
		for (x < y) && !((96 < t[y] && t[y] < 123) || (47 < t[y] && t[y] < 58)){
			y--;
		}
		if (t[x] != t[y]) {
			return false;
		} else {
			x++;
			y--;
		}
	}
	return true;
}

