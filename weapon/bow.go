package weapon

import "log"

type bow struct{}

func NewBow() Weapon {
	return bow{}
}

func (b bow) LightAttack() {
	log.Println("shoot an arrow")
}

func (b bow) HeavyAttack() {
	log.Println("shoot three arrows in a row")
}
