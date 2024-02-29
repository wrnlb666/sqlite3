package main

/*
#include <sqlite3.h>
*/
import "C"

import (
    // "errors"
)


type Error struct {
    code    C.int
    errmsg  string
}

func newError(code C.int, errmsg string) Error {
    return Error{
        code:   code,
        errmsg: errmsg,
    }
}

func (e Error) Error() string {
    errstr := C.GoString(C.sqlite3_errstr(C.int(e.code)))
    if e.errmsg == "" {
        return errstr
    }
    return errstr + ": " + e.errmsg
}

func (e Error) Code() int {
    return int(e.code)
}

func (e Error) Sqlmsg() string {
    return e.errmsg
}
