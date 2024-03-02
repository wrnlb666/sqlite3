package main

/*
#include <stdlib.h>
#include <sqlite3.h>
*/
import "C"
import (
    "strconv"
)

type Elem string

func (e Elem) String() string {
    return string(e)
}

func (e Elem) Int() (int, error) {
    res, err := strconv.Atoi(string(e))
    if err != nil {
        return *new(int), err
    }
    return res, nil
}

func (e Elem) IntOr(or int) int {
    res, err := strconv.Atoi(string(e))
    if err != nil {
        return or
    }
    return res
}

func (e Elem) Float() (float64, error) {
    res, err := strconv.ParseFloat(string(e), 64)
    if err != nil {
        return *new(float64), err
    }
    return res, nil
}

func (e Elem) FloatOr(or float64) float64 {
    res, err := strconv.ParseFloat(string(e), 64)
    if err != nil {
        return or
    }
    return res
}
