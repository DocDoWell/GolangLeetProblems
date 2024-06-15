package pow

func output(){
    var a uint16 = 12
	var b uint16 = 213

	fmt.Printf("Number %d in binary is %b\n", a , a)

	fmt.Printf("Number %d in binary is %b\n", b , b)

	fmt.Printf("Number %d in binary is %b\n", a & b , a& b)

	fmt.Printf("Number %d in binary is %b\n", a << 1, a << 1)
// 11010101 << 1  = 110101010 (We shift it once)

	fmt.Printf("Number %d in binary is %b\n", a << 2, a << 2)	

	fmt.Printf("Number %d in binary is %b\n", a >> 1, a >> 1)

	fmt.Printf("Number %d in binary is %b\n", a >> 2, a >> 2)
}