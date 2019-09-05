package io

import (
	"fmt"
	"io" // 最基本的接口
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

/*
  Go io包提供最基础两个接口 Reader 和 Writer 接口
 */

func ReaderExample() {
	//读取文件内容
	// filedata, fileerror := ioutil.ReadFile("/Users/mac/dev/goio.text")
	//fmt.Println("file date:",string(filedata))
	// fmt.Println("error info:",fileerror)

	// 字符串读取
	data, err := ReaderForm(strings.NewReader("/Users/mac/dev/goio.text"), 15)
	// 将byte数组转为字符
	fmt.Println("data:", string(data))
	fmt.Println("error info:", err)

	for idx, args := range os.Args {
		fmt.Println("参数" + strconv.Itoa(idx) + ":", args)
	}

	dir:= os.Args[0]
	ReaderListDir(dir,0)

}



// 可以读取任意地方数据
func ReaderForm(reader io.Reader,num int)([]byte,error){
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

// 读取 某目录所有文件
func ReaderListDir(path string,curHier int){
	fileInfo,err := ioutil.ReadDir(path)
	if err!= nil{
		fmt.Println("err:",err)
		return
	}

	for _,info := range fileInfo {
		if info.IsDir(){
			for tmpHier :=curHier;tmpHier>0;tmpHier-- {
				fmt.Println("|\t")
			}
			fmt.Printf(info.Name(),"\\")
			ReaderListDir(path+"\\"+info.Name(),curHier+1)
		} else {
			for tmpHier :=curHier;tmpHier>0 ; tmpHier--  {
				fmt.Println("|\t")
			}
			fmt.Println(info.Name())
		}
	}

}

