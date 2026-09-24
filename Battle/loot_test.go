package ProjetRED

import (
	"testing"

	personnage "ProjetRED/Personnage"
	enemies "ProjetRED/enemies"
)

func TestMonsterVictoryDropsMaterial(t *testing.T) {
	player := personnage.CharacterCreation("test", personnage.Classes["Ronin"])
	monster := &enemies.MONSTER{
		NOM: "monstre de test",
		PV:  0,
		PVR: 0,
		Loot: []enemies.LootEntry{
			{Name: "Fer", Rate: 100},
		},
	}

	dropped := GiveMonsterLoot(&player, monster)

	if len(dropped) != 1 || dropped[0] != "Fer" {
		t.Fatalf("expected Fer to drop, got %#v", dropped)
	}
	if player.Inventory.Materials["Fer"] != 1 {
		t.Fatalf("expected 1 Fer in inventory, got %d", player.Inventory.Materials["Fer"])
	}
}

func TestMonsterVictoryDropsEquipment(t *testing.T) {
	player := personnage.CharacterCreation("test", personnage.Classes["Ronin"])
	monster := &enemies.MONSTER{
		NOM: "monstre de test",
		PV:  0,
		PVR: 0,
		Loot: []enemies.LootEntry{
			{Name: "LE BEDOU", Rate: 100},
		},
	}

	dropped := GiveMonsterLoot(&player, monster)

	if len(dropped) != 1 || dropped[0] != "LE BEDOU" {
		t.Fatalf("expected LE BEDOU to drop, got %#v", dropped)
	}
	if player.Inventory.Items["Le Bédou"] != 1 {
		t.Fatalf("expected Le Bédou in inventory, got %d", player.Inventory.Items["Le Bédou"])
	}
}
