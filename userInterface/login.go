package userInterface

import (
    "fmt"
    "bufio"
    "os"    
    "strings"
)

func GetUsername() (un string) {
    fmt.Print("Enter thy username: ")
    
    reader := bufio.NewReader(os.Stdin)
    un, _ = reader.ReadString('\n')
    return
}

func GetPassword() (pw string){
    fmt.Print("Enter thy password: ")
    reader := bufio.NewReader(os.Stdin)
    pw, _ = reader.ReadString('\n')
    return
}

func GetUsernameAndPassword() (un, pw string) {
    un = trim(GetUsername())
    pw = trim(GetPassword())
    return
}

func trim(s string) string {
    return strings.Replace(s, "\n", "", -1)
}
