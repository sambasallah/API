package main
import "fmt"
func sum(x ...int) int {
	sum := 0
	for _, el := range x {
		sum += el
	}

	return sum
}

func main() {
	fmt.Println(sum(10,20))
}