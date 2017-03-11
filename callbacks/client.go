package main

/*
#include "binding.h"
*/
import "C"
import "fmt"

type Person struct {
    age int
    name string
}

func simple_callback(arg int, data interface{}) int {
    fmt.Println("Hello from inside our callback with argument ", arg);
    person := data.(Person)
    fmt.Printf("Person is %+v\n", person)
    return -arg;
}

func main() {
    fmt.Println("Will call without callback set");
    fmt.Println("Result is ", C.call_callback(1));
    fmt.Println();

    fmt.Println("Setting callback");
    person := Person{32, "Lauro"}
    set_callback(simple_callback, person);
}
