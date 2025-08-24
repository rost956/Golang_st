package main

import (
	"fmt"
	"structures_import/magazine"
)

func main() {
	address := magazine.Address{Street: "123 Oak st",
		City: "Omaha", State: "NE", PostalCode: "68111"}
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.Street = address.Street //анонимные поля
	//...
	fmt.Println(subscriber)
}
