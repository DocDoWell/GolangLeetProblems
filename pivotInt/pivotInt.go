package pivotInt

func pivotInteger(n int) int {
    for index:=1; index<= n; index++{
		toPrevIndex:= (index-1)*(index)/2
        toIndex:= index*(index+1)/2
        fromIndex:= n*(n+1)/2 - (toPrevIndex)
        if toIndex == fromIndex{
            return index
        }
    }
    return -1
}
