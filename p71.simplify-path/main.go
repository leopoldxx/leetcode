package main

import "strings"

func split(path []byte) [][]byte {
	index := func(p []byte) int {
		for i := 0; i < len(p); i++ {
			if p[i] == '/' {
				return i
			}
		}
		return -1
	}

	res := make([][]byte, 0, 32)

	i := 0
	for len(path) > 0 {
		i = index(path)
		if i == 0 {
			i++
			path = path[i:]
			continue
		}
		var dir []byte
		if i < 0 {
			dir = path
			path = nil
		} else {
			dir = path[0:i]
			i++
			path = path[i:]
		}

		if string(dir) == "." {
			continue
		} else if string(dir) == ".." {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
			continue
		}
		res = append(res, dir)
	}
	return res
}

func simplifyPath(path string) string {
	res := split([]byte(path))
	if len(res) == 0 {
		return "/"
	}
	sb := &strings.Builder{}

	for i := 0; i < len(res); i++ {
		sb.WriteByte('/')
		sb.Write(res[i])
	}
	return sb.String()
}
