package ProjetRED

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveAndLoadCharacter(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "save.json")

	player := CharacterCreation("Jean", Classes["Ronin"])
	player.PV = 25
	player.Purse = 320
	player.Inventory.Items["Swordshield"] = 1
	player.Inventory.Consumables["Healing Potion"] = 2
	player.Skills["Strike"] = Skill{Name: "Strike", Damage: 12, Reiki: 5}

	if err := SaveCharacterToFile(path, player); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	loaded, err := LoadCharacterFromFile(path)
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}

	if loaded.Nom != player.Nom {
		t.Fatalf("expected name %q, got %q", player.Nom, loaded.Nom)
	}
	if loaded.PV != player.PV {
		t.Fatalf("expected PV %d, got %d", player.PV, loaded.PV)
	}
	if loaded.Inventory.Items["Swordshield"] != 1 {
		t.Fatalf("expected item inventory to persist")
	}
	if loaded.Inventory.Consumables["Healing Potion"] != 2 {
		t.Fatalf("expected consumables to persist")
	}
	if _, ok := loaded.Skills["Strike"]; !ok {
		t.Fatalf("expected skills to persist")
	}

	if _, err := os.Stat(path); err != nil {
		t.Fatalf("save file was not created: %v", err)
	}
}
