package main

import (
	"fmt"
	"sort"
)

//假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。
//
//对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；
//并且每块饼干 j，都有一个尺寸 s[j] 。如果 s[j] >= g[i]，
//我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 2, 3}
	core(g, s)
}

// 保证每一步是最优解
func core(g []int, s []int) {
	sort.Ints(g)
	sort.Ints(s)

	children := 0

	for i, appetite := range g {
		if s[i] >= appetite {
			children++
			fmt.Println("appetite is", appetite)
		}
	}

}

// 每个孩子胃口值
func getAppetite(i int) int {
	var g []int
	g = append(g, 1, 2, 3)

	value := g[i]
	return value
}

// 每个饼干尺寸
func getSize(j int) int {
	var s []int
	s = append(s, 1, 2, 3)

	size := s[j]
	return size
}
