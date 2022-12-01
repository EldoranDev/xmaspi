package math

func Map(val, fromMin, fromMax, toMin, toMax int) int {
	return (val-fromMin)*(toMax-toMin)/(fromMax-fromMin) + toMin
}
