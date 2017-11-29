package algorithms

//Given n, how many structurally unique BST's (binary search trees) that store values 1...n?
//
//For example,
//Given n = 3, there are a total of 5 unique BST's.
//
//1         3     3      2      1
//\       /     /      / \      \
//3     2     1      1   3      2
///     /       \                 \
//2     1         2                 3

// 1
// 1
// 1

// 1，2
// 1		2
// \  	   	/
// 2	  	1
// 2

// 1，2，3
// 1		1		2		3		3
// \		\		/ \		/		/
// 2		3		1 3		2		1
// \		/				/		\
// 3		2				1		2
// 5

//卡特兰数
//当数组为 1，2，3，4，.. i，.. n时，基于以下原则的BST建树具有唯一性：
//以i为根节点的树，其左子树由[0, i-1]构成， 其右子树由[i+1, n]构成。
func numTrees(n int) int {
	if n < 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}
	}
	return dp[n]
}
