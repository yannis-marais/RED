package ProjetRED

import (
	"testing"

	personnage "ProjetRED/Personnage"
)

func TestUseConsumableHealsAndConsumesOne(t *testing.T) {
	player := personnage.CharacterCreation("test", personnage.Classes["Ronin"])
	player.PV = 1
	AddConsumable(&player, HealingPotion)

	UseConsumable(&player, HealingPotion)

	if player.PV != player.PVMax {
		t.Fatalf("expected PV to be capped at %d, got %d", player.PVMax, player.PV)
	}
	if player.Inventory.Consumables[HealingPotion.Name] != 0 {
		t.Fatalf("expected one potion to be consumed")
	}
}

func TestUpdateCooldownsDecrementsByOne(t *testing.T) {
	player := personnage.CharacterCreation("test", personnage.Classes["Ronin"])
	player.Cooldowns["Healing"] = 3

	player.UpdateCooldowns()

	if player.Cooldowns["Healing"] != 2 {
		t.Fatalf("expected cooldown 2, got %d", player.Cooldowns["Healing"])
	}
}

func TestEquipItemByNameReplacesOldEquipmentAndUpdatesStats(t *testing.T) {
	player := personnage.CharacterCreation("test", personnage.Classes["Ronin"])
	player.Inventory.Items = map[string]int{
		"Swordshield":  1,
		"Bandit Spear": 1,
	}
	player.Weapon = "Bandit_Spear"

	EquipItem(&player, "Swordshield")

	if player.Weapon != "Swordshield" {
		t.Fatalf("expected weapon to be equipped as Swordshield, got %s", player.Weapon)
	}
	if player.Inventory.Items["Swordshield"] != 0 {
		t.Fatalf("expected Swordshield to be removed from inventory after equip, got %d", player.Inventory.Items["Swordshield"])
	}
	if _, ok := player.Inventory.Items["Bandit Spear"]; !ok {
		t.Fatalf("expected previous weapon to be returned to inventory")
	}
	if player.Strength != player.Classe.Strength+5 {
		t.Fatalf("expected strength %d, got %d", player.Classe.Strength+5, player.Strength)
	}
	expectedDefense := player.Classe.Defense + 10 + 12 + 7 + 10
	if player.Defense != expectedDefense {
		t.Fatalf("expected defense %d from all equipped items, got %d", expectedDefense, player.Defense)
	}

	RecalculateStats(&player)
	if player.Defense != expectedDefense {
		t.Fatalf("expected recalculating stats to stay at %d, got %d", expectedDefense, player.Defense)
	}
}
