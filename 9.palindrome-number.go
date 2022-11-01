/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
import "strconv"

func isPalindrome(x int) bool {

	// stringify number
	// loop through the []rune of the stringified number and check if the beginning isn't equal to the endings
	stringifiedNum := []rune(strconv.Itoa(x))
	for i,j:=0,len(stringifiedNum)-1;i<j;i,j=i+1,j-1{
		if stringifiedNum[i] != stringifiedNum[j]{
			return false
		}
	}
	return true
    
}
// @lc code=end

