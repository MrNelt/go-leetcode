package min_stack

type MinStack struct {
	elems []int
	mins  []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(val int) {
	if len(this.elems) == 0 {
		this.elems = append(this.elems, val)
		this.mins = append(this.mins, val)
	} else {
		this.elems = append(this.elems, val)
		this.mins = append(this.mins, min(this.mins[len(this.mins)-1], val))
	}
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func (this *MinStack) Pop() {
	this.elems = this.elems[:len(this.elems)-1]
	this.mins = this.mins[:len(this.mins)-1]
}

func (this *MinStack) Top() int {
	return this.elems[len(this.elems)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}
