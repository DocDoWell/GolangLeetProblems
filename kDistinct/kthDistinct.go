package kdistinct

func kthDistinct(arr []string, k int) string {
    m := make(map[string]int)
    pointer := 0
    
    for i:=0; i< len(arr); i++{
        m[arr[i]]++
    }

    for i:=0; i< len(arr); i++{
        if m[arr[i]] == 1{
            pointer++
        }
        if pointer == k{
            return arr[i]
        }
    }

    return ""
}