package trappingRainWater

func trap(height []int) int {
	output, leftPointer, rightPointer, leftBoundary, rightBoundary:= 0, 0, len(height)-1, 0, 0
    for leftPointer <= rightPointer {
        if leftBoundary <= rightBoundary {
            if height[leftPointer] > min(leftBoundary, rightBoundary) {
                leftBoundary = height[leftPointer]
            }
            output += max(min(leftBoundary, rightBoundary) - height[leftPointer], 0)
            leftPointer++
        } else {
            if height[rightPointer] > min(leftBoundary, rightBoundary) {
                rightBoundary = height[rightPointer]
            }
            output += max(min(leftBoundary, rightBoundary) - height[rightPointer], 0)
            rightPointer--
        }
    }
    return output
}

func min(leftBoundary, rightBoundary int) int {
	if(leftBoundary < rightBoundary){
		return leftBoundary
	}
	return rightBoundary
}

func max(leftBoundary, rightBoundary int) int {
	if(leftBoundary > rightBoundary){
		return leftBoundary
	}
	return rightBoundary
}