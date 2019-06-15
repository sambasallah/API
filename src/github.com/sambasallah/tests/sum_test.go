package sambasallah

import (
	"testing"	
)


func TestSum(t *testing.T) {
	sum := sum(10,20)
	if sum != 30 {
		t.Error("Expected 30, got ", sum)
	}
}