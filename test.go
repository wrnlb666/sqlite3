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

    err = db.Exec(`
        select datetime('now');
        `,
        func (argc int, res Iterator) bool {
            for i, v := range res() {
                fmt.Println(i, ": ", v)
            }
            return true
        })
    if err != nil {
        log.Fatal(err)
    }

    var names []string
    var result [][]string
    i := 0
    err = db.Exec(`select datetime('now')`,
        func(argc int, argv Iterator) bool {
            temp := make([]string, argc)
            j := 0
            if i == 0 {
                names = make([]string, argc)
                for k, v := range argv() {
                    names[j] = k
                    temp[j] = v
                    j += 1
                }
            } else {
                for _, v := range argv() {
                    temp[j] = v
                    j += 1
                }
            }
            result = append(result, temp)
            i += 1
            return true
        })
    fmt.Println(names)
    fmt.Println(result)
}
