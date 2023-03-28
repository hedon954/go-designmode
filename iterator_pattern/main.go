package main

import (
	"fmt"
)

func main() {
	productList := ProductList{
		Products: []*Product{
			{Name: "iPhone", Price: 9.9},
			{Name: "iPad", Price: 8.8},
			{Name: "Mac", Price: 6.6},
		},
	}
	iterator := productList.CreateIterator()
	for iterator.HasNext() {
		product := iterator.Next().(*Product)
		fmt.Println(product)
	}
}
