package util

import (
	"math/rand"
	"time"
)

var seedNum = time.Now().UnixNano()

/* *获取[0,max)的随机数
* @param max 最大值
 */
func RandInt(max int) int {
	rand.Seed(seedNum)
	seedNum++
	return rand.Intn(max)
}
