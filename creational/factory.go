// https://golangbyexample.com/golang-factory-design-pattern/
// Factory method is a creational design pattern which solves the problem of creating product objects without specifying their concrete classes.
// This pattern provides a way to hide the creation logic of the instances being created.
// The client only interacts with a factory struct and tells the kind of instances that needs to be created.
// The factory class interacts with the corresponding concrete structs and returns the correct instance back.
package creational

import "fmt"

type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getPower() int {
	return g.power
}

type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

type maverick struct {
	gun
}

func newMaverick() iGun {
	return &maverick{
		gun: gun{
			name:  "Maverick gun",
			power: 5,
		},
	}
}

func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "maverick" {
		return newMaverick(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}
