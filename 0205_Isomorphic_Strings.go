package leetcode

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sTot := make(map[byte]byte)
	tTos := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		sChar, tChar := s[i], t[i]

		if sTot[sChar] > 0 && sTot[sChar] != tChar || tTos[tChar] > 0 && tTos[tChar] != sChar {
			return false
		}

		sTot[sChar] = tChar
		tTos[tChar] = sChar
	}
	return true
}
