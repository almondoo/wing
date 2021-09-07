package array

// IsArray 配列の中に値が入っているか
func IsArray(array []interface{}, key string) bool {
	for _, value := range array {
		if value == key {
			return true
		}
	}
	return false
}
