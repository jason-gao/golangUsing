package string

import "time"

func StrJoin(s string) string{
	s1 := s + " - " + time.Now().String()

	return s1
}