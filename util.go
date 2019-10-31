package klpartitinlin

func CheckInclusion(num int, arr []int) bool {
	for _, elem := range arr {
		if num == elem {
			return true
		}
	}
	return false
}
