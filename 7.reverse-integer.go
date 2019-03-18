/*
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (25.19%)
 * Total Accepted:    630.1K
 * Total Submissions: 2.5M
 * Testcase Example:  '123'
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 * 
 * Example 1:
 * 
 * 
 * Input: 123
 * Output: 321
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: -123
 * Output: -321
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 120
 * Output: 21
 * 
 * 
 * Note:
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of
 * this problem, assume that your function returns 0 when the reversed integer
 * overflows.
 * 1.注意go语言中的定义， var res int 跟java中的区别，第二次犯错了
 * 2.math.Abs 只针对浮点型参数
 */
func reverse(x int) int {
	var res int;
	for x != 0 {
		if ( res > math.MaxInt32 / 10 || res < math.MinInt32 / 10 ){
			return 0;
		}
		res = res * 10 + x % 10;
		x = x / 10;
	}
	return res;
}
