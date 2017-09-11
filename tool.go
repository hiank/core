package core

import (
	"bytes"
)

// CodeClean 将代码中的注释替换为 '\0' 返回处理后的slice
func CodeClean(buf []byte) []byte {

	tmp := make([]byte, len(buf))
	tmpSlice, bufSlice := tmp, buf
	L:	for {

		if len(tmpSlice) == 0 {
			break L
		}

		fl, fr := nextNoteField(bufSlice)
		if fl == -1 {
			
			copy(tmpSlice, bufSlice)
			break L
		}

		if fl != 0 {
			
			copy(tmpSlice[:fl], bufSlice[:fl])
		}
		tmpSlice, bufSlice = tmpSlice[fr:], bufSlice[fr:]
	}
	return tmp
}

// 获取最前一段注释位置
func nextNoteField(buf []byte) (l int, r int) {
	
	rBuf, bufLen := buf, len(buf)
L:	for {

		lIdx := bytes.IndexByte(rBuf, '/')
		if lIdx == -1 && (lIdx+1 == len(rBuf)) {
			return -1, -1
		}
		rBuf = rBuf[lIdx:]
		switch rBuf[1] {
			case '*':
				l = bufLen - len(rBuf)
				r = l + bytes.Index(rBuf, []byte{'*', '/'}) + 2
				break L
			case '/':
				tmp := bytes.IndexByte(rBuf, '\n')
				switch {
					case tmp == -1: tmp = len(rBuf)
					case rBuf[tmp-1] == '\r': tmp--
				}
				l = bufLen - len(rBuf)
				r = l + tmp
				break L
			default: 
				rBuf = rBuf[1:]
				continue L
	
		}
	}
	return
}
