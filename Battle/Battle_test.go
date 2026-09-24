package ProjetRED

import (
	"testing"

	personnage "ProjetRED/Personnage"
	enemies "ProjetRED/enemies"
)

func TestRunDemoBattleDropsLootAfterVictory(t *testing.T) {
	player := personnage.CharacterCreation("test", personnage.Classes["Ronin"])
	player.Strength = 10
	player.Spd = 100

	monster := &enemies.MONSTER{
		NOM:    "monstre de test",
		PVMax:  1,
		PV:     1,
		PVMAXR: 1,
		PVR:    1,
		Spd:    1,
		Loot: []enemies.LootEntry{
			{Name: "Fer", Rate: 100},
		},
	}

	RunDemoBattle(&player, monster, []int{1})

	if player.Inventory.Materials["Fer"] != 1 {
		t.Fatalf("expected 1 Fer after victory, got %d", player.Inventory.Materials["Fer"])
	}
}
