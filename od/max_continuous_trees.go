package od

// 滑动窗口
func MaxContinuousTrees(n int, m int, deadIndices []int, k int) int {
	if len(deadIndices) == k {
		return n
	}

	//胡杨（编号1-N）
	//deadIndices 里面是杨树编号.
	trees := make([]int, n)
	for _, v := range deadIndices {
		// 标记死树
		trees[v-1] = 1
	}

	var kCnt = k
	var max int = 0
	var wdLeft int = 0
	for wdRight := 0; wdRight < n; {
		if trees[wdRight] == 1 {
			if kCnt > 0 {
				//补种一颗
				kCnt--
				//往前走
				wdRight++
			} else {
				//当前 wdRight 指向的是死树
				nowLen := wdRight - wdLeft
				if max < nowLen {
					max = nowLen
				}

				if wdLeft == wdRight {
					//窗口内没死树可以移除.
					break
				}

				if trees[wdLeft] == 1 {
					//前头移除一颗补种树, kCnt加1
					kCnt++
				}
				wdLeft = wdLeft + 1
			}
		} else {
			//当前 wdRight 指向的是活树
			nowLen := wdRight - wdLeft + 1
			if max < nowLen {
				max = nowLen
			}

			//往前走
			wdRight++
		}
	}

	return max
}
