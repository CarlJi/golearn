package strs

// FiledFunc方法会根绝func f来split字符串 s, 而这个判断函数f可以由外部指定
// 其目的是可以让外部自由的指定分割条件,比如指定多个分隔符,来分割字符串.
// 如果s为空,或者s中所有的code point都满足函数f(也就是是f(rune in s)==true), 则返回一个空slice
// 算法实现原理:
//    1. 要实现此算法,我们可以先计算出最后将返回的slice,包含多少个元素,这样我们就可以初始化响应长度的slice,然后遍历字符串s,来一一填充
//       要返回的slice
//    2. 而要计算有多少元素会返回(源码里使用单词Field表示字段,一个字段就是一个元素), 方法是, 定义两个指针,一个叫wasInField, 表示前一个字符
//       是否在Field中,一个叫InField, 表示当前字符是否是在Field. 这样就可以遍历字符串s,如果某一个字符不是分隔符(InField=true), 且其
//       前一个字符是分隔符(wasInField = false), 那么久表明这个是一个新的元素了, 那么最后返回的元素长度就得++了,所以由此方法,我们就可以计算出
//       要最后返回的切片长度
//    3. 有了切片长度,接下来我们就得来填充这个切片. 同样遍历字符串s, 要获得一个切片元素, 如果我们知道元素的开始位置和结束位置,那么就可以直接获得
//       oneElement := s[start:end]
//       那么如果获得开始位置和结束位置呢? 参考步骤2不难想出, 如果一个字符不在Filed中(也就是字符为分隔符,f(rune)==true), 这个字符的位置就是
//       前一个字符的结束位置. 而我们只需要一个指针来标记开始位置就好了.
func FieldsFunc(s string, f func(rune) bool) []string {

	// 计算最后分割的slice长度
	count := 0
	inFiled := false
	for _, char := range s {
		wasInFiled := inFiled
		inFiled = !f(char)
		if inFiled && !wasInFiled {
			count++
		}
	}

	// 初始化,并填充这个slice
	res := make([]string, count)
	i := 0
	start := -1 // start 是切片元素在字符串s中的开始位置, -1表明不存在
	for index, char := range s {
		if f(char) {
			if start >= 0 {
				res[i] = s[start:index]
			}
		}
		start = index
	}

	if i == 0 && start == len(s) {
		res[i] = s
	}

	return res
}
