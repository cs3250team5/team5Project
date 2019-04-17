package smtp

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var name, choice, clean string

	fmt.Print("Enter your name:")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello", name)
	fmt.Print("Now type out your message and finish with a single line with '.': ")
	reader := bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')
	fmt.Print("Would you like to save as a draft or send the message?// S/s or D/d")
	fmt.Scanf("%s", &choice)
	if choice == "S" || choice == "s" {
		fmt.Print("Message Sent ")
	}
	if choice == "D" || choice == "d" {
		fmt.Print("Message saved as a Draft ")
	}
	clean = CleanMessage(message)
	fmt.Print(strings.TrimSpace(clean))
}

func CleanMessage(s string) string {
	s = strings.TrimSuffix(s, "\n.")
	return s
}
