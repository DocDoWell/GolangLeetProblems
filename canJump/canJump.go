package canjump

import "fmt"

func canJump(nums []int) bool {
    arr_len := len(nums)
	dp := make([]bool, arr_len)

	fmt.Println(dp)

	dp[0] = true

	for i := 0; i < arr_len; i++ {
        if !dp[i] {
            continue
        }
		for j := nums[i]; j > 0; j-- {
			if i+j < arr_len {
				dp[i+j] = true
			}
		}
	}

	return dp[arr_len-1]
}


