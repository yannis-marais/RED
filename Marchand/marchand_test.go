package ProjetRED

import (
	"testing"

	personnage "ProjetRED/Personnage"
)

func TestUpgradeInventoryUsesPurseAndAddsCapacity(t *testing.T) {
	player := personnage.CharacterCreation("test", personnage.Classes["Ronin"])
	player.Purse = 100

	if !UpgradeInventory(&player) {
		t.Fatal("expected inventory upgrade to succeed")
	}
	if player.Purse != 0 {
		t.Fatalf("expected purse 0, got %d", player.Purse)
	}
	if player.Inventory.Capacity != personnage.DefaultInventoryCapacity+5 {
		t.Fatalf("expected capacity %d, got %d", personnage.DefaultInventoryCapacity+5, player.Inventory.Capacity)
	}
}

func TestUpgradeInventoryRejectsInsufficientPurse(t *testing.T) {
	player := personnage.CharacterCreation("test", personnage.Classes["Ronin"])
	player.Purse = 99

	if UpgradeInventory(&player) {
		t.Fatal("La demande à été refusé")
	}
	if player.Inventory.Capacity != personnage.DefaultInventoryCapacity {
		t.Fatalf("Pas assez de purse")
	}
}
