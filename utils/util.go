package utils

func RemoveDuplicateStrings(strList []string) []string {
	stringMap := make(map[string]bool)
	for _, str := range strList {
		stringMap[str] = true
	}
	newStrList := []string{}
	for key := range stringMap {
		if key != "" {
			newStrList = append(newStrList, key)
		}
	}
	return newStrList
}
