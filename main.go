package main

import (
	"archive/zip"
	"fmt"
	"github.com/atotto/clipboard"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	licensedUrl = "http://idea.medeming.com/jets/images/jihuoma.zip"
	zipFile     = "idea.medeming.com.zip"
)

func main() {

	// 下载文件内容并保存
	response, err := http.Get(licensedUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	os.Remove(zipFile)
	ioutil.WriteFile(zipFile, data, 0644)

	// 获取压缩包内容
	r, err := zip.OpenReader(zipFile)

	if err != nil {
		panic(err)
	}

	var license string
	for _, file := range r.File {
		if file.Name[0:4] != "2018"{
			continue
		}
		f, err := file.Open()

		if err != nil {
			panic(err)
		}
		fd, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		err = f.Close()
		if err != nil {
			panic(err)
		}
		license = string(fd)
	}

	fmt.Println(license)

	// 拷贝到粘贴板
	clipboard.WriteAll(license)

	r.Close()
	if err != nil {
		panic(err)
	}
	dir, err := os.Getwd()
	err = os.Remove(dir + "/" + zipFile)
	if err != nil {
		panic(err)
	}

	// 设置到ide licensed中

}
