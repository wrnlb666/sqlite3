package main

/*
#include <stdlib.h>
#include <stdint.h>
#include <sqlite3.h>

extern int go_sqlite3_exec_cb(uintptr_t cb, int argc, char** argv, char** col);

static inline int sqlite3_exec_wrapper(sqlite3* db,
        const char* sql,
        void* callback,
        uintptr_t args,
        char** errmsg) {
    return sqlite3_exec(db,
        sql,
        (int (*)(void*, int, char**, char**))callback,
        (void*)args,
        (char**)errmsg);
}
*/
import "C"
import (
    "unsafe"
    "runtime/cgo"
)

type ExecCB func(row Row) int

func (db *Sqlite3) Exec(sql string, cb ExecCB) error {
    Cdb := (*C.sqlite3)(db)
    Csql := C.CString(sql)
    defer C.free(unsafe.Pointer(Csql))
    var errmsg *C.char
    defer C.sqlite3_free(unsafe.Pointer(errmsg))
    // pin callback closure
    Ccb := cgo.NewHandle(cb)
    defer Ccb.Delete()
    r := C.sqlite3_exec_wrapper(Cdb,
        Csql,
        C.go_sqlite3_exec_cb,
        C.uintptr_t(Ccb),
        &errmsg)
    goerrmsg := C.GoString(errmsg)
    if r != C.int(0) {
        return newError(r, goerrmsg)
    }
    return nil
}

//export go_sqlite3_exec_cb
func go_sqlite3_exec_cb(callback C.uintptr_t, argc C.int, argv **C.char, col **C.char) C.int {
    goCallback := cgo.Handle(callback).Value().(ExecCB)
    
    goArgc := int(argc)
    goArgv := uintptr(unsafe.Pointer(argv))
    goCol := uintptr(unsafe.Pointer(col))
    ptrSize := uintptr(C.sizeof_uintptr_t)

    keys := make([]string, goArgc)
    vals := make([]Elem, goArgc)
    for i := range goArgc {
        offset := ptrSize * uintptr(i)
        keys[i] = C.GoString(*(**C.char)(unsafe.Pointer(goCol + offset)))
        vals[i] = Elem(C.GoString(*(**C.char)(unsafe.Pointer(goArgv + offset))))
    }

    row := Row{
        size:   goArgc,
        fields: keys,
        elems:  vals,
    }

    return C.int(goCallback(row))
}

