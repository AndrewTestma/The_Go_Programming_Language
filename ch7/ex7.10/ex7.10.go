package ex7_10

import "sort"

// 练习 7.10： sort.Interface类型也可以适用在其它地方。编写一个IsPalindrome(ssort.Interface) bool函数表明序列s是否是回文序列，换句话说反向排序不会改变这个序列。假设如果!s.Less(i, j) && !s.Less(j, i)则索引i和j上的元素相等。
func equal(i, j int, s sort.Interface) bool {
	return !s.Less(i, j) && s.Less(j, i)
}
func IsPalindrome(s sort.Interface) bool {
	max := s.Len() - 1
	for i := 0; i < s.Len()/2; i++ {
		if !equal(i, max-i, s) {
			return false
		}
	}
	return true
}
