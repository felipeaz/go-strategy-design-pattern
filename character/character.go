package character

import (
	"go-strategy-design-pattern/class"
	"go-strategy-design-pattern/weapon"
)

type Character struct {
	Class  class.Class
	Weapon weapon.Weapon
}

func NewCharacter(c class.Class, w weapon.Weapon) Character {
	return Character{
		Class:  c,
		Weapon: w,
	}
}

func (c Character) LightAttack() {
	c.Weapon.LightAttack()
}

func (c Character) HeavyAttack() {
	c.Weapon.HeavyAttack()
}

func (c Character) PowerAttack() {
	c.Class.PowerAttack()
}

func (c Character) SpecialAttack() {
	c.Class.SpecialAttack()
}
