/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */
package main

import (
	"strconv"
	"math"
)
// @lc code=start
// convert the integer to string
// convert the string to a slice of rune 
// loop through the slice of rune
	// reverse the rune values
// convert the rune back to string then back to integer
// check if the intial integer is negative and return a negative reversed integer
// else return the reversed integer
func reverse(x int) int {
	strEquivalent:=strconv.Itoa(int(math.Abs(float64(x))))
	runeEquivalent:=[]byte(strEquivalent)

	for i,j:=0,len(runeEquivalent)-1;i<j;i,j=i+1,j-1{
		runeEquivalent[i],runeEquivalent[j] = runeEquivalent[j],runeEquivalent[i]
	}

	finalAns,_:=strconv.Atoi(string(runeEquivalent))
	if finalAns > math.MaxInt32 || finalAns < math.MinInt32{
		return 0
	}
	if math.Signbit(float64(x)){
		return -finalAns
	}
	return finalAns
	

}
// @lc code=end

