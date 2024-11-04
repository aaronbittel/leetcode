package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sBytes := []byte(s)

	chMap := make(map[byte]byte)
	mapped := make(map[byte]struct{})

	for i, ch := range sBytes {
		if _, ok := chMap[ch]; !ok {
			chMap[ch] = t[i]
			if _, ok := mapped[t[i]]; ok {
				return false
			}
			mapped[t[i]] = struct{}{}
		}
		sBytes[i] = chMap[ch]
	}

	return string(sBytes) == t
}
