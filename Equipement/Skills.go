package ProjetRED

import (
	personnage "ProjetRED/Personnage"
)

type Skill struct {
	Name       string
	BaseDamage int
	Heal       int
	Cooldown   int
	Type       string // "Nature" ou "Magic"
}

var SkillList = map[string]Skill{
	"Slash": {
		Name:       "Slash",
		BaseDamage: 15,
		Cooldown:   1,
		Type:       "Nature",
	},
	"Punch": {
		Name:       "Punch",
		BaseDamage: 5,
		Type:       "Nature",
	},
	"Dragon Slayer": {
		Name:       "Dragon Slayer",
		BaseDamage: 50,
		Type:       "Nature",
	},
	"Charge": {
		Name:       "Charge",
		BaseDamage: 10,
		Type:       "Nature",
	},
	"Fireball": {
		Name:       "Fireball",
		BaseDamage: 30,
		Cooldown:   3,
		Type:       "Magic",
	},
	"Power Strike": {
		Name:       "Power Strike",
		BaseDamage: 40,
		Cooldown:   3,
		Type:       "Magic",
	},
	"Healing": {
		Name:     "Healing",
		Heal:     40,
		Cooldown: 3,
		Type:     "Magic",
	},
}

func CalculateSkillDamage(p personnage.Character, skill Skill) int {
	dmg := skill.BaseDamage

	switch skill.Type {
	case "Nature":
		dmg += p.Strength
	case "Magic":
		dmg += p.Reiki
	}

	return dmg
}
