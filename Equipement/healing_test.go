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