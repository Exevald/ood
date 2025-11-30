package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"slides/pkg/model"
)

func main() {
	gm := model.NewGumballMachine(5)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(gm)
		fmt.Println("MENU:")
		fmt.Println("1. Insert Quarter")
		fmt.Println("2. Eject Quarter")
		fmt.Println("3. Turn Crank")
		fmt.Println("4. Refill (add 5 balls)")
		fmt.Println("5. Exit")
		fmt.Print("Select option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			gm.InsertQuarter()
		case "2":
			gm.EjectQuarter()
		case "3":
			gm.TurnCrank()
		case "4":
			gm.Refill(5)
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
