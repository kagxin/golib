package function

// IIF 三元运算
func IIF(expr bool, turePart, falsePart interface{}) interface{} {
	if expr {
		return turePart
	}
	return falsePart
}
