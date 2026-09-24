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
	if player.Purse != 50 {
		t.Fatalf("expected purse 50, got %d", player.Purse)
	}
	if player.Inventory.Capacity != personnage.DefaultInventoryCapacity+5 {
		t.Fatalf("expected capacity %d, got %d", personnage.DefaultInventoryCapacity+5, player.Inventory.Capacity)
	}
}

func TestUpgradeInventoryRejectsInsufficientPurse(t *testing.T) {
	player := personnage.CharacterCreation("test", personnage.Classes["Ronin"])
	player.Purse = 49

	if UpgradeInventory(&player) {
		t.Fatal("expected inventory upgrade to fail when purse is too low")
	}
	if player.Inventory.Capacity != personnage.DefaultInventoryCapacity {
		t.Fatalf("expected capacity to remain %d, got %d", personnage.DefaultInventoryCapacity, player.Inventory.Capacity)
	}
}

func TestBuyConsumableUsesAddConsumable(t *testing.T) {
	player := personnage.CharacterCreation("test", personnage.Classes["Ronin"])
	player.Purse = 100

	if !acheterObjet(&player, "Pain") {
		t.Fatal("expected Pain purchase to succeed")
	}
	if player.Inventory.Consumables["Pain"] != 1 {
		t.Fatalf("expected Pain quantity 1, got %d", player.Inventory.Consumables["Pain"])
	}
	if player.Purse != 85 {
		t.Fatalf("expected purse 85 after buying Pain, got %d", player.Purse)
	}

	if !acheterObjet(&player, "Healing_Potion") {
		t.Fatal("expected Healing_Potion purchase to succeed")
	}
	if player.Inventory.Consumables["Healing Potion"] != 1 {
		t.Fatalf("expected Healing Potion quantity 1, got %d", player.Inventory.Consumables["Healing Potion"])
	}
	if player.Purse != 65 {
		t.Fatalf("expected purse 65 after buying both consumables, got %d", player.Purse)
	}
}
