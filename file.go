package core

import (
	"io"
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

	dirLen := len(src)
	for {

		name := d.NextFile()
		if name == "" {
			break
		}

		rName := dst + name[dirLen:]
		DuplicateFile(name, rName)
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

	MkPathDir(rName)
	dst, err := os.OpenFile(rName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("error open dstfile :" + err.Error())
		return
	}
	defer dst.Close()
	io.Copy(dst, src)
}

// MkPathDir used to make dir of path
func MkPathDir(path string) {

	dirName := ParentDir(path)
	if _, e := os.Stat(dirName); e != nil {

		os.MkdirAll(dirName, 0755)
	}
}
