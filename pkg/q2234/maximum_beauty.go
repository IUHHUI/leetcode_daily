package q2234

import (
	"sort"
)

// 1 <= flowers.length <= 10^5
// 1 <= flowers[i], target <= 10^5
// 1 <= newFlowers <= 10^10
// 1 <= full, partial <= 10^5
func maximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {
	lenGarden := len(flowers)
	if lenGarden == 1 {
		if flowers[0] >= target {
			//一个美丽花园
			return int64(full)
		} else {
			if newFlowers+int64(flowers[0]) >= int64(target) {
				return max(int64(full), int64(target-1)*int64(partial))
			} else {
				return int64(target-1) * int64(partial)
			}
		}
	}

	maxBeauty := int64(0)
	//尽量少的使用补种去完成需要的完美花园数量.剩下更多的补种去填补不完美花园.
	//通过排序可以把需要首先补种的花园放在后面.
	sort.Sort(sort.IntSlice(flowers))
	//有 k 个完美花园
	k := 0
	sumArr := make([]int64, lenGarden)
	for i, v := range flowers {
		if i == 0 {
			sumArr[i] = int64(v)
		} else {
			sumArr[i] = sumArr[i-1] + int64(v)
		}
		if v >= target {
			k++
		}
	}

	//i表示完美花园个数.
	for i := k; i <= lenGarden; i++ {
		j := i - k
		if j == 0 {
			//不需要补种为完美花园
			maxBeauty = int64(k) * int64(full)
			if k < lenGarden {
				//存在不完美花园
				subFlowers := flowers[:lenGarden-k]
				maxBeauty = maxBeauty + maxBeautyOfPartial(subFlowers, newFlowers, target, partial, sumArr)
			}
		} else {
			//需要补种为完美花园
			//尽量少的使用补种去完成需要的完美花园数量.剩下更多的补种去填补不完美花园.
			//1.判断补种的数量是否足够
			var sumUsed int64
			if i == lenGarden {
				sumUsed = int64(j)*int64(target) - (sumArr[lenGarden-1-k])
			} else {
				sumUsed = int64(j)*int64(target) - (sumArr[lenGarden-1-k] - sumArr[lenGarden-1-i])
			}

			if sumUsed > newFlowers {
				//2.补种数量不足
				continue
			}

			//3.补种数量足够
			if i == lenGarden {
				//全部补种为了完美花园
				maxTmp := int64(lenGarden) * int64(full)
				if maxTmp > maxBeauty {
					maxBeauty = maxTmp
				}
			} else {
				maxTmp := int64(i) * int64(full)

				subFlowers := flowers[:lenGarden-i]
				tmpNewFlowers := newFlowers - sumUsed
				maxTmp = maxTmp + maxBeautyOfPartial(subFlowers, tmpNewFlowers, target, partial, sumArr)
				if maxTmp > maxBeauty {
					maxBeauty = maxTmp
				}
			}
		}
	}

	return maxBeauty
}

/**
 * @Description: 不完美花园子最大美丽值
 * @param flowers: 不完美花园子集
 * @param newFlowers: 可补种的数量
 * @param target: 完美花园的阈值
 * @param partial: 不满园的美丽值
 * @return: 不满园的最大美丽值
 */
func maxBeautyOfPartial(subFlowers []int, newFlowers int64, target int, partial int, sumArr []int64) int64 {
	lenGarden := len(subFlowers)
	if lenGarden == 1 {
		if newFlowers+int64(subFlowers[0]) >= int64(target) {
			return int64(target-1) * int64(partial)
		} else {
			return (newFlowers + int64(subFlowers[0])) * int64(partial)
		}
	}
	maxBeauty := int64(0)
	sumFlow := sumArr[lenGarden-1]
	//1.补种数量足够
	// 每个花园尽量逼近 target, 使得美丽值最大.
	if int64(target-1)*int64(lenGarden)-sumFlow <= newFlowers {
		maxBeauty = int64(target-1) * int64(partial)
	} else {
		//2.补种数量不足
		// 每个花园尽量提高最小值, 使得美丽值最大.
		avg := (sumFlow + newFlowers) / int64(lenGarden)
		for i := lenGarden - 1; i >= 0; i-- {
			if i == 0 {
				maxBeauty = int64(subFlowers[i]+int(newFlowers)) * int64(partial)
				break
			}
			if subFlowers[i] > int(avg) {
				sumFlow -= int64(subFlowers[i])
				avg = (sumFlow + newFlowers) / int64(i)
			} else {
				//avg is min
				maxBeauty = int64(avg) * int64(partial)
				break
			}
		}
	}

	return maxBeauty
}

//可以优化的地方:
// 1. 对应单调递增的数组, 求出最低水位, 可以从前往后推.
// 必然有: j<i && flowers[j]*j <= presum+leftFlowers
// 最低水位 avg = (presum+leftFlowers)/j
