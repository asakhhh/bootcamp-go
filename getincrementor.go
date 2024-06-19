package bootcamp

var num_, step_ int

func incrementor_() int {
	num_ += step_
	return num_
}

func GetIncrementor(start, step int) func() int {
	num_, step_ = start, step
	return incrementor_
}
