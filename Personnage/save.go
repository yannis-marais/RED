package ProjetRED

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func DefaultSavePath() string {
	return filepath.Join(".", "save.json")
}

func SaveCharacterToFile(path string, character Character) error {
	if path == "" {
		path = DefaultSavePath()
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(character, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644)
}

func LoadCharacterFromFile(path string) (Character, error) {
	if path == "" {
		path = DefaultSavePath()
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Character{}, err
	}

	var character Character
	if err := json.Unmarshal(data, &character); err != nil {
		return Character{}, err
	}

	if character.Inventory.Capacity == 0 {
		character.Inventory.Capacity = DefaultInventoryCapacity
	}
	if character.Inventory.Items == nil {
		character.Inventory.Items = make(map[string]int)
	}
	if character.Inventory.Consumables == nil {
		character.Inventory.Consumables = make(map[string]int)
	}
	if character.Inventory.Materials == nil {
		character.Inventory.Materials = make(map[string]int)
	}
	if character.Inventory.SkillBooks == nil {
		character.Inventory.SkillBooks = make(map[string]int)
	}
	if character.Cooldowns == nil {
		character.Cooldowns = make(map[string]int)
	}
	if character.Skills == nil {
		character.Skills = make(map[string]Skill)
	}
	if character.Effects == nil {
		character.Effects = []StatusEffect{}
	}

	return character, nil
}
