/*
 * [1] Two Sum
 *
 * https://leetcode.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (42.40%)
 * Total Accepted:    1.6M
 * Total Submissions: 3.7M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers, return indices of the two numbers such that they
 * add up to a specific target.
 * 
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 * 
 * Example:
 * 
 * 
 * Given nums = [2, 7, 11, 15], target = 9,
 * 
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 * 注意点：
 * 1，定义变量不用会报错
 * 2，for range和 java中for及for each的区别，for range同时拿到下标和对应值信息，有用
 * 3，= 和 := 的区别，注意go语言中声明定义变量的方式和java中的区别
 */
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums {
		val2 := target - val;
		index2, ok := m[val2];
		if ok && index2 != index {
			return []int{index2,index};
		}	
		m[val]=index
	}
	return []int{-1,-1};
}
