package class

import (
	"go-strategy-design-pattern/weapon"
	"log"
)

type hunter struct {
	weapon weapon.Weapon
}

func NewHunter() Class {
	return hunter{}
}

func (h hunter) PowerAttack() {
	log.Println("<< Hunter called two wolves that rushed into the enemy >>")
}

func (h hunter) SpecialAttack() {
	log.Println("<< Hunter shoot a charged arrow >>")
}
