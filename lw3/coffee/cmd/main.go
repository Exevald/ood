package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"coffee/pkg/model"
	"coffee/pkg/model/beverages"
	"coffee/pkg/model/condiments"
)

func main() {
	DialogWithUser()

	fmt.Println()

	{
		beverage := condiments.NewChocolateCrumbs(
			condiments.NewIceCube(
				condiments.NewLemon(
					condiments.NewCinnamon(beverages.NewLatte(beverages.LatteSizeDefault)),
					2,
				),
				2, condiments.IceCubeTypeDry,
			),
			2,
		)
		fmt.Printf("%s costs %.2f\n", beverage.GetDescription(), beverage.GetCost())
	}
	{

		beverage := beverages.NewLatte(beverages.LatteSizeDefault)
		beverage = condiments.NewCinnamon(beverage)
		beverage = condiments.NewLemon(beverage, 2)
		beverage = condiments.NewIceCube(beverage, 2, condiments.IceCubeTypeDry)
		beverage = condiments.NewChocolateCrumbs(beverage, 2)
		fmt.Printf("%s costs %.2f\n", beverage.GetDescription(), beverage.GetCost())

		beverage = beverages.NewMilkshake(beverages.MilkshakeSizeMedium)
		beverage = condiments.NewSyrup(beverage, condiments.SyrupTypeMaple)
		beverage = condiments.NewCoconutFlakes(beverage, 8)
		fmt.Printf("%s costs %.2f\n", beverage.GetDescription(), beverage.GetCost())
	}
	{

		beverage := beverages.NewTea(beverages.TeaTypeTieguanyin)
		beverage = condiments.NewCream(beverage)

		beverage = condiments.NewLiqueur(beverage, condiments.LiqueurTypeNut)
		fmt.Printf("%s costs %.2f\n", beverage.GetDescription(), beverage.GetCost())
	}
}

func DialogWithUser() {
	fmt.Println("Choose beverage:")
	fmt.Println("1 - Coffee")
	fmt.Println("2 - Latte (Standard)")
	fmt.Println("3 - Double Latte")
	fmt.Println("4 - Cappuccino (Standard)")
	fmt.Println("5 - Double Cappuccino")
	fmt.Println("6 - Tea")
	fmt.Println("7 - Milkshake")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	choice, _ := strconv.Atoi(strings.TrimSpace(input))

	var beverage model.Beverage

	switch choice {
	case 1:
		beverage = beverages.NewCoffee()
	case 2:
		beverage = beverages.NewLatte(beverages.LatteSizeDefault)
	case 3:
		beverage = beverages.NewLatte(beverages.LatteSizeDouble)
	case 4:
		beverage = beverages.NewCappuccino(beverages.CappuccinoSizeDefault)
	case 5:
		beverage = beverages.NewCappuccino(beverages.CappuccinoSizeDouble)
	case 6:
		beverage = chooseTea(reader)
	case 7:
		beverage = chooseMilkshake(reader)
	default:
		fmt.Println("Invalid choice")
		return
	}

	for {
		fmt.Println("\nAdd condiment:")
		fmt.Println("1 - Lemon (2 slices)")
		fmt.Println("2 - Cinnamon")
		fmt.Println("3 - Cream")
		fmt.Println("4 - Chocolate (3 slices)")
		fmt.Println("5 - Nut Liqueur")
		fmt.Println("0 - Checkout")
		input, _ = reader.ReadString('\n')
		condChoice, _ := strconv.Atoi(strings.TrimSpace(input))

		switch condChoice {
		case 1:
			beverage = condiments.NewLemon(beverage, 2)
		case 2:
			beverage = condiments.NewCinnamon(beverage)
		case 3:
			beverage = condiments.NewCream(beverage)
		case 4:
			beverage = condiments.NewChocolateSlice(beverage, 3)
		case 5:
			beverage = condiments.NewLiqueur(beverage, condiments.LiqueurTypeNut)
		case 0:
			goto checkout
		default:
			fmt.Println("Invalid choice")
			continue
		}
	}

checkout:
	fmt.Printf("\nYour order: %s\nTotal cost: %.2f\n", beverage.GetDescription(), beverage.GetCost())
}

func chooseTea(reader *bufio.Reader) model.Beverage {
	fmt.Println("Choose tea type:")
	fmt.Println("1 - 7542")
	fmt.Println("2 - Good 2.0")
	fmt.Println("3 - Tieguanyin")
	fmt.Println("4 - Da Hong Pao")
	input, _ := reader.ReadString('\n')
	choice, _ := strconv.Atoi(strings.TrimSpace(input))

	switch choice {
	case 1:
		return beverages.NewTea(beverages.TeaType7542)
	case 2:
		return beverages.NewTea(beverages.TeaTypeGood2)
	case 3:
		return beverages.NewTea(beverages.TeaTypeTieguanyin)
	case 4:
		return beverages.NewTea(beverages.TeaTypeDaHongPao)
	default:
		fmt.Println("Defaulting to DaHongPao tea")
		return beverages.NewTea(beverages.TeaTypeDaHongPao)
	}
}

func chooseMilkshake(reader *bufio.Reader) model.Beverage {
	fmt.Println("Choose milkshake size:")
	fmt.Println("1 - Small (50)")
	fmt.Println("2 - Medium (60)")
	fmt.Println("3 - Large (80)")
	input, _ := reader.ReadString('\n')
	choice, _ := strconv.Atoi(strings.TrimSpace(input))

	switch choice {
	case 1:
		return beverages.NewMilkshake(beverages.MilkshakeSizeSmall)
	case 2:
		return beverages.NewMilkshake(beverages.MilkshakeSizeMedium)
	case 3:
		return beverages.NewMilkshake(beverages.MilkshakeSizeLarge)
	default:
		fmt.Println("Defaulting to Medium")
		return beverages.NewMilkshake(beverages.MilkshakeSizeMedium)
	}
}
