package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"

	"simuduck/pkg/canvas/tdewolff"
	"simuduck/pkg/picture"
)

func main() {
	pic := picture.NewPicture()
	canvas := tdewolff.NewCanvas(800, 600, "shapes.png")

	handler := NewCommandHandler(pic, canvas)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		command := scanner.Text()
		if command == "exit" {
			break
		}

		err := handler.HandleCommand(command)
		if err != nil {
			if errors.Is(err, errEmptyCommand) {
				continue
			}
			log.Fatal(err)
		}

		pic.DrawPicture(canvas)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if err := canvas.SaveToFile(""); err != nil {
		log.Fatal(err)
	}
}
