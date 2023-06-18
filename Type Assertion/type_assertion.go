package main

import "fmt"

func main() {
    var i interface{} = 42

    // Type assertion to extract underlying int value
    num, ok := i.(int)
    if ok {
        fmt.Println("The value is an integer:", num)
    } else {
        fmt.Println("The value is not an integer")
    }

    // Type assertion to extract underlying string value
    str, ok := i.(string)
    if ok {
        fmt.Println("The value is a string:", str)
    } else {
        fmt.Println("The value is not a string")
    }
}
