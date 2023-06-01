package generate_parentheses

func gen(n, open, close int, ans string, list *[]string) {
	if open+close == 2*n {
		*list = append(*list, ans)
	}
	if open < n {
		gen(n, open+1, close, ans+"(", list)
	}
	if open > close {
		gen(n, open, close+1, ans+")", list)
	}
}

func generateParenthesis(n int) []string {
	list := []string{}
	gen(n, 0, 0, "", &list)
	return list
}
