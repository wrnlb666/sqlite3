package main

/*
#include <stdlib.h>
#include <sqlite3.h>
*/
import "C"
import (
    "iter"
)

type Row struct {
    size    int
    fields  []string
    elems   []Elem
}

func (r Row) Size() int {
    return r.size
}

func (r Row) Fields() []string {
    return r.fields
}

func (r Row) Elems() []Elem {
    return r.elems
}

func (r Row) All() iter.Seq2[string, Elem] {
    return func(yield func(string, Elem) bool) {
        for i := range r.size {
            if !yield(r.fields[i], r.elems[i]) {
                return
            }
        }
    }
}

func (r Row) Backward() iter.Seq2[string, Elem] {
    return func(yield func(string, Elem) bool) {
        for i := r.size-1; i >= 0; i-- {
            if !yield(r.fields[i], r.elems[i]) {
                return
            }
        }
    }
}
