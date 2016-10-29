package main

import (
	"fmt"
	"os"
)

func main() {

	// File Open A, B
	rfile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rfile.Close()

	wfile, err := os.OpenFile(
		os.Args[2],
		os.O_WRONLY|os.O_CREATE,
		os.FileMode(700))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wfile.Close()
	////////////////////////////////

	fi, err := rfile.Stat() //파일 정보 가져오기
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fi.Size()) // 파일 크기만큼 바이트 슬라이스 생성
	rfile.Seek(0, os.SEEK_SET)

	//Read A file
	_, err = rfile.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = wfile.Write([]byte(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Copy Success")

}
