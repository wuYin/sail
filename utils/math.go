package utils

import "log"

func Max(vs ...int) int {
	if len(vs) == 0 {
		log.Fatal("max args empty")
	}
	m := vs[0]
	for _, v := range vs[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

func Min(vs ...int) int {
	if len(vs) == 0 {
		log.Fatal("max args empty")
	}
	m := vs[0]
	for _, v := range vs[1:] {
		if v < m {
			m = v
		}
	}
	return m
}
