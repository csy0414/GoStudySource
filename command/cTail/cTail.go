package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var ReadSize int64
	var notReadSize int64

	rfile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rfile.Close()

	fi, err := rfile.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	ReadSize = fi.Size()

	for {
		time.Sleep(500 * time.Millisecond)

		fi, err := rfile.Stat()
		if err != nil {
			fmt.Println(err)
			return
		}

		if ReadSize < fi.Size() {
			notReadSize = fi.Size() - ReadSize
			rfile.Seek(-notReadSize, os.SEEK_END)

			var data = make([]byte, notReadSize)
			_, err = rfile.Read(data)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Print(string(data))
			ReadSize = fi.Size()
		}
	}
}
