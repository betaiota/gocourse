package tasks

import (
	"fmt"
	"math/rand"
)

type Character struct {
	Name       string
	HP         int
	Level      int
	Experience int
	Class      string
}

type Fighter interface {
	AttackSpecial() bool
}

func (c *Character) AttackHands(enemy *Character) {
	fmt.Printf("%v %v is attacking %v %v with his bare hands! Dealing %v damage.\n", c.Class, c.Name, enemy.Class, enemy.Name, 5)
	enemy.HP = enemy.HP - 5
	fmt.Printf("%v %v HP is now %v\n", enemy.Class, enemy.Name, enemy.HP)
}

type Warrior struct {
	Rage int
	Character
}

func (w *Warrior) AttackSpecial(enemy *Character) bool {
	if w.Rage > 0 {
		damage := (rand.Intn(4)) * 4
		fmt.Printf("%v %v is attacking %v %v with his sword! Dealing %v damage.\n", w.Class, w.Name, enemy.Class, enemy.Name, damage)
		enemy.HP = enemy.HP - damage
		fmt.Printf("%v %v HP is now %v\n", enemy.Class, enemy.Name, enemy.HP)
		w.Rage = w.Rage - 10
	} else {
		w.Character.AttackHands(enemy)
	}
	if enemy.HP <= 0 {
		expModifier := rand.Intn(6) + 1
		fmt.Printf("%v %v has won! Gained %v exp!\n", w.Class, w.Name, enemy.Level*expModifier)
		w.Experience = w.Experience + enemy.Level*expModifier
		return true
	}
	return false
}

func (m *Mage) AttackSpecial(enemy *Character) bool {
	if m.Mana > 0 {
		damage := (rand.Intn(6) + 1) * 2
		fmt.Printf("%v %v is throwing fireball to %v %v! Dealing %v damage.\n", m.Class, m.Name, enemy.Class, enemy.Name, damage)
		enemy.HP = enemy.HP - damage
		fmt.Printf("%v %v HP is now %v\n", enemy.Class, enemy.Name, enemy.HP)
	} else {
		m.Character.AttackHands(enemy)
	}
	if enemy.HP <= 0 {
		expModifier := rand.Intn(12) + 1
		fmt.Printf("%v %v has won! Gained %v exp!\n", m.Class, m.Name, enemy.Level*expModifier)
		m.Experience = m.Experience + enemy.Level*expModifier
		return true
	}
	return false
}

type Mage struct {
	Mana int
	Character
}

type Archer struct {
	Arrows int
	Character
}

func CheckPossibleLevelup(chars ...*Character) {
	for _, ch := range chars {
		if ch.Experience >= 25 {
			ch.Experience = 0
			ch.Level = ch.Level + 1
			fmt.Printf("%v %v leveled up, congrats! He now has level of %v\n", ch.Class, ch.Name, ch.Level)
		}
	}
}

func LabSixteen() {
	mage := &Mage{
		Character: Character{
			Name:       "Gandalf",
			HP:         50,
			Level:      3,
			Experience: 0,
			Class:      "Mage",
		},
		Mana: 30,
	}

	warrior := &Warrior{
		Character: Character{
			Name:       "Sauron the Deciever",
			HP:         60,
			Level:      3,
			Experience: 0,
			Class:      "Warrior",
		},
		Rage: 30,
	}

	for {
		if mageRes := mage.AttackSpecial(&warrior.Character); mageRes == true {
			break
		}
		if warRes := warrior.AttackSpecial(&mage.Character); warRes == true {
			break
		}
	}

	CheckPossibleLevelup(&mage.Character, &warrior.Character)

}
