package minimumcommonelement

func getCommon(nums1 []int, nums2 []int) int {
    for i:=0; i < len(nums1); i++ {
        if binarySearch(nums1[i], nums2){
            return nums1[i]
        }
    }
    return -1
}


func binarySearch(val int, array []int) (output bool) {
    mid := len(array)/2
    switch {
    case len(array) == 0:
        output = false 
    case array[mid] == val:
        output = true
    case array[mid] > val:
        if mid > 0 && array[mid-1] < val{
            output = false 
        }else{
            output = binarySearch(val, array[:mid])
        }
    case array[mid] < val:
        if mid < len(array)-1 && array[mid+1] > val{
            output = false 
        }else{
            output = binarySearch(val, array[mid+1:])
        }
    }
    return
}