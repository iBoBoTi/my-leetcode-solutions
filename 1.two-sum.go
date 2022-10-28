package main
/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
// subtract the


func twoSum(nums []int, target int) []int {
	checker := make(map[int]int)
	for i, _ := range nums {
		comp := target - nums[i]
		if _, exist := checker[comp]; exist{
			return []int{checker[comp],i}
		}
		checker[nums[i]] = i
	}
	return []int{}
}

// @lc code=end

