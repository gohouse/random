package random

import (
	"math/rand"
	"time"
)

// RandType ...
type RandType int

const (
	// 大写字母
	CAPITAL RandType = iota + 1
	// 小写字母
	LOWERCASE
	// 数字
	NUMBERIC
	// 小写字母+数字
	LOWERCASE_NUMBERIC
	// 大写字母+数字
	CAPITAL_NUMBERIC
	// 大写字母+小写字母
	CAPITAL_LOWERCASE
	// 数字+字母
	ALL
)

var (
	numberic  = `0123456789`
	lowercase = `abcdefghijklmnopqrstuvwxyz`
	capital   = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Rand 随机生成6-32位的随机字符串(长度类型皆随机)
func Rand() string {
	var rt = getRandType()
	var length = randBetween(6, 32)
	return random(length, rt)
}

// Random 随机生成指定长度的随机字符串(类型可选或随机)
func Random(length int, fill ...RandType) string {
	var rt RandType
	if len(fill) > 0 {
		rt = fill[0]
	} else {
		rt = getRandType()
	}
	return random(length, rt)
}

// RandomBetween 随机生成指定长度区间的随机字符串(类型可选或随机)
func RandomBetween(min, max int, fill ...RandType) string {
	var rt RandType
	if len(fill) > 0 {
		rt = fill[0]
	} else {
		rt = getRandType()
	}
	var length = randBetween(min, max)
	return random(length, rt)
}

func random(length int, fill RandType) string {
	str := getFillStr(fill)
	var res string
	var i = length
	for i > 0 {
		res += getCharactorFromStr(str)
		i--
	}
	return res
}
