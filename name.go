package core

import (
	"math/rand"
	"strings"
	"os"
)

// TrimApart used to delete os.PathSepartor in the path suffix
func TrimApart(path string) string {


	if strings.HasSuffix(path, string(os.PathSeparator)) {

		return path[:len(path)-1]
	}
	return path
}

// AddApart used to add os.PathSeparator in the path suffix
func AddApart(path string) string {

	apart := string(os.PathSeparator)
	if strings.HasSuffix(path, apart) {

		return path
	}
	return path + apart
}

// ParentDir used to get parent dir string
func ParentDir(path string) string {

	path = TrimApart(path)

	idx := strings.LastIndexByte(path, os.PathSeparator)
	if idx == -1 {
		return ""
	}

	return path[:idx]
}

// RandBytes 用于生成一组随机字，字母及数字
func RandBytes(num int) []byte {

	buf := make([]byte, num)

	for k := range buf {

		idx := rand.Intn(52)
		var base int
		switch {
		case idx < 26: base = 65
		case idx < 52: base = 97 - 26
		default: base = 48 - 52
		}

		buf[k] = byte(base + idx)
	}
	return buf
}
