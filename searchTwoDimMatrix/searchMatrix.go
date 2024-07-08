package searchMatrix

func searchMatrix(matrix [][]int, target int) bool {
    for i:=0; i < len(matrix); i++ {
        if  matrix[i][len(matrix[i])-1] == target{
            return true
        }
        if matrix[i][len(matrix[i])-1] >target{
            return binarySearch(matrix, target, i, 0, len(matrix[i])-1)
        }
    }    
    return false
}

func binarySearch(matrix [][]int, target int, row int, start int, end int) bool{

    if start > end{
        return false
    }

    mid:= (start+end)/2

    if matrix[row][mid] == target{
        return true
    }

    if matrix[row][mid] < target{
        return binarySearch(matrix, target, row, mid+1, end)
    }else{
        return binarySearch(matrix, target, row, start, mid-1)
    }
     

}