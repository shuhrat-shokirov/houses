package main

import (
	//"nettextproto"
	//"sort"
	//"syscall"
	"sort"
)

type tel struct {
	countryCode string
	number string
}

type houses struct {
	id int64
	name string
	//postedBy int64
	price int64
	currency string
	distanceFromTheCenter int64
	distanceValue string
	district string
	houseArea int64
	houseAreaUnit string
	landArea int64
	landAreaUnit string
	flatRoomCount int64
	storeys int64
	phone tel
	//text string
}

func sortByPriceAsc(house []houses) []houses {
	result := make([]houses, len(house))
	copy(result, house)
	sort.Slice(result, func(i, j int) bool {
		return result[i].price < result[j].price
	})
	return result
}

func sortByPriceDec(house []houses) []houses {
	result := make([]houses, len(house))
	copy(result, house)
	sort.Slice(result, func(i, j int) bool {
		return result[i].price > result[j].price
	})
	return result
}
func sortByDistanceFromTheCenterAsc(house []houses) []houses {
	result := make([]houses, len(house))
	copy(result, house)
	sort.Slice(result, func(i, j int) bool {
		return result[i].distanceFromTheCenter < result[j].distanceFromTheCenter
	})
	return result
}

func sortByDistanceFromTheCenterDec(house []houses) []houses {
	result := make([]houses, len(house))
	copy(result, house)
	sort.Slice(result, func(i, j int) bool {
		return result[i].distanceFromTheCenter > result[j].distanceFromTheCenter
	})
	return result
}

func searchByPrice(house []houses, limit int64) []houses {
	result := make([]houses, 0)
	for _, house :=range house {
		if house.price <= limit {
			result = append(result, house)
		}
	}
	return result
}

func searchWithinPrice(house []houses, lowerLimit, upperLimit int64) []houses {
	result := make([]houses, 0)
	for _, house :=range house {
		if lowerLimit <= house.price && house.price <= upperLimit {
			result = append(result, house)
		}
	}
	return result
}

func findByDistrict(house []houses, district string) []houses {
	result := make([]houses, 0)
	for _, house := range house {
		if house.district == district{
			result = append(result, house)
		}
	}
	return result
}

func findByDistricts(homes []houses, districts []string) []houses {
	result := make([]houses, 0)
	for _, district := range districts {
		for _, home := range homes {
			if home.district == district{
				result = append(result, home)
			}
		}
	}
	return result
}

func main() {
	
}
