package main

import (
    "fmt"
)

func main() {
    var a,b,n int
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&n)

    sol := solution(a,b,n)
    fmt.Println(sol)
}


func solution(a,b,n int) string {
    if n > a && n > b && a > 1 {
        return "Yes"
    }
    
    divA := a / n
    divB := b / n
    if divA > divB {
        return "Yes"
    }

    modA := a % n
    modB := b % n
    if modA > modB && modB == 0{
        return "Yes"
    }

    return "No"
}