package mymath

const Delta = 0.0001

func isConverged(d float64) bool {
	if d < 0.0 {
		d = -d
	}
	if d < Delta {
		return true
	}
	return false
}

func Sqrt2(x float64) float64 {
	z := 1.0
	tmp := 0.0
	for {
		tmp = z - (z * z - x) / 2 * z
		if d := tmp - z; isConverged(d) {
			return tmp
		}
		z = tmp
	}
	return z
}
