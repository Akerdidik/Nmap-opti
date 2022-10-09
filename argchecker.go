package nmapper

import "strconv"

func Argcheck() {
	if len(Args) == 0 {
		Helper()
		return
	}
	for i, j := range Args {
		if j == "--help" {
			Helper()
			return
		} else if j == "--target" && i+1 != len(Args) {
			Target = Args[i+1]
			Flag = true
		} else if j == "--port" && i+1 != len(Args) {
			t, _ := strconv.Atoi(Args[i+1])
			Port = t
		} else if j == "--target" && i+1 == len(Args) {
			Helper()
			return
		} else {
			if !Flag {
				Helper()
				return
			}
		}
	}
}
