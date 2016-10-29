package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//ExcuteName FileName String
	if len(os.Args) == 3 {
		rfile, err := os.Open(os.Args[1])
		if err != nil {
			return
		}
		scanner := bufio.NewScanner(rfile)
		for scanner.Scan() {
			readStr := scanner.Text()
			if strings.Contains(readStr, os.Args[2]) != false {
				fmt.Println(readStr)
			}
		}
	} else { //ExcuteName String
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			readStr := scanner.Text()
			//fmt.Println(readStr)
			if strings.Contains(readStr, os.Args[1]) != false {
				fmt.Println(readStr)
			}
		}
	}
}
