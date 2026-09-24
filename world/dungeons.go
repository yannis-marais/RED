package world

import "math/rand"
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
		Monsters: []string{"orc", "troll", "loup-garou", "vouivre", "zombie", "loup-garou"},
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

func Aleatoire(ville string) string {
    mons := GetMonsterNamesForCity(ville)
    if len(mons) == 0 {
        return ""
    }
    return mons[rand.Intn(len(mons))]
}