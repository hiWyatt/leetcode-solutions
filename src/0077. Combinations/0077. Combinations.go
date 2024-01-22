package main

import "fmt"

// 全局变量
var (
	path []int
	res  [][]int
)

// 给定两个参数n和k，函数combine返回一个包含所有由1到n中选取k个数字组成的组合的二维数组
func combine(n int, k int) [][]int {
	path, res = make([]int, 0, k), make([][]int, 0)
	dfs(n, k, 1)
	return res
}

// 函数使用深度优先搜索（DFS）的方法来生成组合。它通过递归的方式不断构建路径path。
func dfs(n int, k int, startIndex int) {
	// 终止条件：当路径的长度等于k时，就将当前的组合加入结果数组res中
	if len(path) == k {
		// 由于切片是引用类型，在递归过程中，path会不断地被修改和回溯，最终所有的组合都会变成相同的结果，即最后一次递归回溯时的path。
		// 为了避免这种情况，需要创建一个新的临时切片temp，将当前的组合复制到temp中，然后将temp添加到结果数组res中。这样每次递归都会记录不同的组合，避免了结果的错误。
		temp := make([]int, k) //目标切片必须分配过空间且足够承载复制的元素个数
		copy(temp, path)
		res = append(res, temp) //存放结果
		return
	}
	// 为了避免重复的组合，每次递归都从startIndex开始，不往回走
	for i := startIndex; i <= n; i++ { // 选择本层集合中元素
		// 在每一层递归中，还会进行剪枝操作。如果剩余可选数字的数量不足以凑齐剩余需要的数字个数，则不再进行递归。
		if n-i+1 < k-len(path) { // 剪枝
			break
		}
		path = append(path, i)    // 处理节点
		dfs(n, k, i+1)            // 递归
		path = path[:len(path)-1] // 回溯
	}
}

func main() {
	//得到一个二维数组，其中包含了从1到4中选取2个数字的所有组合
	res := combine(4, 2)
	fmt.Println(res)
	//var path []int
	//path = make([]int, 0)
	//fmt.Println(len(path))
	//path = append(path, 1, 2)
	//fmt.Println(path)
	//temp := make([]int, 0)
	//copy(temp, path)
	//fmt.Println(temp)
}
