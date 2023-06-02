package design_add_and_search_words_data_structure

type WordDictionary struct {
	next   [26]*WordDictionary
	isLeaf bool
}

func Constructor() WordDictionary {
	return WordDictionary{[26]*WordDictionary{}, false}
}

func (this *WordDictionary) AddWord(word string) {
	temp := this
	for _, v := range word {
		if temp.next[v-'a'] == nil {
			node := Constructor()
			temp.next[v-'a'] = &node
		}
		temp = temp.next[v-'a']
	}
	temp.isLeaf = true
}

func (this *WordDictionary) Search(word string) bool {
	if word == "" {
		return this.isLeaf
	}
	temp := this
	for i, v := range word {
		if v == '.' {
			flag := false
			for j := 0; j < 26; j++ {
				if temp.next[j] != nil {
					if temp.next[j].Search(word[i+1:]) {
						flag = true
					}
				}
			}
			return flag
		} else if temp.next[v-'a'] == nil {
			return false
		} else {
			temp = temp.next[v-'a']
		}
	}
	return temp.isLeaf
}
