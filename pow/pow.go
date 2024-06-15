package pow

func myPow(x float64, n int) (output float64) {
    if n==0 || x == 1 {
        return 1
    }
    if n==1 {
        return x
    }
    minus := n <0 
	if minus {
		n=(-1)*n
	}
	return compute(x,n, 1, minus)
}


func compute(x float64, n int, output float64, minus bool) float64 {
     if n == 0  {
        if(minus){
            return 1/output
        }
		    return output
	}
    
    if n&1 != 0 { 
		output *= x
	}
	x *= x     
	n = n >> 1 
    return compute(x, n, output, minus)
}