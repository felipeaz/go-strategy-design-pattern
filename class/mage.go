package class

import (
	"go-strategy-design-pattern/weapon"
	"log"
)

type mage struct {
	weapon weapon.Weapon
}

func NewMage() Class {
	return mage{}
}

func (m mage) PowerAttack() {
	log.Println("<< Mage cast a fireball at the enemy >>")
}

func (m mage) SpecialAttack() {
	log.Println("<< Mage called a meteor rain on the enemy area >>")
}
