package utils

func RemoveDuplicateValues(stringSlice []string, existingValue string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range stringSlice {
		_, value := keys[entry]
		if !value && existingValue != entry {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
