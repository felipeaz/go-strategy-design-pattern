package main

import (
	"fmt"
	"go-strategy-design-pattern/character"
	"go-strategy-design-pattern/class"
	"go-strategy-design-pattern/weapon"
)

func main() {
	mage := character.NewCharacter(
		class.NewMage(),
		weapon.NewIceStaff(),
	)
	fmt.Println("----------- Mage -----------")
	mage.LightAttack()
	mage.HeavyAttack()
	mage.PowerAttack()
	mage.SpecialAttack()

	hunter := character.NewCharacter(
		class.NewHunter(),
		weapon.NewBow(),
	)
	fmt.Println("----------- Hunter -----------")
	hunter.LightAttack()
	hunter.HeavyAttack()
	hunter.PowerAttack()
	hunter.SpecialAttack()

	warrior := character.NewCharacter(
		class.NewWarrior(),
		weapon.NewSword(),
	)
	fmt.Println("----------- Warrior -----------")
	warrior.LightAttack()
	warrior.HeavyAttack()
	warrior.PowerAttack()
	warrior.SpecialAttack()

	assassin := character.NewCharacter(
		class.NewAssassin(),
		weapon.NewDaggers(),
	)
	fmt.Println("----------- Assassin -----------")
	assassin.LightAttack()
	assassin.HeavyAttack()
	assassin.PowerAttack()
	assassin.SpecialAttack()
}
