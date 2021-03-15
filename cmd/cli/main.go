package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/quycao/xstats/pkg/handler"
)

func main() {
	fmt.Print("Enter text: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")

	// TCBS
	result, err := handler.StatsTCBS(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", result)
	}

	// BVSC
	// result, err := handler.StatsBVSC(input)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%+v\n", result)
	// }
}
