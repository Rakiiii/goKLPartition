package klpartitinlin

//CheckInclusion checks if @num contains in @arr
func CheckInclusion(num int, arr []int) bool {
	for _, elem := range arr {
		if num == elem {
			return true
		}
	}
	return false
}
