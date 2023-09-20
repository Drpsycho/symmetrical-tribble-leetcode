package anagram

func isAnagram(s string, t string) bool {
	if len(s) == 0 || len(t) == 0 || len(s) != len(t) {
		return false
	}
	var temp_map_s = make(map[rune]int)
	var temp_map_t = make(map[rune]int)
	for i := 0; i < len(s); i++ {
		temp_map_s[rune(s[i])] += 1
		temp_map_t[rune(t[i])] += 1
	}
	for key, value := range temp_map_s {
		if temp_map_t[key] != value {
			return false
		}
	}
	return true
}
