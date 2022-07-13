package weapon

import "log"

type iceStaff struct{}

func NewIceStaff() Weapon {
	return iceStaff{}
}

func (s iceStaff) LightAttack() {
	log.Println("shoot an ice block")
}

func (s iceStaff) HeavyAttack() {
	log.Println("throw an ice ball that slows and damaged the target")
}
