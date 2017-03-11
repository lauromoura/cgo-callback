package main

/*
#include "helper.h"
*/
import "C"
import "fmt"

//export Exported_To_C
func Exported_To_C(x int) {
    fmt.Println("Exported_To_C: Received ", x)
}

func main() {
    C.call_exported(42);
    C.call_exported(71);
}
