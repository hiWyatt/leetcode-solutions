package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	maphash := make(map[int]int)
	for _, value := range nums {
		// 统计频率
		maphash[value]++
	}
	// 取出去重的数字,即字典的key
	keys := make([]int, 0)
	for key, _ := range maphash {
		keys = append(keys, key)
	}
	// 按频率对数字排序,即按value对key排序
	sort.Slice(keys, func(i, j int) bool {
		// 排序规则(降序)
		return maphash[keys[i]] > maphash[keys[j]]
	})
	// 返回频率前k高的数字
	return keys[:k]
}

func main() {
	nums := []int{3, 3, 3, 2, 2, 1, 1, 1, 1}
	k := 2
	topKFrequent(nums, k)
}
