// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package io_ioutil

import (
	"bytes"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func getFile(t *testing.T, fileName string) *os.File {
	//注意直接用open打开是只读模式，如果需要写入用openfile函数
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("open file err ", err)
		file, err = os.Create(fileName)
		if err != nil {
			t.Error("create file err ", err)
		}
	}
	return file
}

func Test_Copy(t *testing.T) {
	Convey(
		"从 src读到eof不会出现error ", t, func() {
			fileName := "liuhao.txt"
			//注意直接用open打开是只读模式，如果需要写入用openfile函数
			file := getFile(t, fileName)
			//这里是直接抛出panic，抛出的panic只会进入defer的recover函数
			//defer file.Close()
			defer func() {
				if err := recover(); err != nil {
					file.Close()
					panic(err)
				}
				file.Close()
			}()
			n, err := file.WriteAt([]byte("hello沙发斯蒂芬 liu阿斯顿发斯蒂芬hao"), 10)
			file.Seek(0, io.SeekEnd)
			file.WriteString("hahahaha")

			if err != nil {
				t.Error(n, err)
				t.FailNow()
			}
			file.Seek(0, io.SeekStart)
			written, err := io.Copy(os.Stdout, file)
			fmt.Println(written, err)
			i, err := io.CopyN(os.Stdout, strings.NewReader("Go语言中文网"), 100)
			fmt.Println(i, err)
		},
	)
}

func TestMultiRW(t *testing.T) {
	Convey(
		"mutil reader", t, func() {
			readers := []io.Reader{
				strings.NewReader("from strings reader"),
				bytes.NewBufferString("from bytes buffer"),
			}
			reader := io.MultiReader(readers...)
			data := make([]byte, 0, 128)
			buf := make([]byte, 10)

			for n, err := reader.Read(buf); err != io.EOF; n, err = reader.Read(buf) {
				if err != nil {
					panic(err)
				}
				data = append(data, buf[:n]...)
			}
			fmt.Printf("%s\n", data)

		},
	)
	//go (Issue 18232)
	//mr1 := io.MultiReader(strings.NewReader("def"), strings.NewReader("ghi"))
	//mr2 := io.MultiReader(strings.NewReader("abc"), mr1)
	//
	//io.ReadFull(mr2, make([]byte, 4))
	//io.ReadFull(mr1, make([]byte, 4))
	//io.ReadFull(mr2, make([]byte, 3))

	Convey(
		"mutil write", t, func() {
			file := getFile(t, "liuhao")
			writer := io.MultiWriter(file, os.Stdout)
			writer.Write([]byte("刘号"))
		},
	)
	// 将读者写者封装起来，从读者写入写者。所有的io的操作都是从读者写入写着，有的方法接受者是读者有的方法接受者是写着，有的通过管道传输
	//io.TeeReader
}

func Test_IoutilRead(t*testing.T){
	Convey(
		"readall 是对io包read 增强了一种常见用法",t,func(){

			all, e := ioutil.ReadAll(getFile(t,"liuhao"))
			fmt.Println(all,e)
		},
		)
	Convey(
		"readdir 读该目录层次下的所有文件或目录，而不继续递归查询",t,func(){

			infos, e := ioutil.ReadDir(`D:\GO\src\prcitf\io操作与文本文件处理`)
			fmt.Println(infos,e)
			for _,info:= range infos{
				fmt.Println( info.Name(),info.Size(),info.Mode(),info.ModTime(),info.Sys(),info.IsDir())
			}
		},
	)
}

func Test_IoutilFile(t *testing.T){
	file, e := ioutil.ReadFile("liuhao")
	fmt.Println(file,e)
	e= ioutil.WriteFile("liuhao", []byte(" 我要学习"), 0666)
}
