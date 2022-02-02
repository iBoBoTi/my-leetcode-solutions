package main
/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
// subtract the


func twoSum(nums []int, target int) []int {
	ans := make([]int,0)
	for i,j:= range nums{
		num:= target - j
		ind:= contains(nums,num, i)
		if ind != -1 {
			ans = append(ans,i,ind)
			break
		}
	}
    return ans
}

func contains(nums []int, num, ind int)int{
	n:=-1
	for i,j:= range nums{
		if j == num && ind != i{
		 n= i 	
		}
	}
	return n
}
// @lc code=end

