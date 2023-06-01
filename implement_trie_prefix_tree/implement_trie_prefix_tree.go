package implement_trie_prefix_tree

type Trie struct {
	next   [26]*Trie
	isLeaf bool
}

func Constructor() Trie {
	return Trie{next: [26]*Trie{}, isLeaf: false}
}

func (this *Trie) Insert(word string) {
	temp := this
	for _, v := range word {
		if temp.next[v-'a'] != nil {
			temp = temp.next[v-'a']
		} else {
			node := Constructor()
			temp.next[v-'a'] = &node
			temp = temp.next[v-'a']
		}
	}
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

/**
* Your Trie object will be instantiated and called as such:
* obj := Constructor();
* obj.Insert(word);
* param_2 := obj.Search(word);
* param_3 := obj.StartsWith(prefix);
 */
