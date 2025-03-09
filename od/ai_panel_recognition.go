package od

import (
	"fmt"
	"sort"
)

/*
AI识别到面板上有N（1 ≤ N ≤ 100）个指示灯，灯大小一样，任意两个之间无重叠。
由于AI识别误差，每次别到的指示灯位置可能有差异，以4个坐标值描述AI识别的指示灯的大小和位置(左上角x1,y1，右下角x2,y2)，

请输出先行后列排序的指示灯的编号，排序规则：
1. 每次在尚未排序的灯中挑选最高的灯作为的基准灯，
2. 找出和基准灯属于同一行所有的灯进行排序。两个灯高低偏差不超过灯半径算同一行（即两个灯坐标的差 ≤ 灯高度的一半）。
*/

/*
输入描述
1. 第一行为N，表示灯的个数
2. 接下来N行，每行为1个灯的坐标信息，格式为： {编号, x1, y1, x2, y2}
编号全局唯一
1 ≤ 编号 ≤ 100
0 ≤ x1 < x2 ≤ 1000
0  ≤  y1 < y2 ≤ 1000
*/

type Light struct {
	id int
	rx int
	ry int
	rl int
}

// LightSliceY implements sort.Interface based on the sorting rules provided.
type LightSliceY []Light

func (ls LightSliceY) Len() int {
	return len(ls)
}

func (ls LightSliceY) Less(i, j int) bool {
	// 比较 y 坐标（行），如果在同一行，则比较 x 坐标（列）
	if ls[i].ry == ls[j].ry {
		return ls[i].rx < ls[j].rx
	}
	return ls[i].ry < ls[j].ry
}

func (ls LightSliceY) Swap(i, j int) {
	ls[i], ls[j] = ls[j], ls[i]
}

func DoAiPanelRecognition() string {
	var num int
	fmt.Scan(&num)
	lights := make([][]int, num)

	for i := 0; i < num; i++ {
		var id int
		fmt.Scan(&id)

		var x1 int
		fmt.Scan(&x1)
		var y1 int
		fmt.Scan(&y1)
		var x2 int
		fmt.Scan(&x2)
		var y2 int
		fmt.Scan(&y2)
		lights[i] = []int{id, x1, y1, x2, y2}
	}

	return AiPanelRecognition(lights)
}

func AiPanelRecognition(lights [][]int) string {
	myLights := make([]Light, len(lights))
	for i, v := range lights {
		myLights[i] = Light{
			id: v[0],
			rx: (v[3] + v[1]) / 2,
			ry: (v[4] + v[2]) / 2,
			rl: (v[3] - v[1]) / 2,
		}
	}

	sort.Sort(LightSliceY(myLights))

	baseRyArr := make([]int, 0)

	m := make(map[int]([]Light))
	for _, v := range myLights {
		if len(baseRyArr) == 0 {
			baseRyArr = append(baseRyArr, v.ry)
			s := make([]Light, 0)
			m[v.ry] = append(s, v)
		} else {
			newLine := true
			for _, lastTopY := range baseRyArr {
				if lastTopY <= v.ry && lastTopY+v.rl >= v.ry {
					//同一组
					s := m[lastTopY]
					m[lastTopY] = append(s, v)
					newLine = false
					break
				}
			}
			if newLine {
				//new line
				baseRyArr = append(baseRyArr, v.ry)
				s := make([]Light, 0)
				m[v.ry] = append(s, v)
			}
		}
	}

	res := ""
	for _, v := range baseRyArr {
		for _, v := range m[v] {
			res = res + fmt.Sprintf("%d ", v.id)
		}
	}

	return res[:len(res)-1]
}
