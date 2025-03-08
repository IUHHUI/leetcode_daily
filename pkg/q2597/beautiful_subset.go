package q2597

// 1 <= nums.length <= 18
// 1 <= nums[i], k <= 1000
//
// nums 的子集定义为：可以经由 nums 删除某些元素（也可能不删除）得到的一个数组。
// 只有在删除元素时选择的索引不同的情况下，两个子集才会被视作是不同的子集。
//
// 任意两个整数的绝对差均不等于 k
func beautifulSubsets(nums []int, k int) int {
	ans := 0
	m := make(map[int]int)

	var dfs func(i int)

	dfs = func(i int) {
		if i == len(nums) {
			ans++
			return
		}
		dfs(i + 1)
		if m[nums[i]-k] == 0 && m[nums[i]+k] == 0 {
			//将 nums[i] 添加到子集中，需要满足 (nums[i] −k) 和 (nums[i] +k) 还没有被添加到子集中。
			m[nums[i]]++

			dfs(i + 1)
			m[nums[i]]--
		}
	}

	//这里递归最大的好处是, 输出的参是第一个元素下标.
	//首先计算出结果的确实 受最后一个元素影响的结果.

	dfs(0)
	return ans - 1
}
