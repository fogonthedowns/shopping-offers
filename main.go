package main

import (
	"fmt"
)

func shoppingOffers(price []int, special [][]int, needs []int) int {
	return shopping(price, special, needs)
}

func shopping(price []int, specials [][]int, needs []int) int {

	res := noDiscountPrice(needs, price)

	// s []int a,b,price
mainLoop:
	for _, s := range specials {
		needsClone := make([]int, len(needs))
		copy(needsClone, needs)

		for i := range needsClone {
			// is special possible?
			diff := needsClone[i] - s[i]

			if diff < 0 {
				continue mainLoop
			}
			needsClone[i] = diff

		}
		r := shopping(price, specials, needsClone)
		fmt.Println(r)
		fmt.Println(s[len(price)])
		fmt.Println("***")
		res = min(res, s[len(price)]+r)
	}
	return res
}

// calculates price without discounts applied
func noDiscountPrice(needs []int, price []int) int {
	var sum int
	for i := 0; i < len(price); i++ {
		sum += needs[i] * price[i]
	}
	return sum
}

// min func
func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}
