package core

import (
	// "runtime"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// DirInfo used to manage director
type DirInfo struct {
	_focus   int
	_fileArr []string

	_root 	 string		//the root dir full path
	_filter  Filter

	// _curArr  []string
	// _curIdx  int
}

// NewDirInfo used create en object of DirInfo with root dir
func NewDirInfo(root string, filter Filter) (*DirInfo, error) {

	// apart := string(os.PathSeparator)
	// if !strings.HasSuffix(root, apart) {
	// 	root += apart
	// }
	root = AddApart(root)

	info := new(DirInfo)
	*info = DirInfo{
		_focus:   0,
		_filter: filter,
		_fileArr: nil,
		_root: root,
	}

	err := filepath.Walk(root, info.handleWalk)
	if err != nil {

		fmt.Println("there's an error in dict reset : " + err.Error())
	}
	return info, err
}


// FilesIn used to get all filenames in director dirName
// dirName maybe a full path
func (info *DirInfo) FilesIn(dirName string) []*string {

	if !strings.HasPrefix(dirName, info._root) {
		dirName = info._root + dirName
	}

	dirName = TrimApart(dirName)
	
	arr := make([]*string, 0, 10)
	dirLen, focus := len(dirName), false
	for _, v := range info._fileArr {

		if strings.HasPrefix(v, dirName) && (-1 == strings.IndexByte(v[dirLen+1:], os.PathSeparator)) {

			focus = true
			arr = append(arr, &v)			
		} else if focus {
			
			break
		}
	}
	return arr
}

// NextFile used to pop next file name
func (info *DirInfo) NextFile() string {

	var name string
	if info._focus < len(info._fileArr) {

		name = info._fileArr[info._focus]
		info._focus++
	}
	return name
}

func (info *DirInfo) handleWalk(path string, file os.FileInfo, err error) error {

	if file == nil {
		return err
	}

	if strings.HasSuffix(path, string(os.PathSeparator)) {
		return nil
	}

	switch {
	case file.Mode()&os.ModeSymlink > 0:
	case strings.HasPrefix(file.Name(), "."):
	case file.IsDir():
	default:
		info.addFile(path)
	}

	return nil
}

func (info *DirInfo) addFile(path string) {

	if info._filter != nil && !info._filter.Match(&path) {
		return
	}

	info._fileArr = append(info._fileArr, path)
	// info._fileCnt++
}

// var dirApart byte
// // DirApartByte is the function to get the byte of the dir apart
// func DirApartByte() byte {

// 	if dirApart == 0 {

// 		switch runtime.GOOS {
// 			case "windows":	dirApart = '\\'
// 			default:		dirApart = '/'
// 		}
// 	}
// 	return dirApart
// }
