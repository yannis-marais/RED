package world

import (
	"math/rand"

	enemies "ProjetRED/enemies"
)

// DungeonDefinition représente un donjon associé à une ville.
type DungeonDefinition struct {
	Name     string
	Monsters []string
}

// DungeonsByCity contient, pour chaque ville, le donjon associé et les monstres qui y vivent.
// Le premier donjon est volontairement limité à gobelin, squelette et zombie comme demandé.
var DungeonsByCity = map[string]DungeonDefinition{
	"Ville 1": {
		Name:     "Foret",
		Monsters: []string{"gobelin", "squelette", "zombie"},
	},
	"Ville 2": {
		Name:     "Grotte",
		Monsters: []string{"orc", "troll", "loup-garou"},
	},
	"Ville 3": {
		Name:     "Manoir",
		Monsters: []string{"vouivre", "zombie", "loup-garou"},
	},
	"Ville 4": {
		Name:     "Plaine",
		Monsters: []string{"dragon", "orc", "troll"},
	},
}

func GetDungeonForCity(cityName string) (DungeonDefinition, bool) {
	dungeon, ok := DungeonsByCity[cityName]
	return dungeon, ok
}

func GetMonsterNamesForCity(cityName string) []string {
	dungeon, ok := GetDungeonForCity(cityName)
	if !ok {
		return nil
	}
	return append([]string(nil), dungeon.Monsters...)
}

func RandomMonsterNameForCity(cityName string) (string, bool) {
	names := GetMonsterNamesForCity(cityName)
	if len(names) == 0 {
		return "", false
	}
	return names[rand.Intn(len(names))], true
}

func RandomMonsterForCity(cityName string) (*enemies.MONSTER, bool) {
	monsterName, ok := RandomMonsterNameForCity(cityName)
	if !ok {
		return nil, false
	}
	monster, ok := enemies.NewMonsterByName(monsterName)
	if !ok {
		return nil, false
	}
	return monster, true
}
