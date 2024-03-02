package main

import (
    "fmt"
    "log"
)

func main() {
    db, err := Open("db.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()


    table, err := db.GetTable(`select datetime('now');`)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(table)
}
