/*
 * @lc app=leetcode id=137 lang=golang
 *
 * [137] Single Number II
 *
 * https://leetcode.com/problems/single-number-ii/description/
 *
 * algorithms
 * Medium (45.30%)
 * Total Accepted:    159.8K
 * Total Submissions: 352.6K
 * Testcase Example:  '[2,2,3,2]'
 *
 * Given a non-empty array of integers, every element appears three times
 * except for one, which appears exactly once. Find that single one.
 * 
 * Note:
 * 
 * Your algorithm should have a linear runtime complexity. Could you implement
 * it without using extra memory?
 * 
 * Example 1:
 * 
 * 
 * Input: [2,2,3,2]
 * Output: 3
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [0,1,0,1,0,1,99]
 * Output: 99
 * 
 */
func singleNumber(nums []int) int {
	a,b := 0,0;
	for i, num := range nums {
		b = (b ^ num) & ^a;
		a = (a ^ nums[i]) & ^b;
	}
	return b;
}

