package main

/*
#include <stdio.h>
#include <stdlib.h>

// Wrapper around variadic printf.
void myprintf(const char *msg) {
    printf("%s", msg);
}
*/
import "C"
import "unsafe"

func main() {
    cstr := C.CString("Hello, cgo!\n")
    defer C.free(unsafe.Pointer(cstr))

    C.myprintf(cstr)
}
