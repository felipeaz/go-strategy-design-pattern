package weapon

import "log"

type daggers struct{}

func NewDaggers() Weapon {
	return daggers{}
}

func (d daggers) LightAttack() {
	log.Println("hits the target four times in a row")
}

func (d daggers) HeavyAttack() {
	log.Println("deep strike on the target")
}
