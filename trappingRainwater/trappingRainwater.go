package trappingRainWater

import (
	"fmt"
	"math"
)

func trap(height []int) int {
	var output = 0

	if len(height) < 3 {
		return 0
	}

	var leftBoundry = instantiateLeftBoundary(height)
	var rightBoundary = getClosingBoundary(leftBoundry, height)
	output = output + computeRainWaterTrapped(leftBoundry, rightBoundary, height)
	fmt.Println("left is ", leftBoundry, " right is ", rightBoundary, " output is ", output)

	for rightBoundary < (len(height) - 1) {
		leftBoundry = rightBoundary
		rightBoundary = getClosingBoundary(leftBoundry, height)
		output = output + computeRainWaterTrapped(leftBoundry, rightBoundary, height)
		fmt.Println("left is ", leftBoundry, " right is ", rightBoundary, " output is ", output)
	}

	return output
}

func instantiateLeftBoundary(height []int) int {
	var leftBoundary = 0
	for {
		if height[leftBoundary] == 0 {
			leftBoundary++
		} else {
			break
		}
	}
	return leftBoundary
}

func computeRainWaterTrapped(leftBoundry, rightBoundary int, height []int) (output int) {
	output = 0
	var pointer = leftBoundry + 1
	for pointer < rightBoundary {
		fmt.Print("at ", pointer, " with ", int(math.Min(float64(height[leftBoundry]), float64(height[rightBoundary]))) - height[pointer])
		if height[pointer] < height[leftBoundry] {
			output = output + int(math.Min(float64(height[leftBoundry]), float64(height[rightBoundary]))) - height[pointer]
			fmt.Print(" output is ", output)
		}	
		pointer++
		fmt.Println(" ")
	}
	return output
}

func getClosingBoundary(leftPointer int, height []int) int {
	max := getMaximumHeight(height)
	for i := leftPointer + 1; i < len(height); i++ {
		if (height[i] >= height[leftPointer]) && (height[i]<max) {
			return i
		}
	}
	return len(height) - 1
}


func getMaximumHeight(height []int) (output int){
	output = 0
	for i:=0; i< len(height); i++{
		if height[i] > output{
			output = height[i]
		}
	}
	return output
}
