package utils

func StatusContains(slice []int, key int) bool {
	for _, v := range slice {
		if v == key {
			return true
		}
	}
	return false
}

func RemoveDuplicate(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
