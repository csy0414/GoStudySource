package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

// FILE_NAME_FORMAT is ...
const FILE_NAME_FORMAT = "200601021504"
const LOG_TIME_FORMAT = "20060102150405"

func main() {
	var fpLogFile *os.File
	var bufWriter *bufio.Writer

	sPort, slogPath := getArg()

	udpAddr, err := net.ResolveUDPAddr("udp", ":"+sPort)
	checkError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	defer conn.Close()

	fmt.Println("ListenUDP...")

	t := time.Now()
	sFileTime := t.Format(FILE_NAME_FORMAT)
	fpLogFile, err = os.OpenFile(
		slogPath+"\\"+sFileTime+".log",
		os.O_WRONLY|os.O_APPEND|os.O_CREATE,
		os.FileMode(700))
	if err != nil {
		fmt.Println(err)
		return
	}

	var buf = make([]byte, 1024)

	bufWriter = bufio.NewWriterSize(fpLogFile, 1024)

	for {
		err := conn.SetReadBuffer(1024 * 1024 * 10)
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			return
		}
		t := time.Now()
		sRecvTime := t.Format(FILE_NAME_FORMAT)

		if sFileTime != sRecvTime {
			fpLogFile.Close() // 기존 로그파일 Close
			sFileTime = sRecvTime

			fpLogFile, err := os.OpenFile(
				slogPath+"\\"+sFileTime+".log",
				os.O_WRONLY|os.O_APPEND|os.O_CREATE,
				os.FileMode(700))
			if err != nil {
				fmt.Println(err)
				return
			}
			defer fpLogFile.Close()
			bufWriter = bufio.NewWriterSize(fpLogFile, 1024)
		}

		sLogRecTime := t.Format(LOG_TIME_FORMAT)
		usedBufSize := bufWriter.Available()
		bufWriter.WriteString(fmt.Sprintf("%s|%s|%s\n", addr.IP, sLogRecTime, string(buf[0:n])))

		fmt.Printf("bufsize=%d\n", usedBufSize)
	}
}

func getArg() (string, string) {
	var sPort, slogPath string

	if len(os.Args) < 2 {
		fmt.Println("Usage : udpServer [Port] [LogPath]")
		os.Exit(-1)
	}

	sPort = os.Args[1]
	slogPath = os.Args[2]

	return sPort, slogPath
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
		os.Exit(1)
	}
}
