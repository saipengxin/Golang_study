package split

import "strings"

// Split 实现了字符串s根据 sep进行分割
func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)
	// 最开始的代码是append的时候才会动态的增加空间，现在我们改成初始就手动分配好空间
	// 容量是分割的字符串数量加1，例如a:b:c按照:分割，:有两个，我们分割完成有3份
	// strings.Count 返回字符串s中有几个不重复的sep子串。
	result = make([]string, 0, strings.Count(s, sep)+1)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
