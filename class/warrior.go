package class

import (
	"go-strategy-design-pattern/weapon"
	"log"
)

type warrior struct {
	weapon weapon.Weapon
}

func NewWarrior() Class {
	return warrior{}
}

func (w warrior) PowerAttack() {
	log.Println("<< Warrior hit hard and stuns the enemy >>")
}

func (w warrior) SpecialAttack() {
	log.Println("<< Warrior spins his weapon around the enemy >>")
}
