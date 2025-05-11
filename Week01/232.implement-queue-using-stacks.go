/*
 * @lc app=leetcode id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 */

// @lc code=start
type MyQueue struct {
	input  []int
	output []int
}

func Constructor() MyQueue {
	return MyQueue{
		input:  make([]int, 0),
		output: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.input = append(this.input, x)
}

func (this *MyQueue) moveInputToOutput() {
	if len(this.output) == 0 {
		for len(this.input) > 0 {
			val := this.input[len(this.input)-1]
			this.input = this.input[:len(this.input)-1]
			this.output = append(this.output, val)
		}
	}
}

func (this *MyQueue) Pop() int {
	this.moveInputToOutput()

	if len(this.output) == 0 {
		return -1
	}

	val := this.output[len(this.output)-1]
	this.output = this.output[:len(this.output)-1]
	return val
}

func (this *MyQueue) Peek() int {
	this.moveInputToOutput()

	if len(this.output) == 0 {
		return -1
	}

	return this.output[len(this.output)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.input) == 0 && len(this.output) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

