package main

/*
#include <stdlib.h>
#include <stdint.h>
#include <sqlite3.h>


*/
import "C"
import (
    // "unsafe"
    // "runtime"
)

type Table struct {
    names   []string    // name slice for the attributes
    result  [][]string  // result of the sql
}

// sqlite3 itself has a `sqlite3_get_table` function, but is just a wrapper of `sqlite3_exec`
// So I am just going to create my own `GetTable` wrapper in go
func (db *Sqlite3) GetTable(sql string) (Table, error) {
    res := Table{}
    i := 0
    err := db.Exec(sql,
        func(argc int, argv Iterator) bool {
            temp := make([]string, argc)
            j := 0
            if i == 0 {
                res.names = make([]string, argc)
                for k, v := range argv() {
                    res.names[j] = k
                    temp[j] = v
                    j += 1
                }
            } else {
                for _, v := range argv() {
                    temp[j] = v
                    j += 1
                }
            }
            res.result = append(res.result, temp)
            i += 1
            return true
        })
    return res, err
}
