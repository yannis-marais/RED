package ProjetRED

import (
	"fmt"

	personnage "ProjetRED/Personnage"
)

func EquipItem(p *personnage.Character, slot string, item Item) {
	if p == nil {
		return
	}

	switch slot {
	case "weapon", "Weapon":
		p.Weapon = item.Name
	case "armor", "Armor":
		p.Armor = item.Name
	case "boots", "Boots":
		p.Boots = item.Name
	case "helmet", "Helmet":
		p.Helmet = item.Name
	default:
		fmt.Println("Slot inconnu :", slot)
		return
	}

	RecalculateStats(p)
}

func FindItemByName(itemName string) (string, Item, bool) {
	if item, ok := Items[itemName]; ok {
		return itemName, item, true
	}

	for key, item := range Items {
		if item.Name == itemName {
			return key, item, true
		}
	}

	return "", Item{}, false
}

func EquipItemByName(p *personnage.Character, itemName string, itemType ItemType) {
	if p == nil {
		return
	}

	itemKey, item, ok := FindItemByName(itemName)
	if !ok {
		fmt.Println("Objet introuvable :", itemName)
		return
	}
	if item.Type != itemType {
		fmt.Printf("L'objet %s n'est pas de type %s.\n", item.Name, itemType)
		return
	}

	if p.Inventory.Items == nil {
		p.Inventory.Items = make(map[string]int)
	}

	if qty, exists := p.Inventory.Items[item.Name]; !exists || qty <= 0 {
		if qty, exists := p.Inventory.Items[itemKey]; !exists || qty <= 0 {
			fmt.Println("Tu n'as pas", item.Name, "dans l'inventaire.")
			return
		}
	}

	var currentSlot string
	switch itemType {
	case Weapon:
		currentSlot = p.Weapon
	case Armor:
		currentSlot = p.Armor
	case Boots:
		currentSlot = p.Boots
	case Helmet:
		currentSlot = p.Helmet
	default:
		fmt.Println("Type d'équipement inconnu :", itemType)
		return
	}

	if currentSlot != "" && currentSlot != itemKey {
		if oldItem, ok := Items[currentSlot]; ok {
			AddItem(p, oldItem)
		}
	}

	if qty, exists := p.Inventory.Items[item.Name]; exists && qty > 0 {
		if qty <= 1 {
			delete(p.Inventory.Items, item.Name)
		} else {
			p.Inventory.Items[item.Name] = qty - 1
		}
	} else if qty, exists := p.Inventory.Items[itemKey]; exists && qty > 0 {
		if qty <= 1 {
			delete(p.Inventory.Items, itemKey)
		} else {
			p.Inventory.Items[itemKey] = qty - 1
		}
	}

	switch itemType {
	case Weapon:
		p.Weapon = itemKey
	case Armor:
		p.Armor = itemKey
	case Boots:
		p.Boots = itemKey
	case Helmet:
		p.Helmet = itemKey
	}

	RecalculateStats(p)
	fmt.Println(item.Name, "équipé.")
}
