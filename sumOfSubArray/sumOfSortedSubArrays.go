package sumOfSubArray

import (
	"sort"
)

func rangeSum(nums []int, n int, left int, right int) int {

	output:= []int{}
    
    for i:=0; i<n; i++ {
        output = append(output, nums[i])
        sum := nums[i]
        for j:=i+1; j<n; j++{
            sum = sum + nums[j]
            output = append(output, sum)
        }
    }

	sort.Slice(output, func(i, j int) bool {
		return output[i] < output[j]
	})

    outputSum := 0
    mod:= int(1e9+7)
    
    for i:=left-1; i<=right-1; i++ {
        outputSum = (outputSum + output[i])%mod
    }

    return outputSum
}