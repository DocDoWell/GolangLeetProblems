package binaryAddition

func addBinary(a string, b string) string {    
	if(len(a)>=len(b)){
	   return computeBinaryString(b,a)
	}else{
	   return computeBinaryString(a,b)
	}
}

func computeBinaryString(a string, b string) string {
   diff := len(b) - len(a)

  for i := 0; i < diff; i++ {
	 a = "0" + a
  }

  carry := "0"
  output := ""

   for i := len(a) - 1; i >= 0; i-- {
	 if a[i] == '1' && b[i] == '1' {
		if carry == "1" {
		   output = "1" + output
		} else {
		   output = "0" + output
		   carry = "1"
		}
	 } else if a[i] == '0' && b[i] == '0' {
		if carry == "1" {
		   output = "1" + output
		   carry = "0"
		} else {
		   output = "0" + output
		}
	 } else if a[i] != b[i] {
		if carry == "1" {
		   output = "0" + output
		} else {
		   output = "1" + output
		}
	 }
  }
  if carry == "1" {
	 output = "1" + output
  }
  return output
}
