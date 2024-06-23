package main

import "C"

func ConvertC2Go(input *C.char) string {
	return C.GoString(input)
}

func ConvertGo2C(input string) *C.char {
	return C.CString(input)
}
