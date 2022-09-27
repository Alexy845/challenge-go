package main

import (
    "os"

    "fmt"
)

func main() {
    if len(os.Args) == 1 || os.Args[1] == "--help" || os.Args[1] == "-h" {
        fmt.Println("--insert
        -i
               This flag inserts the string into the string passed as argument.
      --order
        -o
               This flag will behave like a boolean, if it is called it will order the argument.")
        return
    }
    args := os.Args[1:]
    insert := false
    toInsert := ""
    var str string
    var nb []int
    order := false
    if len(args) > 1 {
        if (len(args[0]) > 10 && args[0][0:9] == "--insert") || (len(args[0]) > 4 && args[0][0:3] == "-i=") {
            insert = true
            toInsert = args[0][Index(args[0], "=")+1:]
        }
        if (args[0] == "--order" || args[0] == "-o") || (args[1] == "--order" || args[1] == "-o") {
            order = true
        }
    } else {
        str = args[0]
    }
    if insert {
        if order {
            str = args[2] + toInsert
        } else {
            str = args[1] + toInsert
        }
    }
    if order {
        if !insert {
            str = args[1]
        }
        for _, c := range str {
            nb = append(nb, int(rune(c)))
        }
        SortIntegerTable(nb)
        str = ""
        for _, v := range nb {
            str += string(rune(v))
        }
    }
    fmt.Println(str)
}

func SortIntegerTable(table []int) {
    for i := 0; i < len(table); i++ {
        for j := i + 1; j < len(table); j++ {
            if table[i] > table[j] {
                tmp := table[i]
                table[i] = table[j]
                table[j] = tmp
            }
        }
    }
}

func Index(s string, toFind string) int {
    if len(toFind) == 0 || len(s) == 0 {
        return 0
    }
    for v := range s {
        if v+len(toFind) >= len(s) {
            break
        }
        var slice string = s[v : v+len(toFind)]
        if slice == toFind {
            return v
        }
    }
    return -1
}