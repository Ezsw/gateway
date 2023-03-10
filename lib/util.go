package lib

// InArrayString 检查string数组arr中是否有字符串s
func InArrayString(s string, arr []string) bool {
	for _, i := range arr {
		if i == s {
			return true
		}
	}
	return false
}
