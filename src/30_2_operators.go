package main

import (
	"fmt"
)

func main() {
	var mealCost float32
	var tipPercent int32
	var taxPercent int32

	fmt.Scanf("%f\n%d\n%d\n", &mealCost, &tipPercent, &taxPercent)
	var tip float32 = mealCost * float32(tipPercent) / 100.0
	var tax float32 = mealCost * float32(taxPercent) / 100.0
	var totalCost int32 = int32(mealCost + tip + tax + 0.5)
	fmt.Printf("The total meal cost is %d dollars.", totalCost)
}
