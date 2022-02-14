/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
package main
func removeDuplicates(nums []int) int {
	finalAns := make([]int,0)
	checker := make(map[int]int)
    for _,j := range nums {
		_,ok:=checker[j]
		if !ok {
			checker[j] = 1
			finalAns = append(finalAns, j)
		} 
		checker[j]++
	}
	for i:=0; i<len(finalAns); i++{
		nums[i] = finalAns[i]
	} 
	return len(finalAns)
}
// @lc code=end

