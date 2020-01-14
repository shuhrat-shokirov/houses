package main

import "fmt"

var house = []houses{
	{
		id:   1,
		name: "Дом №1",
		//postedBy:				1578841080,
		price:                 560_000,
		currency:              "Сомони",
		distanceFromTheCenter: 26_000_000,
		distanceValue:         "м",
		district:              "Варзоб",
		houseArea:             133,
		houseAreaUnit:         "м^2",
		landArea:              12,
		landAreaUnit:          "соток",
		flatRoomCount:         6,
		storeys:               2,
		phone:                 tel{countryCode: "+992", number: "93 58xxxxx"},
		//text:					"...."
	},
	{
		id:   2,
		name: "Дом №2",
		//postedBy:				1579030380,
		price:                 2_500_000,
		currency:              "Сомони",
		distanceFromTheCenter: 10_000_000,
		distanceValue:         "м",
		district:              "Варзоб",
		landArea:              11,
		landAreaUnit:          "соток",
		flatRoomCount:         7,
		storeys:               3,
		phone:                 tel{countryCode: "+992", number: "91 10xxxxx"},
		//text:					"...."
	},
	{
		id:   3,
		name: "Дом №3",
		//postedBy:				1578784200,
		price:                 2_700_000,
		currency:              "Сомони",
		distanceFromTheCenter: 5_000_000,
		distanceValue:         "м",
		houseArea:             540,
		houseAreaUnit:         "м^2",
		district:              "Рудаки",
		landArea:              30,
		landAreaUnit:          "соток",
		flatRoomCount:         6,
		storeys:               3,
		phone:                 tel{countryCode: "+992", number: "90 22xxxxx"},
		//text:					"...."
	},
	{
		id:   4,
		name: "Дом №4",
		//postedBy:				1578916620,
		price:                 497_000,
		currency:              "Сомони",
		distanceFromTheCenter: 5_000_000,
		distanceValue:         "м",
		district:              "Зарафшон",
		landArea:              3,
		landAreaUnit:          "соток",
		flatRoomCount:         3,
		storeys:               1,
		phone:                 tel{countryCode: "+992", number: "98 84xxxxx"},
		//text:					"...."
	},
}

func ExampleSortByPriceAsc() {
	ascSorted := sortByPriceAsc(house)
	fmt.Println(ascSorted)
	// Output: [{4 Дом №4 497000 Сомони 5000000 м Зарафшон 0  3 соток 3 1 {+992 98 84xxxxx}} {1 Дом №1 560000 Сомони 26000000 м Варзоб 133 м^2 12 соток 6 2 {+992 93 58xxxxx}} {2 Дом №2 2500000 Сомони 10000000 м Варзоб 0  11 соток 7 3 {+992 91 10xxxxx}} {3 Дом №3 2700000 Сомони 5000000 м Рудаки 540 м^2 30 соток 6 3 {+992 90 22xxxxx}}]
}

func ExampleSortByPriceDec() {
	decSorted := sortByPriceDec(house)
	fmt.Println(decSorted)
	// Output:[{3 Дом №3 2700000 Сомони 5000000 м Рудаки 540 м^2 30 соток 6 3 {+992 90 22xxxxx}} {2 Дом №2 2500000 Сомони 10000000 м Варзоб 0  11 соток 7 3 {+992 91 10xxxxx}} {1 Дом №1 560000 Сомони 26000000 м Варзоб 133 м^2 12 соток 6 2 {+992 93 58xxxxx}} {4 Дом №4 497000 Сомони 5000000 м Зарафшон 0  3 соток 3 1 {+992 98 84xxxxx}}]
}

func ExampleSortByDistanceFromTheCenterAsc() {
	ascSorted := sortByDistanceFromTheCenterAsc(house)
	fmt.Println(ascSorted)
	// Output: [{3 Дом №3 2700000 Сомони 5000000 м Рудаки 540 м^2 30 соток 6 3 {+992 90 22xxxxx}} {4 Дом №4 497000 Сомони 5000000 м Зарафшон 0  3 соток 3 1 {+992 98 84xxxxx}} {2 Дом №2 2500000 Сомони 10000000 м Варзоб 0  11 соток 7 3 {+992 91 10xxxxx}} {1 Дом №1 560000 Сомони 26000000 м Варзоб 133 м^2 12 соток 6 2 {+992 93 58xxxxx}}]
}

func ExampleSortByDistanceFromTheCenterDec() {
	decSorted := sortByDistanceFromTheCenterDec(house)
	fmt.Println(decSorted)
	// Output: [{1 Дом №1 560000 Сомони 26000000 м Варзоб 133 м^2 12 соток 6 2 {+992 93 58xxxxx}} {2 Дом №2 2500000 Сомони 10000000 м Варзоб 0  11 соток 7 3 {+992 91 10xxxxx}} {3 Дом №3 2700000 Сомони 5000000 м Рудаки 540 м^2 30 соток 6 3 {+992 90 22xxxxx}} {4 Дом №4 497000 Сомони 5000000 м Зарафшон 0  3 соток 3 1 {+992 98 84xxxxx}}]
}

func ExampleSearchByPrice_NoResults() {
	result := searchByPrice(house, 400_000)
	fmt.Println(result)
	// Output: []
}

func ExampleSearchByPrice_OneResult() {
	result := searchByPrice(house, 500_000)
	fmt.Println(result)
	// Output: [{4 Дом №4 497000 Сомони 5000000 м Зарафшон 0  3 соток 3 1 {+992 98 84xxxxx}}]
}

func ExampleSearchByPrice_SeveralResults() {
	result := searchByPrice(house, 2_500_000)
	fmt.Println(result)
	// Output: [{1 Дом №1 560000 Сомони 26000000 м Варзоб 133 м^2 12 соток 6 2 {+992 93 58xxxxx}} {2 Дом №2 2500000 Сомони 10000000 м Варзоб 0  11 соток 7 3 {+992 91 10xxxxx}} {4 Дом №4 497000 Сомони 5000000 м Зарафшон 0  3 соток 3 1 {+992 98 84xxxxx}}]
}

func ExampleSearchWithinPrice_NoResults() {
	result := searchWithinPrice(house, 100_000, 400_000)
	fmt.Println(result)
	// Output: []
}

func ExampleSearchWithinPrice_OneResults() {
	result := searchWithinPrice(house, 400_000, 500_000)
	fmt.Println(result)
	// Output: [{4 Дом №4 497000 Сомони 5000000 м Зарафшон 0  3 соток 3 1 {+992 98 84xxxxx}}]
}

func ExampleSearchWithinPrice_SeveralResults() {
	result := searchWithinPrice(house, 500_000, 2_500_000)
	fmt.Println(result)
	// Output: [{1 Дом №1 560000 Сомони 26000000 м Варзоб 133 м^2 12 соток 6 2 {+992 93 58xxxxx}} {2 Дом №2 2500000 Сомони 10000000 м Варзоб 0  11 соток 7 3 {+992 91 10xxxxx}}]
}

func ExampleFindByDistrict_NoResults() {
	result := findByDistrict(house, "Душанбе")
	fmt.Println(result)
	// Output: []
}

func ExampleFindByDistrict_OneResults() {
	result := findByDistrict(house, "Зарафшон")
	fmt.Println(result)
	// Output: [{4 Дом №4 497000 Сомони 5000000 м Зарафшон 0  3 соток 3 1 {+992 98 84xxxxx}}]
}

func ExampleFindByDistrict_SeveralResults() {
	result := findByDistrict(house, "Варзоб")
	fmt.Println(result)
	// Output: [{1 Дом №1 560000 Сомони 26000000 м Варзоб 133 м^2 12 соток 6 2 {+992 93 58xxxxx}} {2 Дом №2 2500000 Сомони 10000000 м Варзоб 0  11 соток 7 3 {+992 91 10xxxxx}}]
}

func ExampleFindByDistricts_NoResults() {
	result := findByDistricts(house, []string{"Сомони"})
	fmt.Println(result)
	// Output: []
}

func ExampleFindByDistricts_OneResults() {
	result := findByDistricts(house, []string{"Сомони", "Душанбе", "Зарафшон"})
	fmt.Println(result)
	// Output: [{4 Дом №4 497000 Сомони 5000000 м Зарафшон 0  3 соток 3 1 {+992 98 84xxxxx}}]
}

func ExampleFindByDistricts_SeveralResults() {
	result := findByDistricts(house, []string{"Сомони", "Душанбе", "Зарафшон", "Рудаки", "Варзоб"})
	fmt.Println(result)
	// Output: [{4 Дом №4 497000 Сомони 5000000 м Зарафшон 0  3 соток 3 1 {+992 98 84xxxxx}} {3 Дом №3 2700000 Сомони 5000000 м Рудаки 540 м^2 30 соток 6 3 {+992 90 22xxxxx}} {1 Дом №1 560000 Сомони 26000000 м Варзоб 133 м^2 12 соток 6 2 {+992 93 58xxxxx}} {2 Дом №2 2500000 Сомони 10000000 м Варзоб 0  11 соток 7 3 {+992 91 10xxxxx}}]
}
