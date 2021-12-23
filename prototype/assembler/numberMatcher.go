package main

import "strconv"

type decimalMatcher struct {
	found string
}

func (m *decimalMatcher) match(check string) bool {
	valid := "0123456789"
	for i := 0; i < len(check); i++ {
		var found bool
		for _, v := range valid {
			if rune(check[i]) == v {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	m.found = check
	return true
}

func (m *decimalMatcher) value() []byte {
	i, _ := strconv.Atoi(m.found)
	return intToBytes(i)
}
