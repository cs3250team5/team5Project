package connection

import (
	"fmt"
	"strconv"
	"strings"
)

func ExtractFromList(s string) map[int]int {
	// Puts message number and data size into a map
	lines := strings.Split(s, "\n")
	listMap := make(map[int]int)
	for _, line := range lines[1 : len(lines)-2] {
		y := strings.Split(line, " ")
		key, err := strconv.Atoi(y[0])
		value, err := strconv.Atoi(y[1])
		if err == nil {
			listMap[key] = value
		}
		fmt.Println(listMap)
	}
	return listMap
}
