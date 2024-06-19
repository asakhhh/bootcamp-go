package bootcamp

type mathoper func(int, int) int

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func GetMathOperation(op string) *mathoper {
	var ptr mathoper
	if op == "add" {
		ptr = add
	} else if op == "subtract" {
		ptr = subtract
	} else if op == "multiply" {
		ptr = mult
	} else if op == "divide" {
		ptr = div
	} else {
		ptr = nil
	}
	return &ptr
}

// func main() {
// 	add := GetMathOperation("add")
// 	if add != nil {
// 		result := (*add)(2, 3)
// 		fmt.Println(result) // 5
// 	} else {
// 		fmt.Println("Invalid operation")
// 	}

// 	subtract := GetMathOperation("subtract")
// 	if subtract != nil {
// 		result := (*subtract)(5, 2)
// 		fmt.Println(result) // 3
// 	} else {
// 		fmt.Println("Invalid operation")
// 	}

// 	multiply := GetMathOperation("multiply")
// 	if multiply != nil {
// 		result := (*multiply)(3, 4)
// 		fmt.Println(result) // 12
// 	} else {
// 		fmt.Println("Invalid operation")
// 	}

// 	divide := GetMathOperation("divide")
// 	if divide != nil {
// 		result := (*divide)(10, 2)
// 		fmt.Println(result) // 5
// 		result = (*divide)(10, 0)
// 		fmt.Println(result) // 0
// 	} else {
// 		fmt.Println("Invalid operation")
// 	}

// 	invalid := GetMathOperation("mod")
// 	if invalid != nil {
// 		result := (*invalid)(10, 3)
// 		fmt.Println(result)
// 	} else {
// 		fmt.Println("Invalid operation")
// 	}
// }
