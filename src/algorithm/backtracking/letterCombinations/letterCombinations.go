package main

var dsMap map[string][]string = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

var res []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res = []string{}
	backtrack(0, digits, "")
	return res
}

func backtrack(idx int, digits string, str string) {
	if idx == len(digits) {
		res = append(res, str)
		return
	}
	choose := dsMap[string(digits[idx])]
	for _, tmp := range choose {
		backtrack(idx+1, digits, str+tmp)
	}
}

func main() {
	// [17. 电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)
}
