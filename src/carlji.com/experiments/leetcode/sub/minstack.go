package sub

import "math"

// 既然是常量时间找到最小的节点值，那么只要记录这个值就行。
// 而记录之后，这里面解决的重点是这个值在Push和Pop动作之后的更新问题。
// 如果单纯的用两个变量值记录最小以及次小，在Pop是次小，然后在Pop最小的情况下，就无以为继了。
// 所以应该是每个值都能找到前一个最小的节点，故这里面需要用链表来保存最小的阶段。
type MinStack struct {
	data []int
	min  *Node
}

type Node struct {
	idx int // 此处必须是索引，记录最小的节点的位置，
	// 方便在pop时，精确判断是否是最小值被pop出去。而如果用值的话就不行，因为值有可能会重复。
	preMin *Node
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)

	if this.min == nil || this.data[this.min.idx] > x {
		this.min = &Node{idx: len(this.data) - 1, preMin: this.min}
	}
}

func (this *MinStack) Pop() {
	if len(this.data) == 0 {
		return
	}
	if len(this.data)-1 == this.min.idx {
		this.min = this.min.preMin
	}

	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return math.MinInt32
	}

	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.data) == 0 {
		return math.MinInt32
	}

	return this.data[this.min.idx]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
