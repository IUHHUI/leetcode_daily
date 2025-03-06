package q2588

// 1 <= nums.length <= 10**5
// 0 <= nums[i] <= 10**6
// 数组很大,时间和空间复杂度不要超过O(n)
//
// 最美子数组中数组里面的元素 异或运算后 为 0.
// 题目转化为 异或运算后为 0 的子数组 数量
// 对于数组[0..i], 如果存在[j...i]为最美子数组
// a[j]^a[j+1]...a[i]=0
// 那么 a[0]^a[1]...a[j-1]^a[j]...a[i]=a[0]^a[1]...a[j-1]
func beautifulSubarrays(nums []int) int64 {

	//任何一个子数组的xor值不会大于子数组中最大值的两倍，
	//所以xor值最大不会超过2*10^6，或者不会超过数组中最大值的两倍，可以用整型数组代替Map加速。

	ans := int64(0)
	xorCntMap := make(map[int]int, len(nums))
	xorVal := 0
	xorCntMap[0] = 1
	for _, v := range nums {
		xorVal ^= v
		preXorCnt, ok := xorCntMap[xorVal]
		if ok {
			ans += int64(preXorCnt)
		}
		xorCntMap[xorVal]++
	}

	return ans
}
