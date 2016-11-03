package BinaryParser

import "testing"

func TestGetArg(t *testing.T) {
	var testdata = []struct {
		command []string
		arg1    string
		arg2    string
		argc    int
	}{
		{[]string{"command"}, "", "", 0},
		{[]string{"command", "\\var\\log\\wtmp"}, "\\var\\log\\wtmp", "", 2},
		{[]string{"command", "-t", "\\var\\log\\wtmp"}, "-t", "\\var\\log\\wtmp", 3},
	}

	for _, test := range testdata {
		arg1, arg2, argc, _ := GetArg(test.command)
		if (test.arg1 != arg1) || (test.arg2 != arg2) || (test.argc != argc) {
			t.Errorf("GetArg(%s) -> %s:%s / %s:%s / %d:%d", test.command, test.arg1, arg1, test.arg2, arg2, test.argc, argc)
		}

	}
}

func TestBinaryFileOpen(t *testing.T) {
	var testpath = []string{
		"wtmp",
	}

	for _, path := range testpath {
		pFile := BinaryFileOpen(path)
		if pFile == nil {
			t.Errorf("Fail Open File : %s", path)
		}
	}
}

func TestBinaryReadAll(t *testing.T) {

	pFile := BinaryFileOpen("wtmp")
	r := BinaryReadAll(pFile)
	if r != true {
		t.Errorf("BinaryAllError. : %s", pFile.Name)
	}
}

func TestBinaryReadTail(t *testing.T) {

	pFile := BinaryFileOpen("wtmp")
	r := BinaryReadAll(pFile)
	if r != true {
		t.Errorf("BinaryAllError. : %s", pFile.Name)
	}
}
