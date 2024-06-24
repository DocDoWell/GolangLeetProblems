package searchRange

func searchRange(nums []int, target int) []int {
	output := []int{}

	mid := find(nums, target, 0, len(nums)-1)

	if mid == -1{
		output = append(output, mid)
		output = append(output, mid)
		return output
	}

	first:=mid
	second:=mid

	for first > 0{
		if nums[first] == target && nums[first-1] == target{
			first--
		}else{
			break
		}
	}

	for second < len(nums)-1 {
		if nums[second] == target && nums[second+1] == target{
			second++
		}else{
			break
		}
	}
	output = append(output, first)
	output = append(output, second)
	return output
}

func find(nums []int, target, left, right int) int {

	if right < left {
		return -1
	}

	mid := (left + right) / 2

	if nums[mid] == target {
		return mid
	}

	if nums[mid] < target {
		return find(nums, target, mid+1, right)
	} else {
		return find(nums, target, left, mid-1)
	}

}
