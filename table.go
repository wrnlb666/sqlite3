package main

/*
#include <stdlib.h>
#include <stdint.h>
#include <sqlite3.h>


*/
import "C"
import (
    "fmt"
    // "unsafe"
    // "runtime"
)

type Table struct {
    names   []string    // name slice for the attributes
    result  [][]Elem    // result of the sql
}

// sqlite3 itself has a `sqlite3_get_table` function, but is just a wrapper of `sqlite3_exec`
// So I am just going to create my own `GetTable` wrapper in go
func (db *Sqlite3) GetTable(sql string) (*Table, error) {
    res := &Table{}
    i := 0
    err := db.Exec(sql,
        func(row Row) int {
            if i == 0 {
                res.names = row.fields
                res.result = append(res.result, row.elems)
            } else {
                res.result = append(res.result, row.elems)
            }
            i += 1
            return 0
        })
    return res, err
}

func (t *Table) Cols() int {
    return len(t.names)
}

func (t *Table) Rows() int {
    return len(t.result)
}

func (t *Table) Fields() []string {
    return t.names
}

func (t *Table) Row(r int) (Row, error) {
    if r >= len(t.result) {
        return Row{}, fmt.Errorf("index out of bounds")
    }
    return Row{
        size:   len(t.names),
        fields: t.names,
        elems:  t.result[r],
    }, nil
}
