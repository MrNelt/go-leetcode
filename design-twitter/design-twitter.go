package designtwitter

import "container/heap"

type HeapInt []int

func (h HeapInt) Len() int {
	return len(h)
}

func (h HeapInt) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h HeapInt) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *HeapInt) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *HeapInt) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Twitter struct {
	nowTime             int
	userInfoTime        map[int][]int
	userTweetIDFromTime map[int]int
	userInfoFollow      map[int]map[int]bool
}

func Constructor() Twitter {
	userInfoTime := make(map[int][]int)
	userTweetIDFromTime := make(map[int]int)
	userInfoFollow := make(map[int]map[int]bool)
	return Twitter{nowTime: 0, userInfoTime: userInfoTime, userTweetIDFromTime: userTweetIDFromTime, userInfoFollow: userInfoFollow}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.userInfoTime[userId] = append(this.userInfoTime[userId], this.nowTime)
	this.userTweetIDFromTime[this.nowTime] = tweetId
	this.nowTime++
	this.Follow(userId, userId)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	h := HeapInt{}
	for key, value := range this.userInfoFollow[userId] {
		if value {
			h = append(h, this.userInfoTime[key]...)
		}
	}
	heap.Init(&h)
	output := []int{}
	count := 0
	for count < 10 && len(h) > 0 {
		timePost := heap.Pop(&h).(int)
		output = append(output, this.userTweetIDFromTime[timePost])
		count++
	}
	return output
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.userInfoFollow[followerId]; !ok {
		temp := make(map[int]bool)
		this.userInfoFollow[followerId] = temp
	}
	this.userInfoFollow[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followeeId == followerId {
		return
	}
	if _, ok := this.userInfoFollow[followerId]; !ok {
		temp := make(map[int]bool)
		this.userInfoFollow[followerId] = temp
	}
	this.userInfoFollow[followerId][followeeId] = false
}
