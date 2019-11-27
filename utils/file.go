package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 存储文件信息
type File struct {
	ReadpPath string
	WritePath string
	Content   string
}

// 加载文件 并 初始化 文件
func LoadFile(readFile, writePath string) (file *File, err error) {
	//获取文件信息
	fi, err := os.Stat(readFile)
	if err != nil {
		return
	}
	//检测是否是一个目录
	if fi.IsDir() {
		err = fmt.Errorf(readFile + " is not a file.")
		return
	}

	b, err := ioutil.ReadFile(readFile)
	if err != nil {
		return
	}
	file = &File{
		ReadpPath: readFile,
		WritePath: writePath,
		Content:   string(b),
	}
	return
}

// 写入文件
func (file *File) WriteFile() (err error) {
	f, err := os.OpenFile(file.WritePath, os.O_RDWR|os.O_CREATE, 0600)
	defer f.Close()
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(f)
	_, err = writer.Write([]byte(file.Content))
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}
