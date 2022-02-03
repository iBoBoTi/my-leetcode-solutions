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
func reverse(x int) int {
	if int32(x) == 0{
		return 0
	}
	ans:=strconv.Itoa(int(math.Abs(float64(x))))
	anss:=[]rune(ans)
	for i,j:=0,len(anss)-1;i<j;i,j=i+1,j-1{
		anss[i],anss[j] = anss[j],anss[i]
	}
	finalAns,_:=strconv.Atoi(string(anss))
	if math.Signbit(float64(x)){
		return -finalAns
	}
	return finalAns
	

}
// @lc code=end

