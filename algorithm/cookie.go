package main

import (
	"fmt"
	"sort"
)

//假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。
//
//对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，都有一个尺寸 s[j] 。
//如果 s[j] >= g[i]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。

func findContentChildren(g []int, s []int) int {
	// 将孩子的胃口值和饼干的尺寸进行排序
	sort.Ints(g)
	sort.Ints(s)

	// 初始化指向孩子和饼干的指针
	childIndex := 0
	cookieIndex := 0

	// 使用贪心算法分配饼干
	for childIndex < len(g) && cookieIndex < len(s) {
		if s[cookieIndex] >= g[childIndex] {
			// 如果当前饼干可以满足当前孩子，分配饼干
			childIndex++
		}
		// 无论是否满足，饼干指针都需要移动
		cookieIndex++
	}

	// 返回满足的孩子数量
	return childIndex
}

func main() {
	// 示例数据
	g1 := []int{1, 2, 3}
	s1 := []int{1, 1}
	fmt.Println(findContentChildren(g1, s1)) // 输出: 1

	g2 := []int{1, 2}
	s2 := []int{1, 2, 3}
	fmt.Println(findContentChildren(g2, s2)) // 输出: 2
}
