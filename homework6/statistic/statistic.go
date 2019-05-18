package statistic

/*
 * Syntax Go Homework 6
 * Stepanov Anton, 18.05.2019
 * 1 пункт
 */

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func Sum(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total
}
