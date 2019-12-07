package random

import (
	"fmt"
	"math/rand"
)

func getRandType() RandType {
	var ct = []RandType{
		// 大写字母
		CAPITAL,
		// 小写字母
		LOWERCASE,
		// 数字
		NUMBERIC,
		// 小写字母+数字
		LOWERCASE_NUMBERIC,
		// 大写字母+数字
		CAPITAL_NUMBERIC,
		// 大写字母+小写字母
		CAPITAL_LOWERCASE,
		// 数字+字母
		ALL,
	}
	l := len(ct)
	i := rand.Intn(l)
	return ct[i]
}

func getCharactorFromStr(str string) string {
	strLen := len(str)
	return string(([]rune(str))[rand.Intn(strLen-1)])
}

func getFillStr(rt RandType) string {
	switch rt {
	case CAPITAL:
		return capital
	case LOWERCASE:
		return lowercase
	case NUMBERIC:
		return numberic
	case CAPITAL_LOWERCASE:
		return fmt.Sprint(capital, lowercase)
	case CAPITAL_NUMBERIC:
		return fmt.Sprint(capital, numberic)
	case LOWERCASE_NUMBERIC:
		return fmt.Sprint(lowercase, numberic)
	case ALL:
		return fmt.Sprint(capital, lowercase, numberic)
	}
	return ""
}

func randBetween(min, max int) int {
	return rand.Intn(max-min+1) + min
}
