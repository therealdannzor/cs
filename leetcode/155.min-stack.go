/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 *
 * https://leetcode.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (40.89%)
 * Likes:    3035
 * Dislikes: 297
 * Total Accepted:    511.6K
 * Total Submissions: 1.2M
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * Design a stack that supports push, pop, top, and retrieving the minimum
 * element in constant time.
 *
 *
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * getMin() -- Retrieve the minimum element in the stack.
 *
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["MinStack","push","push","push","getMin","pop","top","getMin"]
 * [[],[-2],[0],[-3],[],[],[],[]]
 *
 * Output
 * [null,null,null,null,-3,null,0,-2]
 *
 * Explanation
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin(); // return -3
 * minStack.pop();
 * minStack.top();    // return 0
 * minStack.getMin(); // return -2
 *
 *
 *
 * Constraints:
 *
 *
 * Methods pop, top and getMin operations will always be called on non-empty
 * stacks.
 *
 *
 */

// @lc code=start
type MinStack struct {
	stack []int
	min   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 {
		this.min = x
	}

	if x < this.min {
		this.min = x
	}

	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	// save the element to pop and resize stack
	last := len(this.stack) - 1
	item := this.stack[last]
	this.stack = this.stack[:last]

	// if the element pop is not the minimum,
	// bail out here
	if item != this.min {
		return
	}

	// element was the minimum, next in stack is
	// new (temporary) minimum
	if len(this.stack) > 0 {
		last--
	} else {
		return
	}

	this.min = this.stack[last]

	// for every time in the stack, check if there
	// is a smaller than the temporary one
	for i := last - 1; i > -1; i-- {
		if this.stack[i] < this.min {
			this.min = this.stack[i]
		}
	}
}

func (this *MinStack) Top() int {
	last := len(this.stack) - 1
	return this.stack[last]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
