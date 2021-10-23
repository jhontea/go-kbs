package main

import (
	"fmt"
	"sort"
)

type Config struct {
	Vendor string
	Price  int
	Active bool
}

var configs = []Config{
	{Vendor: "Vendor4", Price: 100, Active: false},
	{Vendor: "Vendor3", Price: 300, Active: true},
	{Vendor: "Vendor2", Price: 200, Active: true},
	{Vendor: "Vendor1", Price: 100, Active: true},
}

func BuildConfig() {
	fmt.Println("configs before", configs)
	SortConfig(configs)
	fmt.Println("configs after", configs)
}

func SortConfig(configs []Config) {

	sort.Slice(configs, func(i, j int) bool {
		var sortedByActive, sortedByLowerPrice bool

		// sort by sold quantity
		sortedByActive = configs[i].Active && !configs[j].Active

		// sort by lowest sold price
		if configs[i].Active == configs[j].Active {
			sortedByLowerPrice = configs[i].Price < configs[j].Price
			return sortedByLowerPrice
		}
		return sortedByActive
	})
}
