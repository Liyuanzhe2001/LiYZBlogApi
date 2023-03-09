package util

func IfUnit8(cond bool, a, b uint8) uint8 {
	if cond {
		return a
	}
	return b
}

func IfInterface(cond bool, a, b interface{}) interface{} {
	if cond {
		return a
	}
	return b
}
