// Ftoc prints two Fahrenheit-to-Celsius conversions.
package main
import "fmt"
func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g F = %g \n", freezingF, fToC(freezingF)) // 32F = 0C
	fmt.Printf("%g F = %g C\n", boilingF, fToC(boilingF)) // 212F = 100C
	var arr[3] int
	fmt.Println("arr[0] = %d", arr[0])
	x := &arr[0]
	fmt.Println("arr[0] = %d", *x)
	*x = 5
	fmt.Println("arr[0] = %d", *x)
	fmt.Println("arr[0] = %d", arr[0])
	fmt.Println(f() == f())
}
func fToC(f float64) float64 {
	return (f -32) * 5/9
}

func f() *int {
	v := 1
	return &v
}
