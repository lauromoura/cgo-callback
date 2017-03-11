package main

/*
#include "binding.h"
*/
import "C"
// import "fmt"
import "sync"

// Callback infrastructure code
type cb_type func(int, interface{}) int;

type callback_info struct {
    callback cb_type
    data interface{}
}

//export go_callback_int
func go_callback_int(key, value int) int {
    callback_data := lookup(key);
    return callback_data.callback(value, callback_data.data);
}

var mu sync.Mutex;
var index int;
var fns = make(map[int]*callback_info)

func lookup(i int) callback_info {
    mu.Lock()
    defer mu.Unlock()
    return *fns[index];
}

func set_callback(callback cb_type, data interface{}) int {
    mu.Lock();
    defer mu.Unlock();

    index++;
    for fns[index] != nil {
        index++;
    }

    fns[index] = &callback_info{callback, data};

    C.set_translator(C.long(index));
    return index;
}


