package main

import (
	"fmt"
	"os"
)

type bill struct { // type - to define the type // bill - name that  // struct - define it's a structure

	// define different properties of this structure
	name  string
	items map[string]float64
	tip   float64
}

// func to make new bill

func newBill(name string) bill {
	//  create bill
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// reciever func to format the bill
func (b bill) format() string {
	fs := "Bill Breakdown: \n"
	var total float64 = 0

	//list items

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}
	// add tip
	fs += fmt.Sprintf("%-25v ...$%v \n", "tip:", b.tip)

	//total
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "total:", total+b.tip)

	return fs
}

// update the tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// Add item to the bill

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save() {
	data := []byte(b.format()) // slice of bytes
	err := os.WriteFile("bills/"+b.name+".txt", data, 0664)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill saved to the file")
}
