package main

import (
	"Team5Project/connection"
	"fmt"
)

func main() {
	mail := connection.ReadMF("3_Test_3_Frederick_Buker.txt")
	fmt.Println(mail.To, mail.Message)
}
