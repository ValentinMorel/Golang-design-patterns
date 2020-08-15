package main

import (
	"fmt"
	"strconv"
)

// Builder pattern aims to :
// abstract complex creations -> object creation is separated from the object user
// create an object step by step
// reuse the object creation algo between many objects

const (
	BLUE = "blue"
	RED  = "red"
)

type car struct {
	topSpeed int
	color    string
}

type Car interface {
	Drive() string
	Stop() string
}

type carBuilder struct {
	speedOption int
	color       string
}

type CarBuilder interface {
	TopSpeed(int) CarBuilder
	Paint(string) CarBuilder
	Build() Car
}

func New() CarBuilder {
	return &carBuilder{}
}

func (cb *carBuilder) TopSpeed(speed int) CarBuilder {
	cb.speedOption = speed
	return cb
}

func (cb *carBuilder) Paint(color string) CarBuilder {
	cb.color = color
	return cb
}

func (cb *carBuilder) Build() Car {
	return &car{
		topSpeed: cb.speedOption,
		color:    cb.color,
	}
}

func (c *car) Drive() string {
	return "Driving at speed : " + strconv.Itoa(c.topSpeed)
}

func (c *car) Stop() string {
	return "Stopping a " + c.color + " car"
}

func main() {
	builder := New()
	car := builder.TopSpeed(50).Paint(BLUE).Build()

	fmt.Println(car.Drive())
	fmt.Println(car.Stop())
}
