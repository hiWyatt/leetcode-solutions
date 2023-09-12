package main

var (
	path []int
	res  [][]int
)

//给定两个参数n和k，函数combine返回一个包含所有由1到n中选取k个数字组成的组合的二维数组
func combine(n int, k int) [][]int {
	path, res = make([]int, 0, k), make([][]int, 0)
	dfs(n, k, 1)
	return res
}

//函数使用深度优先搜索（DFS）的方法来生成组合。它通过递归的方式不断构建路径path。
func dfs(n int, k int, start int) {
	//当路径的长度等于k时，就将当前的组合加入结果数组res中
	if len(path) == k { // 说明已经满足了k个数的要求
		tmp := make([]int, k)
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	//为了避免重复的组合，每次递归都从start开始，不往回走。
	for i := start; i <= n; i++ { // 从start开始，不往回走，避免出现重复组合
		//在每一层递归中，还会进行剪枝操作。如果剩余可选数字的数量不足以凑齐剩余需要的数字个数，则不再进行递归。
		if n-i+1 < k-len(path) { // 剪枝
			break
		}
		path = append(path, i)
		dfs(n, k, i+1)
		path = path[:len(path)-1]
	}
}

func main() {
	//得到一个二维数组，其中包含了从1到4中选取2个数字的所有组合
	combine(4, 2)
}
