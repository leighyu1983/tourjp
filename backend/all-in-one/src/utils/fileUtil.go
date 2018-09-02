package util

import (
	"mime/multipart"
	"io"
	"os"
	"fmt"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func SaveFile(file multipart.File, path string, filename string) (error){

	exist, err := PathExists(path)
	if(err != nil) {
		panic(err)
	}

    if(!exist) {
		err := os.Mkdir(path, os.ModePerm)
		if(err != nil) {
			panic(err)
		}
	}

	out, _ := os.Create(path + filename);
	defer out.Close()

	fmt.Printf("[util.SaveFile] prepare to write into file '%s'\n", path + filename)
	_, err = io.Copy(out, file)
    if err != nil {
		//fmt.Println("=====ccccccccccccc=====" + path)
		//fmt.Println(fmt.Printf("%s", err))
		//fmt.Println("=====dddddddddddd=====" + path)
		panic(err)
	}
	fmt.Printf("[util.SaveFile] successfully write into file '%s'\n", path + filename)
	return nil
}