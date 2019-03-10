package connection

import (
	"fmt"
	"strconv"
	"strings"
)

const ex = `4 messages:
1 3153
2 3183
3 11542
4 13946
.`

func main() {

	fmt.Println(ex)
	extractFromList(ex)
}
func extractFromList(s string) int {


	//This is were user gives input
	fmt.Println("What email would you like to seclect?")
	var userInput int
	_, err := fmt.Scanf("%d", &userInput)

	if err != nil {
        fmt.Println(err)
    }
	
	/*This is taking the big string and spliting it to create a
	  map that holds two int*/ 
	listMap := make(map[int]int)
	x := strings.Split(ex, "\n")
	y := strings.Split(x[userInput], " ")
	int1, err := strconv.Atoi(y[0])
	int2, err := strconv.Atoi(y[1])

	if err == nil {
		listMap[int1] = int2
	}
	fmt.Println(listMap[userInput])


	return userInput
}
