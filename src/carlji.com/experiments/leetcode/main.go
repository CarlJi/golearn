package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(removeDuplicates([]int{1, 2, 2, 3}))

	fmt.Println(findSingleNumber([]int{1, 2, 2, 1, 3}))
	fmt.Println(findSingleNumber([]int{1, 2, 2, 1}))
	fmt.Println(findSingleNumber([]int{1, 2, 2}))

	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{9, 9}))
	fmt.Println(plusOne([]int{9, 9, 8}))
	fmt.Println(plusOne([]int{9, 8, 9}))
	fmt.Println(plusOne([]int{0}))

	fmt.Println(reverse("abc"))
	fmt.Println(reverse("abc 中国"))

	fmt.Println(reverseInt(123))
	fmt.Println(reverseInt(-123))
	fmt.Println(reverseInt(-120))
	fmt.Println(1<<31 - 1)
	fmt.Println(-1 << 31)
	fmt.Println(math.MinInt32)
	fmt.Println(math.MaxInt32)

	s := "我是中国人"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}

	fmt.Printf("\n")

	for _, v := range s {
		fmt.Printf("%c", v)
	}

	fmt.Print("\n")

	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar("love"))
	fmt.Println(firstUniqChar("中国中国爱"))

	fmt.Println(maxSubArr([]int{-2, 1}))
	fmt.Println(maxSubArr([]int{-2, 1, -1}))
	fmt.Println(maxSubArr([]int{1, -2, 1, 3, 0, -1, 2}))
	fmt.Println(maxSubArr([]int{1, -2, 1, 3, 0, -1, -2}))

	testMinStack()
	testHanMingWeight()
}

func testHanMingWeight() {
	fmt.Println(hanMingWeight(1))
	fmt.Println(hanMingWeight(2))
	fmt.Println(hanMingWeight(3))
	fmt.Println(hanMingWeight(4))
	fmt.Println(hanMingWeight(5))
	fmt.Println(hanMingWeight(15))
}

func testMinStack() {
	m := Constructor()
	m.Push(1)
	m.Push(2)
	fmt.Printf("%+v\n", m)
	fmt.Println("m.Pop()", m.Top())
	m.Pop()
	m.Push(-1)
	fmt.Println("m.GetMin()", m.GetMin())
	m.Pop()
	m.Pop()
	fmt.Println("m.GetMin()", m.GetMin())
	m.Push(100)
	fmt.Printf("%+v\n", m)
	m.Pop()
	m.Push(-101)
	m.Push(-102)
	m.Push(-103)
	fmt.Printf("%+v\n", m)
	fmt.Println("m.Pop()", m.Top())
	m = Constructor()
	m.Push(0)
	m.Push(1)
	m.Push(0)
	fmt.Printf("%+v\n", m)

	m.GetMin()
	m.Pop()
	fmt.Printf("%+v\n", m)

	fmt.Println("m.GetMin()", m.GetMin())
}

func maxSubArr(nums []int) int {
	var sum, temp = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if temp+nums[i] > nums[i] {
			temp = temp + nums[i]
		} else {
			temp = nums[i]
		}

		if temp > sum {
			sum = temp
		}
	}

	return sum
}

func firstUniqChar(s string) int {
	// 如果字符范围有序，就可以固定申请的数量，比如26个英文字符
	// 256个ascii码。
	// 但是go原因本身是utf-8的，这个范围就广了。
	var cm = make(map[rune]int)
	for _, c := range s {
		if v, ok := cm[c]; ok {
			cm[c] = v + 1
		} else {
			cm[c] = 1
		}
	}

	for i, c := range s {
		if v := cm[c]; v == 1 {
			return i
		}
	}
	return -1
}

func reverseInt(x int) int {
	var neg bool
	var ret int

	if x < 0 {
		neg = true
		x = -x
	}

	for x > 0 {
		tail := x % 10
		temp := ret*10 + tail
		if temp > (2<<31 - 1) {
			return 0
		}
		ret = temp
		x = x / 10
	}

	if neg {
		return -ret
	}

	return ret
}

func reverse(ss string) string {
	var ret = []rune(ss)
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}

	return string(ret)
}

func plusOne(digits []int) []int {
	var l = len(digits) - 1
	if digits[l]+1 < 10 {
		digits[l] = digits[l] + 1
		return digits
	}

	digits[l] = 0
	if l == 0 {
		var newD []int
		newD = append(newD, 1)
		newD = append(newD, 0)
		return newD
	}

	tempRet := plusOne(digits[:l])
	var ret []int
	ret = append(ret, tempRet...) //copy(ret, tempRet) // 为什么这样写不行？
	ret = append(ret, 0)
	return ret
}

func findSingleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	var result = nums[0]
	for _, n := range nums[1:] {
		result ^= n
	}

	return result
}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	i := 0
	j := 1
	for ; j < len(nums); j++ {
		if nums[i] == nums[j] {
			continue
		}

		if j-i == 1 {
			i++
		} else {
			i++
			nums[i] = nums[j]
		}
	}

	return i
}

func hanMingWeight(n int) int {
	count := 0
	for n != 0 {
		if (n & 1) != 0 {
			count++
		}

		n = n >> 1
	}

	return count
}
