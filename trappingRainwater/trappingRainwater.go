package trappingRainWater

import (
	"math"
)

func trap(height []int) int {	
	if len(height) < 3 {
		return 0
	}
	var output = 0
	start:=geStart(height)
	end:= getEnd(height)
	for i:=start+1; i < end;i++{
		maxRight:= getMaximumRight(i, height)
		maxLeft:= getMaximumLeft(i, height)
		trapped:= int(math.Min(float64(maxLeft), float64(maxRight))) - height[i]
		if trapped >= 0{
			output = output + trapped
		}
	}
	return output
}

func geStart(height []int) (start int) {
	start =0
	for{
		if height[start] == 0{
			start++
		}else{
			break
		}
	}
	return start
}

func getEnd(height []int) (end int) {
	end= len(height)-1
	for{
		if height[end] == 0{
			end--
		}else{
			break
		}
	}
	return end
}

func getMaximumLeft(index int, height []int) (output int){
	output = 0
	for j:=0; j< index; j++{
		if height[j] > output{
			output = height[j]
		}
	}
	return output
}

func getMaximumRight(index int, height []int) (output int){
	output = 0
	for j:=index+1; j< len(height); j++{
		if height[j] > output{
			output = height[j]
		}
	}
	return output
}