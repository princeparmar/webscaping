package main

import (
	"fmt"

	"github.com/princeparmar/webscaping/parser"
)

func main() {
	collections := []string{
		"bridal-lehengas",
		"salwar-kameez",
		"desinger-saree",
	}

	for _, str := range collections {
		products, err := parser.FullCollection(str)
		if err != nil {
			fmt.Println("collection is not completed ", err)
			return
		}

		err = parser.SaveAllProducts(products...)
		if err != nil {
			fmt.Println("collection is not saved in file ", err)
		}
	}
}
