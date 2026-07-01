func groupAnagrams(strs []string) [][]string {
	freqMap := make(map[[26]uint16]int)
	res := [][]string{}
	for _, s := range strs {
		freq := [26]uint16{}
		for _, c := range s {
			freq[c-'a']++
		}
		if _, ok := freqMap[freq]; !ok {
			freqMap[freq] = len(res)
			res = append(res, nil)
		}
		res[freqMap[freq]] = append(res[freqMap[freq]], s)
	}
	return res
}