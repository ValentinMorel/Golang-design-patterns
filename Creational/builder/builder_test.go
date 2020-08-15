package main

import (
	"fmt"
	"testing"
)

func TestBuilderPattern(t *testing.T) {

	speed := 100
	color := "yellow"
	testBuilder := New()

	testCar := testBuilder.TopSpeed(speed).Paint(color).Build()
	expected_speed := "Driving at speed : 100"

	if expected_speed != testCar.Drive() {
		fmt.Println("Fail at drive function from car builder.")
	}

}
