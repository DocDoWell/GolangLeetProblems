package climbstairs

func climbStairs(n int) int {
    m := make(map[int]int)

    m[0] = 0
    m[1] = 1
    m[2] = 2

    if n <= 2{
        return m[n]
    }

    for i:=3; i<n; i++ {
        value := m[i-1] + m[i-2];
        m[i]=value;
    }

    return  m[n-1] + m[n-2]
}
