package connection

import (
	"fmt"
	"strconv"
	"strings"
)

func extractFromList(s string) map[int]int {
	lines := strings.Split(s, "\n")
	listMap := make(map[int]int)
	for _, line := range lines[1 : len(lines)-2] {
		y := strings.Split(line, " ")
		int1, err := strconv.Atoi(y[0])
		int2, err := strconv.Atoi(y[1])
		if err == nil {
			listMap[int1] = int2
		}

		fmt.Println(listMap)
	}

	return listMap
}
