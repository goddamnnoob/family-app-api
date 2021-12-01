package util

type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	} else {
		l := len(*s) - 1
		ele := (*s)[l]
		(*s) = (*s)[:l]
		return ele, true
	}
}

func (s *Stack) push(val string) bool {
	*s = append((*s), val)
	return true
}
