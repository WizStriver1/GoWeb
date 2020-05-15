package mymath

func Sqrt(x float64) float64 {
	m := float64(1)
	m = (m + x/m)/2
	n := m * m - x
	for n > 0 {
		m = (m + x/m)/2
		n = m * m - x
	}
	return m
}