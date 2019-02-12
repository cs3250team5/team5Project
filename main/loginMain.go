package main

import (
    "fmt"
    "Team5Project/userInterface"
)
func main() {


    un, pw := userInterface.GetUsernameAndPassword()
    fmt.Println("got: ", un, ", ", pw)
}
