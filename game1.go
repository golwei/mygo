package main

import "fmt"

func main() {
        fmt.Println("====")
        for i := 0; i < 5; i++ {
                for j := 0; j < i; j++ {
                        fmt.Print("*")
                }       
                fmt.Println("")
        }       
}
