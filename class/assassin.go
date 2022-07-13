package class

import (
	"go-strategy-design-pattern/weapon"
	"log"
)

type assassin struct {
	weapon weapon.Weapon
}

func NewAssassin() Class {
	return assassin{}
}

func (m assassin) PowerAttack() {
	log.Println("<< Assassin became invisible >>")
}

func (m assassin) SpecialAttack() {
	log.Println("<< Assassin summoned clones and attacked the enemy >>")
}
