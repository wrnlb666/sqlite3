package main

/*
#cgo pkg-config: sqlite3 
#include <stdlib.h>
#include <sqlite3.h>
*/
import "C"
import (
    "unsafe"
    "runtime"
)

type Sqlite3 C.sqlite3

func Open(filename string) (*Sqlite3, error) {
    // Cdb := &C.sqlite3{}
    var Cdb *C.sqlite3
    Cfilename := C.CString(filename)
    defer C.free(unsafe.Pointer(Cfilename))
    r := C.sqlite3_open(Cfilename, &Cdb)
    if r != C.int(0) {
        return nil, newError(r, "")
    }
    db := (*Sqlite3)(Cdb)
    runtime.SetFinalizer(db, func(d *Sqlite3) {
        d.Close()
    })
    return db, nil
}

func (db *Sqlite3) Close() error {
    r := C.sqlite3_close((*C.sqlite3)(db))
    if r != C.int(0) {
        return newError(r, "")
    }
    return nil
}
