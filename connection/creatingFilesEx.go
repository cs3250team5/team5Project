package connection

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) { //check error func
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Create("myFile.txt") // creation of a file
	check(err)

	data := []byte("Hello World\n") // byte array data placed in the file

	file.Write(data)
	file.WriteString("This is a message\n") // we can use input text to edit messages or subject names

	buffer := bufio.NewWriter(file) // implementing buffer io

	buffer.WriteString("What's going on?\n")
	buffer.Flush() // clears the buffer and writes the string to the file

	data2 := make([]byte, 20)
	file.Seek(20, 0) // this tells the read how long to read
	file.Read(data2) // the file will be read up to 20 bytes, from the array above
	fmt.Printf(string(data2))

	file.Close()

}
