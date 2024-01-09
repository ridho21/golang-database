package main

import (
	"fmt"

	"github.com/ridho21/geometry-lib/shape"
)

func main() {
	field1 := shape.Rectangle{Width: 15.0, Length: 5.0}
	field2 := shape.Rectangle{Width: 17.5, Length: 15.5}

	harvestField1 := field1.Area() / 100
	harvestField2 := field2.Area() / 100

	fmt.Println(harvestField1)
	fmt.Println(harvestField2)
}
