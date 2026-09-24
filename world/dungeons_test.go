package world

import (
	"reflect"
	"testing"
)

func TestFirstDungeonUsesGoblinSkeletonZombie(t *testing.T) {
	city, ok := DungeonsByCity["Ville 1"]
	if !ok {
		t.Fatalf("la ville 1 doit avoir un donjon associé")
	}

	if city.Name != "Foret" {
		t.Fatalf("donjon attendu = %q, obtenu = %q", "Foret", city.Name)
	}

	want := []string{"gobelin", "squelette", "zombie"}
	if !reflect.DeepEqual(city.Monsters, want) {
		t.Fatalf("monstres attendus = %v, obtenus = %v", want, city.Monsters)
	}
}

func TestEachCityHasADungeon(t *testing.T) {
	if len(DungeonsByCity) != 4 {
		t.Fatalf("4 villes attendues, %d présentes", len(DungeonsByCity))
	}
}
