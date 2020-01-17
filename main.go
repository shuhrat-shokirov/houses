package main

import (
	"sort"
)

type tel struct {
	countryCode string
	number      string
}

type houses struct {
	id   int64
	name string
	//postedBy int64
	price                 int64
	currency              string
	distanceFromTheCenter int64
	distanceValue         string
	district              string
	houseArea             int64
	houseAreaUnit         string
	landArea              int64
	landAreaUnit          string
	flatRoomCount         int64
	storeys               int64
	phone                 tel
	//text string
}

func sortBy(house []houses, less func(a, b houses) bool) []houses {
	result := make([]houses, len(house))
	copy(result, house)
	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})
	return result
}

func sortByPriceAsc(house []houses) []houses {
	return sortBy(house, func(a, b houses) bool {
		return a.price < b.price
	})
}

func sortByPriceDec(house []houses) []houses {
	return sortBy(house, func(a, b houses) bool {
		return a.price > b.price
	})
}
func sortByDistanceFromTheCenterAsc(house []houses) []houses {
	return sortBy(house, func(a, b houses) bool {
		return a.distanceFromTheCenter < b.distanceFromTheCenter
	})
}

func sortByDistanceFromTheCenterDec(house []houses) []houses {
	return sortBy(house, func(a, b houses) bool {
		return a.distanceFromTheCenter > b.distanceFromTheCenter
	})
}

func filterBy(house []houses, predicate func(house houses) bool) []houses {
	result := make([]houses, 0)
	for _, house := range house {
		if predicate(house) {
			result = append(result, house)
		}
	}
	return result
}

func searchByPrice(house []houses, limit int64) []houses {
	return filterBy(house, func(house houses) bool {
		return house.price <= limit
	})
}

func searchWithinPrice(house []houses, lowerLimit, upperLimit int64) []houses {
	return filterBy(house, func(house houses) bool {
		if house.price < lowerLimit {
			return false
		}
		if house.price > upperLimit {
			return false
		}
		return true
	})
}

func findByDistrict(house []houses, district string) []houses {
	return filterBy(house, func(house houses) bool {
		return house.district == district
	})
}

func findByDistricts(homes []houses, districts []string) []houses {
	result := make([]houses, 0)
	for _, district := range districts  {
		result =append(result, findByDistrict(homes, district)...)
	}
	return result
}

func main() {
}
