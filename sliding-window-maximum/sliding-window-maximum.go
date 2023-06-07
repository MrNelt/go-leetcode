package slidingwindowmaximum

type Deque struct {
	list []int
}

func newDeque() Deque {
	return Deque{list: []int{}}
}

func (d *Deque) Push(x int) {
	d.list = append(d.list, x)
}

func (d *Deque) Len() int {
	return len(d.list)
}

func (d *Deque) PopFront() {
	d.list = d.list[1:]
}

func (d *Deque) PopBack() {
	n := len(d.list)
	d.list = d.list[:n-1]
}

func (d *Deque) Front() int {
	return d.list[0]
}

func (d *Deque) Back() int {
	return d.list[len(d.list)-1]
}

func maxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	d := newDeque()
	for i := 0; i < k; i++ {
		for d.Len() > 0 && d.Back() < nums[i] {
			d.PopBack()
		}
		d.Push(nums[i])
	}
	ans = append(ans, d.Front())
	for i := k; i < len(nums); i++ {
		if d.Front() == nums[i-k] {
			d.PopFront()
		}
		for d.Len() > 0 && d.Back() < nums[i] {
			d.PopBack()
		}
		d.Push(nums[i])
		ans = append(ans, d.Front())
	}
	return ans
}
