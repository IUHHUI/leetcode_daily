package od

import (
	"strconv"
	"strings"
)

// Ipv42Int ipv4转int
/*
存在一种虚拟IPv4地址，由4小节组成，每节的范围为0~255，以#号间隔，虚拟IPv4地址可以转换为一个32位的整数，例如：

128#0#255#255，转换为32位整数的结果为2147549183（0x8000FFFF）
1#0#0#0，转换为32位整数的结果为16777216（0x01000000）
现以字符串形式给出一个虚拟IPv4地址，限制第1小节的范围为1~128，即每一节范围分别为(1~128)#(0~255)#(0~255)#(0~255)，要求每个IPv4地址只能对应到唯一的整数上。

如果是非法IPv4，返回invalid IP
*/
func Ipv42Int(ip string) int {
	strs := strings.Split(ip, "#")
	if len(strs) != 4 {
		return 0
	}

	ipInt := 0
	for i, v := range strs {
		num, e := strconv.Atoi(v)
		if e != nil {
			return 0
		}

		if i == 0 {
			if num < 1 || num > 128 {
				return 0
			}
			ipInt = num << 24
		} else {
			if num < 0 || num > 255 {
				return 0
			}
			if i == 1 {
				ipInt += num << 16
			} else if i == 2 {
				ipInt += num << 8
			} else if i == 3 {
				ipInt += num
			}
		}

	}

	return ipInt
}
