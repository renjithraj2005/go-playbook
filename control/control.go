package control

//For loop in Golang
func Forloop() (int, int) {
	sum := 0
	for i := 0; i < 100; i++ {
		sum = sum + i
	}

	//The init and post statement are optional.
	test := 0

	for test < 5000 {
		test += test
	}

	return sum, test
}
