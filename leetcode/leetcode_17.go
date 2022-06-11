package leetcode

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	var m = map[int]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}

	var dfs func(i int)
	var res []string
	var curr []byte
	dfs = func(i int) {
		if i == len(digits) {
			res = append(res, string(curr))
			return
		}

		str := m[int(digits[i]-'0')]
		for j := range str {
			curr = append(curr, str[j])
			dfs(i + 1)
			curr = curr[:len(curr)-1]
		}
	}
	dfs(0)

	return res
}
