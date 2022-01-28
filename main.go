package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)

	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add Item, s - Save Bill, t - add Tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item Name: ", reader)
		price, _ := getInput("Item Price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Item Added: ", name, price)
		promptOptions(b)

	case "t":
		tip, _ := getInput("Enter tip amount: ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip Added: ", tip)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("Bill saved ", b)
	default:
		fmt.Println("wrong option")
		promptOptions(b)
	}

	//fmt.Println(opt)
}

func main() {

	//var bill1 bill = newBill("bill1")
	// mybill := newBill("mario's bill")
	// mybill.updateTip(20)
	// mybill.addItem("Pie", 120)
	// mybill.addItem("Pizza", 450)
	// mybill.addItem("Paneer Role", 150)

	// fmt.Println(mybill.format())

	myBill1 := createBill()
	promptOptions(myBill1)
}
