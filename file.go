package core

import (
	"io"
	"strings"
	"os"
	"fmt"
)

// DuplicateDir used to copy dir src to dir dst
func DuplicateDir(src string, dst string) {

	src = AddApart(src)
	dst = AddApart(dst)

	d, err := NewDirInfo(src, nil)
	if err != nil {
		fmt.Println("there is en error in loadFile : " + err.Error())
		return
	}

	for {

		name := d.NextFile()
		if name == nil {
			break
		}

		idx := strings.LastIndexByte(*name, os.PathSeparator)
		dirName := (*name)[:idx]
		if _, e := os.Stat(dirName); e != nil {

			os.MkdirAll(dirName, 0755)
		}

		rName := dst + (*name)[len(src):]
		DuplicateFile(*name, rName)
	}
}

// DuplicateFile used to copy file
func DuplicateFile(lName string, rName string) {

	src, err := os.Open(lName)
	if err != nil {
		fmt.Println("error open srcfile :" + err.Error())
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(rName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("error open dstfile :" + err.Error())
		return
	}
	defer dst.Close()
	io.Copy(dst, src)
}

