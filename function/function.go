package function

//Constant Declration
const Pi = 3.14

//Method for Swap : Example
func Swap(x, y int) (int, int) {
	return y, x
}

//Named return values example
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
