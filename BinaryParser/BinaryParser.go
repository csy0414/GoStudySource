//Package BinaryParser is ...
//package main
package BinaryParser

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	UTLINENAME = 32
	UTNAMESIZE = 32
	UTHOSTSIZE = 256
)

// EXITSTATUS struct is Defined from utmp.h
type EXITSTATUS struct {
	Etermination int16
	Eexit        int16
}

// TIMEVAL struct is Defined from time.h
type TIMEVAL struct {
	Tvsec  int32
	Tvusec int32
}

// UTMP struct is Defined from utmp.h
type UTMP struct {
	Uttype    int32
	Utpid     int32
	Utline    [UTLINENAME]byte
	Utid      [4]byte
	Utuser    [UTNAMESIZE]byte
	Uthost    [UTHOSTSIZE]byte
	Utexit    EXITSTATUS
	Utsession int32
	Uttv      TIMEVAL
	Utaddrv6  [4]int32
	Unused    [20]byte
}

/*
func main() {
	arg1, arg2, argc, err := GetArg(os.Args)

	if argc == 2 {

	} else if argc == 3 {

	} else {
		return
	}

	return
}
*/
//GetArg parse command from user input.
func GetArg(cmd []string) (string, string, int, error) {

	argc := len(cmd)

	if argc == 1 {
		fmt.Println("Command input error. :", cmd)
		fmt.Println("Usage : # cmd [OPTION -t]... [FILE]... ")
		return "", "", 0, errors.New("func GetArg err")
	} else if argc == 2 {
		return cmd[1], "", argc, nil
	} else if argc == 3 {
		return cmd[1], cmd[2], argc, nil
	}

	fmt.Println("Command input error. :", cmd)
	fmt.Println("Usage : # cmd [OPTION -t]... [FILE]... ")
	return "", "", 0, errors.New("func GetArg err")
}

//BinaryFileOpen open and return *FILE from binary file path.
func BinaryFileOpen(path string) *os.File {

	pwtmp, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return pwtmp
}

//BinaryReadAll Read From beginning to end.
func BinaryReadAll(pwtmp *os.File) bool {
	pwtmp.Seek(0, os.SEEK_SET)
	for {
		ut := UTMP{}
		err := binary.Read(pwtmp, binary.LittleEndian, &ut)
		if err != nil {
			if err == io.EOF {
				return true
			}
			return false
		}

		fmt.Printf("ut_type=%d;ut_pid=%d;ut_line=%s;ut_id=%s;ut_user=%s;ut_host=%s;e_termination=%d;e_exit=%dut_session=%d;tv_sec=%d;tv_usec=%d;\n",
			ut.Uttype,
			ut.Utpid,
			strings.Trim(string(ut.Utline[:]), "\x00"),
			ut.Utid,
			strings.Trim(string(ut.Utuser[:]), "\x00"),
			strings.Trim(string(ut.Uthost[:]), "\x00"),
			ut.Utexit.Etermination,
			ut.Utexit.Eexit,
			ut.Utsession,
			ut.Uttv.Tvsec,
			ut.Uttv.Tvusec)
	}
}

//BinaryReadTail Read From beginning to end.
func BinaryReadTail(pwtmp *os.File) bool {
	pwtmp.Seek(0, os.SEEK_SET)
	return true
}
