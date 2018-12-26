package main

import (
	"fmt"
	"math"
	"sort"

	. "carlji.com/experiments/leetcode/sub"
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

	//TopK related
	testTopKFrequency()
	testPartition()
	testFindKLargest()
	testFrequencySort()

	testT()

	testBinarySearch()
}

func testT() {
	var times = []struct {
		start int64
		end   int64
	}{
		{100100, 100120},
		{100100, 100140},
		{100130, 100135},
		{100134, 100135},
		{100134, 100135},
		{100140, 100155},
		{100141, 100155},
	}

	var pp int
	var endT []int64

	for i, tt := range times {
		if i == 0 {
			endT = append(endT, tt.end)
			pp += 1
			continue
		}

		if tt.start < endT[0] {
			pp += 1
			endT = InsertSort(endT, tt.end)

			fmt.Printf("1=========> %+v \n", endT)
		} else {
			endT[0] = tt.end
			sort.SliceStable(endT, func(i, j int) bool {
				return i < j
			})

			fmt.Printf("2=========> %+v \n", endT)

		}
	}

	fmt.Println("=========>", pp)
	fmt.Printf("=========> %+v", endT)
}

func InsertSort(arr []int64, element int64) []int64 {
	for i, e := range arr {
		if element > e {
			continue
		}
		newArr := make([]int64, len(arr)+1)
		copy(newArr[0:i], arr[0:i])
		newArr[i] = element
		copy(newArr[i+1:], arr[i:])

		return newArr
	}

	arr = append(arr, element)
	return arr
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

func testTopKFrequency() {
	fmt.Println(topKFrequeny([]int{1}, 1))
	fmt.Println(topKFrequeny([]int{1, 1, 1, 1, 2, 2, 0}, 2))
	fmt.Println(topKFrequeny([]int{1, 1, 1, 2, 2, 3}, 2))
}

func testPartition() {
	fmt.Println(partition([]int{1}, 0, 0))
	fmt.Println(partition([]int{2, 5, 3, 4, 2, 4, 1}, 0, 6))
}

func testFindKLargest() {
	fmt.Println("func testFindKLargest():")
	fmt.Println(findKthLargest([]int{1}, 1))
	fmt.Println(findKthLargest([]int{2, 5, 3, 4, 2, 4, 1}, 2))
	fmt.Println(findKthLargest([]int{2, 1}, 2))
}

func testFrequencySort() {
	fmt.Println("func testFrequencySort():")
	fmt.Println(frequencySort("eeeeee"))
	fmt.Println(frequencySort("eeeetessf"))
}

func testBinarySearch() {
	fmt.Println("func testBinarySearch():")
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 8}, 6))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 8}, 1))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 6, 8}, 7))
}

//思路: 按照topKFrequency来做，先统计各个rune出现的次数，然后基于次数来重新组合rune
func frequencySort(s string) string {
	var keyCount = make(map[rune]int, 256)
	for _, k := range s {
		if v, ok := keyCount[k]; ok {
			keyCount[k] = v + 1
		} else {
			keyCount[k] = 1
		}
	}

	// TODO: BUG NOTICE+1
	var countS = make([][]rune, len(s)+1)
	for k, num := range keyCount {
		var temp = make([]rune, 0)
		if len(countS[num]) != 0 {
			temp = countS[num]
		}
		for i := 0; i < num; i++ {
			temp = append(temp, k)
		}
		countS[num] = temp
	}

	var ret []rune
	for index := len(countS) - 1; index > 0; index-- {
		ret = append(ret, countS[index]...)
	}

	return string(ret)
}

// 思路: 用快排的partition算法找出一个数字, 左边的比其小，右边的比其大
// 这样通过多次比较旧可以找到这个第k大的数字
// SUMMARY:
// 1. 排序，然后取值 O(N*logN)
// 2. 最小堆，然后取值 O(N*logK)
// 3. Partion选择，然后取值 O(N*logK)
// TODO: FOCUS
func findKthLargest(nums []int, k int) int {
	var target = len(nums) - k
	var l, r = 0, len(nums) - 1
	var index = 0

	// TODO: [BUG NOTICE] 不可以直接比较index与target的值，必须经过partition之后，index的值才是有价值的！！
	for {
		index = partition(nums, l, r)
		if index > target {
			r = index - 1
			continue
		}

		if index < target {
			l = index + 1
			continue
		}
		break
	}

	return nums[index]
}

// TODO: NOTICE
func partition(nums []int, start, end int) int {
	var l, r = start, end
	var base = nums[l]
	for l < r {
		for l < r && nums[r] >= base {
			r--
		}

		if l < r {
			nums[l] = nums[r]
			l++
		}

		for l < r && nums[l] <= base {
			l++
		}

		if l < r {
			nums[r] = nums[l]
			r--
		}
	}
	nums[l] = base
	return l
}

// 思路: 首先用map来计数每个数字出现的次数, 之后有几种方式:
// 1. 排序， 用快排的话是时间复杂度 O(n*logN)
// 2. 最小堆， 时间复杂度 O(n*logk)
// 3. 用数组，数字出现的次数作为下表, 对应相关数字的slice， 之后倒叙遍历这个数据，可以拿到期望的最多出现的数字
func topKFrequeny(nums []int, k int) []int {
	var countM = make(map[int]int, 0)
	for _, n := range nums {
		if v, ok := countM[n]; ok {
			countM[n] = v + 1
		} else {
			countM[n] = 1
		}
	}

	// BUG NOTICE: 长度这里要加一 ！！
	var frequencyS = make([][]int, len(nums)+1)
	for key, num := range countM {
		var temp []int
		if len(frequencyS[num]) != 0 {
			temp = frequencyS[num]
		}

		temp = append(temp, key)
		frequencyS[num] = temp
	}

	var ret []int
	for i := len(frequencyS) - 1; i > 0; i-- {
		if len(frequencyS[i]) != 0 && len(ret) < k { // 如果有数字的出现的测试是一样的怎么办？比如1,1,2,2, 去topFrequency=1, 应该得哪个？
			ret = append(ret, frequencyS[i]...)
		}
	}

	return ret
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

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return middle
		}

		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}
