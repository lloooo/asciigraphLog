package main

import (
	"fmt"
	"./src/github.com/guptarohit/asciigraph"
	"os"
	"io/ioutil"
	"strings"
	"time"
	"log"
	"bufio"
	"io"
)

var (
	data    []float64
	fileCnt = 0
	dirpath = os.Args[1]
	suffix  = os.Args[2]
)

func main() {
	for {
		fetchData()
		graph := asciigraph.Plot(data, asciigraph.Height(2))

		fmt.Println(graph)
		time.Sleep(2 * time.Second)
	}
}

func initData() {
	data = []float64{}
	files, err := ListDir(dirpath, suffix)
	if err != nil {
		fmt.Println(err)
	}
	fileCnt = len(files)
	for _, file := range files {
		lineCnt := float64(computeLine(file))
		data = append(data, lineCnt)
	}
}

func fetchData() {
	files, err := ListDir(dirpath, suffix)
	if err != nil {
		fmt.Println(err)
	}
	nowFileCnt := len(files)
	if nowFileCnt != fileCnt {
		initData()
	}
	lastFile := files[fileCnt-1]
	data[len(data)-1] = float64(computeLine(lastFile))
}

func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = []string{}
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

func computeLine(path string) (cnt int) {
	num := 0
	f, err := os.Open(path)
	if nil != err {
		log.Println(err)
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		_, err := r.ReadString('\n')
		if io.EOF == err || nil != err {
			break
		}
		num += 1
	}
	return num
}
