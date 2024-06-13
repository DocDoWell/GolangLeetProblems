package pow

func myPow(x float64, n int) (output float64) {

	minus := n < 0
	if minus {
		n = -n
	}

	output = 1
	for n != 0 {
		if n&1 != 0 { 
			output *= x
		}
		x *= x     
		n = n >> 1 
	}
	if minus {
		output = 1 / output
	}
	return
}
