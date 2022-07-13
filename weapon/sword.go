package weapon

import "log"

type sword struct{}

func NewSword() Weapon {
	return sword{}
}

func (s sword) LightAttack() {
	log.Println("hits the target from the right")
}

func (s sword) HeavyAttack() {
	log.Println("jumps and hit the target from above")
}
