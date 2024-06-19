package trappingRainWater

import (
	"fmt"
	"math"
)

func trap(height []int) int {

	fmt.Println(" ")

	var output = 0

	if len(height) < 3 {
		return 0
	}

	start:=0
	end:= len(height)-1

	for{
		if height[start] == 0{
			start++
		}else{
			break
		}
	}

	for{
		if height[end] == 0{
			end--
		}else{
			break
		}
	}

	for i:=start+1; i < end;i++{
			maxRight:= getMaximumRight(i, height)
			maxLeft:= getMaximumLeft(i, height)
			fmt.Print("maxLeft is ", maxLeft, " maxRight is ", maxRight, " i is ", i)
			trapped:= int(math.Min(float64(maxLeft), float64(maxRight))) - height[i]
			if trapped >= 0{
				output = output + trapped
			}
			
			fmt.Print(" output added is ", int(math.Min(float64(maxLeft), float64(maxRight))) - height[i])
			fmt.Println(" ")
	}
	return output
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