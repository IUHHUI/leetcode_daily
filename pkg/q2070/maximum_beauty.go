package q2070

import (
	"sort"
)

/**
 * @author hui
 * @date 20250309
 * @description
 * @items 一个二维整数数组 items ， 其中 items[i] = [pricei, beautyi] 分别表示每一个物品的 价格 和 美丽值 。
 */
func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	minPrice := items[0][0]
	maxPrice := items[len(items)-1][0]

	maxBeautyMap := make(map[int]int)
	tmpMaxBeauty := 0
	for _, v := range items {
		if tmpMaxBeauty == 0 {
			tmpMaxBeauty = v[1]
		} else {
			if tmpMaxBeauty < v[1] {
				tmpMaxBeauty = v[1]
			}
		}
		maxBeautyMap[v[0]] = tmpMaxBeauty
	}

	res := make([]int, len(queries))
	for i, memory := range queries {
		if memory < minPrice {
			res[i] = 0
		} else if memory >= maxPrice {
			res[i] = maxBeautyMap[maxPrice]
		} else {
			//二分查找
			priceIndex := query(items, memory)
			res[i] = maxBeautyMap[items[priceIndex][0]]
		}
	}
	return res
}

// 升序数组可以使用二分查找
func query(items [][]int, memory int) int {
	left := 0
	right := len(items)
	for left < right {
		mid := left + (right-left)/2
		if items[mid][0] < memory {
			left = mid + 1
		} else if items[mid][0] == memory {
			return mid
		} else {
			right = mid
		}
	}
	if items[left][0] > memory {
		return left - 1
	} else {
		return left
	}
}

//可以考虑优化的情况
//去掉 map, 复用items[i][1] 存储该价格的最大美丽值.
