package word_search_ii

type Trie struct {
	next   [26]*Trie
	isLeaf bool
	id     int
}

func Constructor() Trie {
	return Trie{next: [26]*Trie{}, isLeaf: false}
}

func (this *Trie) Insert(word string, id int) {
	temp := this
	for _, v := range word {
		if temp.next[v-'a'] == nil {
			node := Constructor()
			temp.next[v-'a'] = &node
		}
		temp = temp.next[v-'a']
	}
	temp.id = id
	temp.isLeaf = true
}

func (this *Trie) Search(word string) bool {
	temp := this
	for _, v := range word {
		if temp.next[v-'a'] == nil {
			return false
		}
		temp = temp.next[v-'a']
	}
	return temp.isLeaf
}

func (this *Trie) StartsWith(prefix string) bool {
	temp := this
	for _, v := range prefix {
		if temp.next[v-'a'] == nil {
			return false
		}
		temp = temp.next[v-'a']
	}
	return true
}

func FindWords(board [][]byte, words []string) []string {
	ans := []string{}
	t := Constructor()
	for i, v := range words {
		t.Insert(v, i)
	}
	visited := make([][]bool, len(board))
	for i := range board {
		visited[i] = make([]bool, len(board[i]))
	}
	cache := make(map[int]bool)
	ansId := []int{}
	for i := range board {
		for j := range board[i] {
			if t.next[board[i][j]-'a'] != nil {
				dfs(i, j, t.next[board[i][j]-'a'], board, visited, &ansId, cache)
			}
		}
	}
	for _, v := range ansId {
		ans = append(ans, words[v])
	}
	return ans
}

func dfs(i, j int, t *Trie, board [][]byte, visited [][]bool, ansId *[]int, cache map[int]bool) {
	if i < 0 || j > len(board[i])-1 || visited[i][j] {
		return
	}
	visited[i][j] = true
	if t.isLeaf && !cache[t.id] {
		*ansId = append(*ansId, t.id)
		cache[t.id] = true
	}
	if i-1 >= 0 && t.next[board[i-1][j]-'a'] != nil {
		dfs(i-1, j, t.next[board[i-1][j]-'a'], board, visited, ansId, cache)
	}
	if j-1 >= 0 && t.next[board[i][j-1]-'a'] != nil {
		dfs(i, j-1, t.next[board[i][j-1]-'a'], board, visited, ansId, cache)
	}
	if i+1 < len(board) && t.next[board[i+1][j]-'a'] != nil {
		dfs(i+1, j, t.next[board[i+1][j]-'a'], board, visited, ansId, cache)
	}
	if j+1 < len(board[i]) && t.next[board[i][j+1]-'a'] != nil {
		dfs(i, j+1, t.next[board[i][j+1]-'a'], board, visited, ansId, cache)
	}
	visited[i][j] = false
}
