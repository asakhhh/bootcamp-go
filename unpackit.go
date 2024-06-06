package bootcamp

func UnpackIt(p *******int) int {
	if p == nil {
		return 0
	}

	var p6 ******int = *p
	if p6 == nil {
		return 0
	}

	var p5 *****int = *p6
	if p5 == nil {
		return 0
	}

	var p4 ****int = *p5
	if p4 == nil {
		return 0
	}

	var p3 ***int = *p4
	if p3 == nil {
		return 0
	}

	var p2 **int = *p3
	if p2 == nil {
		return 0
	}

	var p1 *int = *p2
	if p1 == nil {
		return 0
	}

	return *p1
}
