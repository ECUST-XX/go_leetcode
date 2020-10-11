package leet13

func RomanToInt(s string) int {
	romanIntMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	specialRomanIntMap := map[byte][2]byte{
		'I': {'V', 'X'},
		'X': {'L', 'C'},
		'C': {'D', 'M'},
	}

	res := 0
	specialTemp := 0
	var specialFlag byte
	for i := 0; i < len(s); i++ {
		if _, ok := specialRomanIntMap[s[i]]; ok && specialRomanIntMap[specialFlag][0] != s[i] && specialRomanIntMap[specialFlag][1] != s[i] {
			if s[i] != specialFlag {
				res += specialTemp
				specialTemp = 0
			}
			specialFlag = s[i]
			specialTemp += romanIntMap[s[i]]
			continue
		}
		if specialTemp != 0 && (specialRomanIntMap[specialFlag][0] == s[i] || specialRomanIntMap[specialFlag][1] == s[i]) {
			res += romanIntMap[s[i]] - specialTemp
			specialTemp = 0
			specialFlag = 0
			continue
		}
		res += romanIntMap[s[i]]
	}
	res += specialTemp

	return res
}
