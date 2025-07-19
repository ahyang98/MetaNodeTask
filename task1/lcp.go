package task1

import "strings"

func LongestCommonPrefix(strs []string) string {

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		prefix = longestCommonPrefix(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}

	return prefix
}

func longestCommonPrefix(str1 string, str2 string) string {
	prefix := strings.Builder{}
	for i := 0; i < len(str1) && i < len(str2); i++ {
		if str1[i] != str2[i] {
			break
		}
		prefix.WriteByte(str1[i])
	}
	return prefix.String()
}
