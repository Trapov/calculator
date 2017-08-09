package utils

import (
	"strings"
)

type Operator string

func (m Operator) IsOperatorAndLevel(operators []string) (bool, int) {
	for l, e := range operators {
		if strings.ContainsAny(m.ToString(), string(e)) {
			return true, l
		}
	}
	return false, -1
}

func (m Operator) ToString() string {
	return string(m)
}
