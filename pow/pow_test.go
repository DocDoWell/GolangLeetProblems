package pow

import (
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {


	z:= math.Round(myPow(2.0000, 10)*100)/100
	c:= math.Round(myPow(2.1000, 3)*1000)/1000
	d:= math.Round(myPow(2.0000, -2)*100)/100

	if z != 1024.00{
		t.Fatalf(`expecting 1024.00000 but got %v`, z)
	}

	if c != 9.261{
		t.Fatalf(`expecting 9.261000 but got %v`, c)
	}

	if d != 0.25{
		t.Fatalf(`expecting 9.26100 but got %v`, d)
	}


}
